package controllers

import (
	"sync"
	"time"

	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

type CommonAttribute struct {
	Name      string
	DB        *gorm.DB
	K8sClient *kubernetes.Clientset
	stopChan  chan struct{}
	aliveChan chan struct{}
}

func (ca *CommonAttribute) chanStop() chan struct{} {
	return ca.stopChan
}

func (ca *CommonAttribute) chanAlive() chan struct{} {
	return ca.aliveChan
}

type Paging struct {
	Limit  int
	Offset int
}

type Controller interface {
	Name() string
	sync(stopChan chan struct{})
	initListerAndInformer()
	total() int64
	CountWithConditions(condition string) int64
	ListWithConditions(condition string, paging *Paging, order string) (int64, interface{}, error)
	Lister() interface{}
	chanAlive() chan struct{}
	chanStop() chan struct{}
}

func countWithConditions(db *gorm.DB, conditions string, object interface{}) int64 {
	var count int64
	if len(conditions) == 0 {
		db.Model(object).Count(&count)
	} else {
		db.Model(object).Where(conditions).Count(&count)
	}
	return count
}

func listWithConditions(db *gorm.DB, total *int64, object, list interface{}, conditions string, paging *Paging, order string) {
	if len(conditions) == 0 {
		db.Model(object).Count(total)
	} else {
		db.Model(object).Where(conditions).Count(total)
	}

	if paging != nil {
		if len(conditions) > 0 {
			db.Where(conditions).Order(order).Limit(paging.Limit).Offset(paging.Offset).Find(list)
		} else {
			db.Order(order).Limit(paging.Limit).Offset(paging.Offset).Find(list)
		}

	} else {
		if len(conditions) > 0 {
			db.Where(conditions).Order(order).Find(list)
		} else {
			db.Order(order).Find(list)
		}
	}
}

func handleCrash(ctl Controller) {
	close(ctl.chanAlive())
	if err := recover(); err != nil {
		return
	}
}

func hasSynced(ctl Controller) bool {
	totalInDb := ctl.CountWithConditions("")
	totalInK8s := ctl.total()

	return totalInDb == totalInK8s
}

func checkAndResync(ctl Controller, stopChan chan struct{}) {
	defer close(stopChan)

	lastTime := time.Now()

	for {
		select {
		case <-ctl.chanStop():
			return
		default:
			if time.Since(lastTime) < time.Minute {
				time.Sleep(time.Second * 20)
				break
			}

			lastTime = time.Now()
			if !hasSynced(ctl) {
				close(stopChan)
				stopChan = make(chan struct{})
				go ctl.sync(stopChan)
			}
		}
	}
}

func listAndWatch(ctl Controller, wg *sync.WaitGroup) {
	defer handleCrash(ctl)
	// defer ctl.CloseDB()
	defer wg.Done()
	stopChan := make(chan struct{})

	go ctl.sync(stopChan)

	checkAndResync(ctl, stopChan)
}
