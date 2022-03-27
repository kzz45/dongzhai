package controllers

import (
	"dongzhai/models"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	lister_appsv1 "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
)

type DeploymentCtl struct {
	CommonAttribute
	lister   lister_appsv1.DeploymentLister
	informer cache.SharedIndexInformer
}

func (ctl *DeploymentCtl) Name() string {
	return ctl.CommonAttribute.Name
}

func (ctl *DeploymentCtl) generateObject(item v1.Deployment) *models.Deployment {
	var app string
	var status string
	var updateTime time.Time
	name := item.Name
	namespace := item.Namespace
	availablePodNum := item.Status.AvailableReplicas
	desirePodNum := *item.Spec.Replicas
	release := item.ObjectMeta.Labels["release"]
	chart := item.ObjectMeta.Labels["chart"]

	if len(release) > 0 && len(chart) > 0 {
		app = release + "/" + chart
	}

	for _, conditon := range item.Status.Conditions {
		if conditon.Type == "Available" {
			updateTime = conditon.LastUpdateTime.Time
		}
	}
	if updateTime.IsZero() {
		updateTime = time.Now()
	}
	if item.Annotations["state"] == "stop" {
		status = models.Stopped
	} else {
		if availablePodNum >= desirePodNum {
			status = models.Running
		} else {
			status = models.Updating
		}
	}

	return &models.Deployment{
		App:       app,
		Name:      name,
		Status:    status,
		Namespace: namespace,
		BaseModel: models.BaseModel{
			UpdatedAt: updateTime,
		},
	}
}

func (ctl *DeploymentCtl) sync(stopChan chan struct{}) {
	db := ctl.DB
	if db.Migrator().HasTable(&models.Deployment{}) {
		db.Migrator().DropTable(&models.Deployment{})
	}
	db.Migrator().CreateTable(&models.Deployment{})

	ctl.initListerAndInformer()
	list, err := ctl.lister.List(labels.Everything())
	if err != nil {
		return
	}

	for _, item := range list {
		obj := ctl.generateObject(*item)
		db.Create(obj)
	}
	ctl.informer.Run(stopChan)
}

func (ctl *DeploymentCtl) initListerAndInformer() {
	db := ctl.DB

	informerFactory := informers.NewSharedInformerFactory(ctl.K8sClient, time.Second*10)

	ctl.lister = informerFactory.Apps().V1().Deployments().Lister()

	informer := informerFactory.Apps().V1().Deployments().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {

			object := obj.(*v1.Deployment)
			mysqlObject := ctl.generateObject(*object)
			db.Create(mysqlObject)
		},
		UpdateFunc: func(old, new interface{}) {
			object := new.(*v1.Deployment)
			mysqlObject := ctl.generateObject(*object)
			db.Save(mysqlObject)
		},
		DeleteFunc: func(obj interface{}) {
			var deploy models.Deployment
			object := obj.(*v1.Deployment)
			db.Where("name=? And namespace=?", object.Name, object.Namespace).Find(&deploy)
			db.Delete(deploy)
		},
	})
	ctl.informer = informer
}

func (ctl *DeploymentCtl) total() int {
	list, err := ctl.lister.List(labels.Everything())
	if err != nil {
		return 0
	}
	return len(list)
}

func (ctl *DeploymentCtl) CountWithConditions(conditions string) int64 {
	var object models.Deployment

	return countWithConditions(ctl.DB, conditions, &object)
}

func (ctl *DeploymentCtl) ListWithConditions(conditions string, paging *Paging) (int64, interface{}, error) {
	var list []models.Deployment
	var object models.Deployment
	var total int64

	order := "updateTime desc"

	listWithConditions(ctl.DB, &total, &object, &list, conditions, paging, order)

	return total, list, nil
}
