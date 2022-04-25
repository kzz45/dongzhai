package client

import (
	"dongzhai/db"
	"dongzhai/models/k8s"
	"errors"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	defaultQPS          = 1e6
	defaultBurst        = 1e6
	defaultResyncPeriod = 30 * time.Second
)

// 通过kubeconfig获取client
func GetK8sClientWithConfig(cluster k8s.Cluster) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
	config.QPS = defaultQPS
	config.Burst = defaultBurst
	if err != nil {
		return nil, errors.New("kubeconfig content error")
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.New("new k8s client with kubeconfig error")
	}
	return clientSet, nil
}

// 通过token获取client
func GetK8sClientWithToken(cluster k8s.Cluster) (*kubernetes.Clientset, error) {
	config := &rest.Config{
		Host:        cluster.Addr,
		BearerToken: cluster.Token,
		QPS:         defaultQPS,
		Burst:       defaultBurst,
	}
	config.TLSClientConfig.Insecure = true
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.New("new k8s client with token error")
	}
	return clientset, nil
}

// 多个集群的k8s客户端
func NewK8SClients() (map[uint]*kubernetes.Clientset, error) {
	var clusters []k8s.Cluster
	if err := db.GlobalGorm.Find(&clusters).Error; err != nil {
		return nil, err
	}
	clientset_map := make(map[uint]*kubernetes.Clientset)
	for _, cluster := range clusters {
		config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
		if err != nil {
			continue
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			continue
		}
		clientset_map[cluster.ID] = clientSet
	}
	return clientset_map, nil
}
