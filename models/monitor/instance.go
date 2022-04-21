package monitor

import "dongzhai/models"

// 监控实例

type Instance struct {
	models.BaseModel
	ServerId    *uint           `json:"server_id"`                               // 关联的Prometheus Server
	Server      Server          `json:"server"  gorm:"foreignKey:ServerId"`      //
	ProductId   *uint           `json:"product_id"`                              // 关联的产品
	Product     models.Product  `json:"product" gorm:"foreignKey:ProductId"`     //
	Name        string          `json:"name"`                                    //
	UUID        string          `json:"uuid"`                                    //
	Type        string          `json:"type"`                                    // 实例类型(ecs/k8s/)
	Endpoint    string          `json:"endpoint"`                                //
	Enable      bool            `json:"enable"`                                  // 是否监控
	PublicIP    string          `json:"public_ip"`                               //
	PrivateIP   string          `json:"private_ip"`                              //
	UsePublicIP bool            `json:"use_public_ip"`                           // 是否使用外网IP监控
	Labels      []Label         `json:"labels" gorm:"many2many:instance_label;"` // 实例上的标签
	Groups      []InstanceGroup `json:"groups" gorm:"many2many:instance_group;"` //
}

func (Instance) TableName() string {
	return models.TableNameInstance
}

type InstanceGroup struct {
	models.BaseModel
	Name string `json:"name"` //
	Desc string `json:"desc"` //
}

func (InstanceGroup) TableName() string {
	return models.TableNameInstanceGroup
}
