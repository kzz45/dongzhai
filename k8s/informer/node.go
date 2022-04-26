package informer

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/informers"
	lister_corev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type NodeCtl struct {
	lister   lister_corev1.NodeLister
	informer cache.SharedIndexInformer
	CommonAttribute
}

func (ctl *NodeCtl) Name() string {
	return ctl.CommonAttribute.Name
}

func (ctl *NodeCtl) initInformer() {
	// db := ctl.DB
	informerFactory := informers.NewSharedInformerFactory(ctl.K8sClient, time.Second*ResyncCircle)
	ctl.lister = informerFactory.Core().V1().Nodes().Lister()
	informer := informerFactory.Core().V1().Nodes().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// object := obj.(*corev1.Node)
			// mysqlObject := ctl.generateObject(*object)
			// logrus.Warnf("create: %v", mysqlObject)
			// db.Create(mysqlObject)
			logrus.Errorln("add...")
		},
		UpdateFunc: func(old, new interface{}) {
			// object := new.(*corev1.Node)
			// mysqlObject := ctl.generateObject(*object)
			// db.Updates(mysqlObject)
		},
		DeleteFunc: func(obj interface{}) {
			// var item models.KNode
			// object := obj.(*corev1.Node)
			// db.Where("name=? ", object.Name).Find(&item).Delete(item)
		},
	})
	ctl.informer = informer
}

func (ctl *NodeCtl) sync(stopChan chan struct{}, wg *sync.WaitGroup) {

	// db := ctl.DB
	defer wg.Done()
	ctl.initInformer()
	// list, err := ctl.lister.List(labels.Everything())
	// if err != nil {
	// 	return
	// }
	// for _, item := range list {
	// obj := ctl.generateObject(*item)
	// db.Create(obj)
	// }
	ctl.informer.Run(stopChan)
}
