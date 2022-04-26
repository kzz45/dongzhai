package k8s

import "dongzhai/models"

type Secret struct {
	models.BaseModel
	ClusterId  uint      `json:"cluster_id"`                  // 所属集群
	Cluster    Cluster   `json:"cluster"`                     //
	ProjectId  uint      `json:"project_id"`                  //
	Project    Project   `json:"project"`                     //
	Name       string    `json:"name" gorm:"primaryKey"`      //
	Namespace  string    `json:"namespace" gorm:"primaryKey"` //
	Type       string    `json:"type"`                        // Opaque(默认)/dockerconfigjson/basic_auth(用户名和密码)/tls
	Data       KeyValue  `json:"data"`                        // Opaque/tls/basic_auth
	Addr       string    `json:"addr"`                        // harbor addr
	Username   string    `json:"username"`                    // harbor username
	Password   string    `json:"password"`                    // harbor password
	Annotation MapString `json:"annotations"`                 //
}

func (Secret) TableName() string {
	return models.TableNameSecret
}
