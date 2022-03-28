package controllers

import (
	"dongzhai/models"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	lister_corev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type NodeCtl struct {
	CommonAttribute
	lister   lister_corev1.NodeLister
	informer cache.SharedIndexInformer
}

func (ctl *NodeCtl) generateObject(item v1.Node) *models.ClusterNode {
	var status string

	name := item.Name
	createTime := item.ObjectMeta.CreationTimestamp.Time

	for _, condition := range item.Status.Conditions {
		if condition.Type == "Ready" {
			if condition.Status == "True" {
				status = models.Running
			} else {
				status = models.Error
			}

		}
	}

	object := &models.ClusterNode{
		Name:   name,
		Status: status,
		BaseModel: models.BaseModel{
			CreatedAt: createTime,
		},
	}
	return object
}

func (ctl *NodeCtl) Name() string {
	return ctl.CommonAttribute.Name
}

func (ctl *NodeCtl) sync(stopChan chan struct{}) {
	db := ctl.DB

	if db.Migrator().HasTable(&models.ClusterNode{}) {
		db.Migrator().DropTable(&models.ClusterNode{})
	}

	db.Migrator().CreateTable(&models.ClusterNode{})

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

func (ctl *NodeCtl) initListerAndInformer() {
	db := ctl.DB

	informerFactory := informers.NewSharedInformerFactory(ctl.K8sClient, time.Second*10)

	ctl.lister = informerFactory.Core().V1().Nodes().Lister()

	informer := informerFactory.Core().V1().Nodes().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			object := obj.(*v1.Node)
			mysqlObject := ctl.generateObject(*object)
			db.Create(mysqlObject)
		},
		UpdateFunc: func(old, new interface{}) {
			object := new.(*v1.Node)
			mysqlObject := ctl.generateObject(*object)
			db.Updates(mysqlObject)
		},
		DeleteFunc: func(obj interface{}) {
			var item models.ClusterNode
			object := obj.(*v1.Node)
			db.Where("name=? ", object.Name).Find(&item)
			db.Delete(item)
		},
	})

	ctl.informer = informer
}

func (ctl *NodeCtl) total() int64 {
	list, err := ctl.lister.List(labels.Everything())
	if err != nil {
		return 0
	}
	return int64(len(list))
}

func (ctl *NodeCtl) CountWithConditions(conditions string) int64 {
	var object models.ClusterNode

	return countWithConditions(ctl.DB, conditions, &object)
}

func (ctl *NodeCtl) ListWithConditions(conditions string, paging *Paging, order string) (int64, interface{}, error) {
	var list []models.ClusterNode
	var object models.ClusterNode
	var total int64

	if len(order) == 0 {
		order = "createTime desc"
	}

	listWithConditions(ctl.DB, &total, &object, &list, conditions, paging, order)

	return total, list, nil
}

func (ctl *NodeCtl) Lister() interface{} {
	return ctl.lister
}
