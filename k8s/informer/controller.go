package informer

import (
	"dongzhai/client"
	"dongzhai/db"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

var ResourceControllers resourceControllers

type resourceControllers struct {
	Controllers map[string]Controller
	k8sClient   *kubernetes.Clientset
}

type Controller interface {
	Name() string
	initInformer()
	sync(stopChan chan struct{}, wg *sync.WaitGroup)
	// total() int64
	// CountWithConditions(condition string) int64
	chanAlive() chan struct{}
	chanStop() chan struct{}
}

// func countWithConditions(db *gorm.DB, conditions string, object interface{}) int64 {
// 	var count int64
// 	if len(conditions) == 0 {
// 		db.Model(object).Count(&count)
// 	} else {
// 		db.Model(object).Where(conditions).Count(&count)
// 	}
// 	return count
// }

// func handleCrash(ctl Controller) {
// 	close(ctl.chanAlive())
// 	if err := recover(); err != nil {
// 		logrus.Errorf("panic in: %s controller's listAndWatch function, reason: %s", ctl.Name(), err)
// 		return
// 	}
// }

// func hasSynced(ctl Controller) bool {
// 	totalInDb := ctl.CountWithConditions("cluster_id = ?")
// 	totalInK8s := ctl.total()
// 	return totalInDb == totalInK8s
// }

// func listAndWatch(ctl Controller, wg *sync.WaitGroup) {
// 	defer handleCrash(ctl)
// 	defer wg.Done()

// 	stopChan := make(chan struct{})
// 	go ctl.sync(stopChan)

// 	checkAndResync(ctl, stopChan)
// }

// func checkAndResync(ctl Controller, stopChan chan struct{}) {
// 	defer close(stopChan)
// 	lastTime := time.Now()

// 	for {
// 		select {
// 		case <-ctl.chanStop():
// 			return
// 		default:
// 			if time.Since(lastTime) < checkPeriod {
// 				time.Sleep(sleepPeriod)
// 				break
// 			}
// 			lastTime = time.Now()
// 			if !hasSynced(ctl) {
// 				logrus.Errorf("the data in db and k8s is inconsistent, resync: %s controller", ctl.Name())
// 				close(stopChan)
// 				stopChan = make(chan struct{})
// 				go ctl.sync(stopChan)
// 			}
// 		}
// 	}
// }

func (rec *resourceControllers) runController(cluster_id uint, name string, stopChan chan struct{}, wg *sync.WaitGroup) {
	var ctl Controller
	attr := CommonAttribute{
		Name:      name,
		ClusterID: cluster_id,
		DB:        db.GlobalGorm,
		K8sClient: rec.k8sClient,
	}
	switch name {
	case Nodes:
		ctl = &NodeCtl{CommonAttribute: attr}
	default:
		return
	}
	rec.Controllers[name] = ctl
	wg.Add(1)
	go ctl.sync(stopChan, wg)
}

func Run(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	clientset_map, err := client.NewK8SClients()
	if err != nil {
		panic(err)
	}
	for cluster_id, clientset := range clientset_map {
		ResourceControllers = resourceControllers{k8sClient: clientset, Controllers: make(map[string]Controller)}
		for _, item := range []string{Deployments, Nodes, Pods, Secrets, ConfigMaps} {
			ResourceControllers.runController(cluster_id, item, stopChan, wg)
		}
		for {
			for ctl_name, controller := range ResourceControllers.Controllers {
				select {
				case <-stopChan:
					return
				case _, isClose := <-controller.chanAlive():
					if !isClose {
						logrus.Errorf("controller: %s have stopped restart it now", ctl_name)
						ResourceControllers.runController(cluster_id, ctl_name, stopChan, wg)
					}
				default:
					time.Sleep(sleepPeriod)
				}
			}
		}
	}
}

func RunInformer() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	wg.Add(1)

	go Run(stopChan, &wg)
}
