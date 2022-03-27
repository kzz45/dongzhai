package controllers

import (
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

type CommonAttribute struct {
	K8sClient *kubernetes.Clientset
	Name      string
	DB        *gorm.DB
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
	total() int
	CountWithConditions(condition string) int64
	ListWithConditions(condition string, paging *Paging) (int, interface{}, error)
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
