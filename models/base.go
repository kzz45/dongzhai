package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableNameUser          = "user"                   //
	TableNameRole          = "role"                   //
	TableNameUserGroup     = "group"                  //
	TableNameCloud         = "cloud"                  //
	TableNameDomain        = "domain"                 //
	TableNameProduct       = "product"                //
	TableNameDomainCert    = "domain_cert"            //
	TableNameDomainRecord  = "domain_record"          //
	TableNameTask          = "monitor_task"           //
	TableNameLabel         = "monitor_label"          //
	TableNameGroup         = "monitor_group"          //
	TableNameServer        = "monitor_server"         //
	TableNameInstance      = "monitor_instance"       //
	TableNameReceiver      = "monitor_receiver"       //
	TableNameAlertRule     = "monitor_alert_rule"     //
	TableNameAlertRoute    = "monitor_alert_route"    //
	TableNameLabelValue    = "monitor_label_value"    //
	TableNameInstanceGroup = "monitor_instance_group" //
	TableNamePod           = "k8s_pod"                //
	TableNameJob           = "k8s_job"                //
	TableNameNode          = "k8s_node"               //
	TableNameImage         = "k8s_image"              //
	TableNameSecret        = "k8s_secret"             //
	TableNameProject       = "k8s_project"            //
	TableNameCluster       = "k8s_cluster"            //
	TableNameService       = "k8s_service"            //
	TableNameIngress       = "k8s_ingress"            //
	TableNameRegistry      = "k8s_registry"           //
	TableNameContainer     = "k8s_container"          //
	TableNameConfigmap     = "k8s_configmap"          //
	TableNameDaemonset     = "k8s_daemonset"          //
	TableNameDeployment    = "k8s_deployment"         //
	TableNameStatefulset   = "k8s_statefulset"        //
	TableNameIngressRule   = "k8s_ingress_rule"       //
	TableNameServicePort   = "k8s_service_port"       //
	TableNameImageVersion  = "k8s_image_version"      //
	TableNameContainerPort = "k8s_container_port"     //
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Pagination struct {
	Page  int         `json:"page" form:"page"`
	Size  int         `json:"size" form:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
