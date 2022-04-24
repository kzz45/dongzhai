package k8s

import "dongzhai/models"

type Container struct {
	models.BaseModel
	DeploymentId    *uint           `json:"deployment_id"`                                 // 所属Deploy
	Deployment      Deployment      `json:"deployment" gorm:"foreignKey:DeploymentId"`     //
	JobId           *uint           `json:"job_id"`                                        // 所属Job
	Job             Job             `json:"job" gorm:"foreignKey:JobId"`                   //
	Name            string          `json:"name" gorm:"primaryKey"`                        // 容器名称
	Desc            string          `json:"desc"`                                          // 容器描述
	Namespace       string          `json:"namespace" gorm:"primaryKey"`                   // 容器命名空间
	Image           string          `json:"image"`                                         // 容器镜像名称
	ImageVersion    string          `json:"image_version"`                                 // 容器镜像版本
	ImagePullPolicy string          `json:"image_pull_policy"`                             // 容器拉取策略 Always/Never/IfNotPresent
	RestartPolicy   string          `json:"restart_policy"`                                // 重启策略
	Status          string          `json:"status" gorm:"default:creating"`                // 容器状态
	CPULimit        int             `json:"cpu_limit"`                                     // CPU限制
	CPURequest      int             `json:"cpu_request"`                                   //
	MemLimit        int             `json:"mem_limit"`                                     // 内存限制
	MemRequest      int             `json:"mem_request"`                                   //
	Environs        MapString       `json:"environs"`                                      // 环境变量
	Commands        MapString       `json:"commands"`                                      // 启动命名
	ContainerPorts  []ContainerPort `json:"container_ports" gorm:"foreignKey:ContainerId"` // 容器端口
	Labels          MapString       `json:"labels"`                                        // 标签
	Annotation      MapString       `json:"annotations"`                                   // 注解
}

func (Container) TableName() string {
	return models.TableNameContainer
}

type ContainerPort struct {
	models.BaseModel
	ContainerId   uint      `json:"container_id"`
	Container     Container `json:"container" gorm:"foreignKey:ContainerId"`
	Name          string    `json:"name" gorm:"primaryKey"`
	Protocol      string    `json:"protocol"`
	ContainerPort int32     `json:"container_port"`
}

func (ContainerPort) TableName() string {
	return models.TableNameContainerPort
}
