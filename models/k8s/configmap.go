package k8s

import "dongzhai/models"

type ConfigMap struct {
	models.BaseModel
	ClusterId  uint      `json:"cluster_id"`                  // 所属集群
	Cluster    Cluster   `json:"cluster"`                     //
	ProjectId  uint      `json:"project_id"`                  // 所属项目
	Project    Project   `json:"project"`                     //
	Name       string    `json:"name" gorm:"primaryKey"`      //
	Namespace  string    `json:"namespace" gorm:"primaryKey"` //
	Data       KeyValue  `json:"data"`                        //
	Labels     MapString `json:"labels"`                      //
	Annotation MapString `json:"annotations"`                 //
}

func (ConfigMap) TableName() string {
	return models.TableNameConfigmap
}
