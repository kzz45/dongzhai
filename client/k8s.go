package client

import (
	"dongzhai/models"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8SClients() (map[string]*kubernetes.Clientset, error) {
	var clusters []models.Cluster
	if err := DBClient.Find(&clusters).Error; err != nil {
		return nil, err
	}
	clientset_map := make(map[string]*kubernetes.Clientset)
	for _, cluster := range clusters {
		config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
		if err != nil {
			continue
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			continue
		}
		clientset_map[cluster.Name] = clientSet
	}
	return clientset_map, nil
}
