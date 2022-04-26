package informer

import (
	"time"

	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

const (
	ResyncCircle          = 60
	Stopped               = "stopped"
	PvcPending            = "pending"
	Running               = "running"
	Updating              = "updating"
	Failed                = "failed"
	Unfinished            = "unfinished"
	Completed             = "completed"
	Pause                 = "pause"
	Warning               = "warning"
	Error                 = "error"
	Creator               = "creator"
	Namespaces            = "namespaces"
	Deployments           = "deployments"
	Daemonsets            = "daemonsets"
	Statefulsets          = "statefulsets"
	Replicasets           = "replicasets"
	Services              = "services"
	Ingresses             = "ingresses"
	Pods                  = "pods"
	Jobs                  = "jobs"
	Roles                 = "roles"
	Nodes                 = "nodes"
	Secrets               = "secrets"
	Cronjobs              = "cronjobs"
	ConfigMaps            = "configmaps"
	RoleBindings          = "role-bindings"
	ClusterRoles          = "cluster-roles"
	StorageClasses        = "storage-classes"
	ClusterRoleBindings   = "cluster-role-bindings"
	PersistentVolumeClaim = "persistent-volume-claims"
	ManagedBy             = "dongzhai.io"
	checkPeriod           = 1 * time.Minute
	sleepPeriod           = 15 * time.Second
)

type CommonAttribute struct {
	Name      string
	ClusterID uint
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
