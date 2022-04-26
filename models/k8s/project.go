package k8s

import "dongzhai/models"

type Project struct {
	models.BaseModel
	Name        string           `json:"name"`                                     //
	Desc        string           `json:"desc"`                                     //
	ProductId   uint             `json:"product_id"`                               //
	Product     models.Product   `json:"product" gorm:"foreignKey:ProductId"`      // 产品名称作为命名空间名称
	ClusterId   uint             `json:"cluster_id"`                               //
	Cluster     Cluster          `json:"cluster" gorm:"foreignKey:ClusterId"`      // 关联集群
	UserGroupId uint             `json:"user_group_id"`                            //
	UserGroup   models.UserGroup `json:"user_group" gorm:"foreignKey:UserGroupId"` // 关联用户组
	// Deployments []Deployment   `json:"deployments"`                              //
	// Jobs        []Job          `json:"jobs"`                                     //
}

func (Project) TableName() string {
	return models.TableNameProject
}
