package k8s

import "dongzhai/models"

type Ingress struct {
	models.BaseModel
	ProjectId  uint          `json:"project_id"`                          // 关联项目
	Project    Project       `json:"project" gorm:"foreignKey:ProjectId"` //
	ClusterId  uint          `json:"cluster_id"`                          // 关联集群
	Cluster    Cluster       `json:"cluster" gorm:"foreignKey:ClusterId"` //
	ServiceId  uint          `json:"service_id"`                          // 关联服务
	Service    Service       `json:"service" gorm:"foreignKey:ServiceId"` //
	Name       string        `json:"name" gorm:"primaryKey"`              //
	Desc       string        `json:"desc"`                                //
	Status     int           `json:"status"`                              //
	Namespace  string        `json:"namespace" gorm:"primaryKey"`         //
	Rules      []IngressRule `json:"rules" gorm:"foreignKey:IngressId"`   //
	Labels     MapString     `json:"labels"`                              //
	Annotation MapString     `json:"annotations"`                         //
}

func (Ingress) TableName() string {
	return models.TableNameIngress
}

type IngressRule struct {
	models.BaseModel
	IngressId     uint        `json:"ingress_id"`                                   //
	Ingress       Ingress     `json:"ingress" gorm:"foreignKey:IngressId"`          //
	Host          string      `json:"host"`                                         //
	Path          string      `json:"path"`                                         //
	ServiceId     uint        `json:"service_id"`                                   // 关联服务
	Service       Service     `json:"service" gorm:"foreignKey:ServiceId"`          // 关联服务端口
	ServicePortId uint        `json:"service_port_id"`                              //
	ServicePort   ServicePort `json:"service_port" gorm:"foreignKey:ServicePortId"` //
}

func (IngressRule) TableName() string {
	return models.TableNameIngressRule
}
