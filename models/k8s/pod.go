package k8s

import "dongzhai/models"

type Pod struct {
	models.BaseModel
	ClusterId    *uint      `json:"cluster_id"`                                // 所属集群
	Cluster      Cluster    `json:"cluster"`                                   //
	DeploymentId *uint      `json:"deployment_id"`                             // 所属Deploy
	Deployment   Deployment `json:"deployment" gorm:"foreignKey:DeploymentId"` //
	Name         string     `json:"name" gorm:"primaryKey"`                    // 容器组唯一名称
	Namespace    string     `json:"namespace" gorm:"primaryKey"`               // 命名空间
	GenerateName string     `json:"generate_name"`                             // 容器组生成的名称
	Status       string     `json:"status"`                                    // 状态
	Reason       string     `json:"reason"`                                    // 原因
	Message      string     `json:"message"`                                   // 信息
	PodIP        string     `json:"pod_ip"`                                    // 容器组IP
	NodeIP       string     `json:"node_ip"`                                   // 节点IP
	NodeName     string     `json:"node_name"`                                 // 节点名称
	RestartCount int        `json:"restart_count"`                             // 重启次数
	Labels       MapString  `json:"labels"`                                    // 标签
	Annotation   MapString  `json:"annotations"`                               // 注解
}

func (Pod) TableName() string {
	return models.TableNamePod
}
