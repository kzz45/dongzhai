package k8s

import "dongzhai/models"

type Service struct {
	models.BaseModel
	ProjectId    uint          `json:"project_id"`                                //
	Project      Project       `json:"project" gorm:"foreignKey:ProjectId"`       //
	ClusterId    uint          `json:"cluster_id"`                                //
	Cluster      Cluster       `json:"cluster" gorm:"foreignKey:ClusterId"`       //
	DeploymentId uint          `json:"deployment_id"`                             //
	Deployment   Deployment    `json:"deployment" gorm:"foreignKey:DeploymentId"` //
	Name         string        `json:"name"`                                      //
	Desc         string        `json:"desc"`                                      //
	Status       string        `json:"status"`                                    //
	Namespace    string        `json:"namespace"`                                 //
	VirtualIP    string        `json:"virtual_ip"`                                //
	ExternalIP   string        `json:"external_ip"`                               //
	NodePort     int           `json:"node_port"`                                 //
	ServiceType  string        `json:"service_type"`                              // ClusterIP/NodePort/LoadBalancer/ExternalName
	Labels       MapString     `json:"labels"`                                    //
	Annotation   MapString     `json:"annotations"`                               //
	ServicePorts []ServicePort `json:"service_ports" gorm:"foreignKey:ServiceId"` //
	Ingresses    []Ingress     `json:"ingresses" gorm:"foreignKey:ServiceId"`     //
	Original     int           `json:"original"`                                  //
}

func (Service) TableName() string {
	return models.TableNameService
}

type ServicePort struct {
	models.BaseModel
	ServiceId  uint    `json:"service_id"`                          //
	Service    Service `json:"service" gorm:"foreignKey:ServiceId"` //
	Name       string  `json:"name"`                                //
	Protocol   string  `json:"protocol"`                            //
	Port       int     `json:"port"`                                //
	TargetPort int     `json:"target_port"`                         //
}

func (ServicePort) TableName() string {
	return models.TableNameServicePort
}
