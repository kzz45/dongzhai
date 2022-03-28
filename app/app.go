package app

import (
	"dongzhai/apis"
	"dongzhai/client"
	"dongzhai/controllers"
	"sync"
)

func Run() {
	client.Run()

	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	wg.Add(1)
	go controllers.Run(stopChan, &wg)

	apis.Run()
}
