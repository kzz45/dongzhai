package controllers

import (
	"dongzhai/client"
	"sync"
	"time"

	"k8s.io/client-go/kubernetes"
)

type resourceControllers struct {
	Controllers map[string]Controller
	k8sClient   map[string]*kubernetes.Clientset
}

var ResourceControllers resourceControllers

func (rec *resourceControllers) runContoller(cluster_name string, name string, stopChan chan struct{}, wg *sync.WaitGroup) {
	var ctl Controller
	attr := CommonAttribute{
		Name:      name,
		DB:        client.NewDBClient(),
		K8sClient: rec.k8sClient[cluster_name],
		stopChan:  stopChan,
		aliveChan: make(chan struct{}),
	}
	switch name {
	case "Node":
		ctl = &NodeCtl{CommonAttribute: attr}
	default:
		return
	}
	rec.Controllers[name] = ctl
	wg.Add(1)
	go listAndWatch(ctl, wg)
}

func Run(stopChan chan struct{}, wg *sync.WaitGroup) {
	clientset_map, err := client.NewK8SClients()
	if err != nil {
		panic(err)
	}
	ResourceControllers = resourceControllers{k8sClient: clientset_map, Controllers: make(map[string]Controller)}
	for _, item := range []string{"Node"} {
		ResourceControllers.runContoller("", item, stopChan, wg)
	}

	for {
		for ctl_name, controller := range ResourceControllers.Controllers {
			select {
			case <-stopChan:
				return
			case _, isClose := <-controller.chanAlive():
				if !isClose {
					ResourceControllers.runContoller("", ctl_name, stopChan, wg)
				}
			default:
				time.Sleep(time.Second * 3)
			}
		}
	}
}
