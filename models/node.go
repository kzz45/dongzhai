package models

type ClusterNode struct {
	BaseModel
	Name   string `json:"name" gorm:"primary_key" `
	Status string `json:"status"`
}

func (ClusterNode) TableName() string {
	return "cluster_node"
}
