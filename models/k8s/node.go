package k8s

import "dongzhai/models"

type Node struct {
	models.BaseModel
	ClusterId   uint      `json:"cluster_id" `                         //
	Cluster     Cluster   `json:"cluster" gorm:"foreignKey:ClusterId"` //
	NodeName    string    `json:"node_name" gorm:"primaryKey"`         //
	ProviderID  string    `json:"provider_id" gorm:"primaryKey"`       //
	Status      string    `json:"status"`                              //
	InternalIP  string    `json:"internal_ip"`                         //
	ExternalIP  string    `json:"external_ip"`                         //
	CPUCap      int64     `json:"total_cpu"`                           //
	CPUAllocate int64     `json:"allocate_cpu"`                        //
	MemCap      int64     `json:"total_mem"`                           //
	MemAllocate int64     `json:"allocate_mem"`                        //
	Labels      MapString `json:"labels"`                              //
	Annotation  MapString `json:"annotations"`                         //
	Taints      Taints    `json:"taints"`                              //
}

func (Node) TableName() string {
	return models.TableNameNode
}
