package k8s

import "dongzhai/models"

type Deployment struct {
	models.BaseModel
	ProjectId    uint        `json:"project_id"`                                //
	Project      Project     `json:"project" gorm:"foreignKey:ProjectId"`       //
	ClusterId    uint        `json:"cluster_id"`                                //
	Cluster      Cluster     `json:"cluster" gorm:"foreignKey:ClusterId"`       //
	Name         string      `json:"name" gorm:"primaryKey"`                    // 部署名称唯一
	Desc         string      `json:"desc"`                                      //
	Status       string      `json:"status" gorm:"default:creating"`            // 状态
	Namespace    string      `json:"namespace" gorm:"primaryKey"`               // 命名空间
	Replicas     int32       `json:"replicas"`                                  // 副本数
	RollMax      int         `json:"roll_max"`                                  //
	RollMin      int         `json:"roll_min"`                                  //
	UpdatePolicy string      `json:"update_policy"`                             // 更新策略
	Available    int32       `json:"available"`                                 //
	Desire       int32       `json:"desire"`                                    //
	Services     []Service   `json:"services" gorm:"foreignKey:DeploymentId"`   // 包含的服务
	Original     int         `json:"original"`                                  // 是否原生标识
	Labels       MapString   `json:"labels"`                                    // 标签
	Annotation   MapString   `json:"annotations"`                               // 注解
	Containers   []Container `json:"containers" gorm:"foreignKey:DeploymentId"` //
}

func (Deployment) TableName() string {
	return models.TableNameDeployment
}
