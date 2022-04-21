package monitor

import "dongzhai/models"

// 监控实例

type Instance struct {
	models.BaseModel
	ServerId        *uint          `json:"server_id"`     // 关联的Prometheus Server
	Server          Server         `json:"server"`        //
	ProductId       *uint          `json:"product_id"`    // 关联的产品
	Product         models.Product `json:"product"`       //
	InstanceGroupId *uint          `json:"group_id"`      // 关联的分组
	InstanceGroup   InstanceGroup  `json:"group"`         //
	Name            string         `json:"name"`          //
	UUID            string         `json:"uuid"`          //
	Type            string         `json:"type"`          // 实例类型
	Endpoint        string         `json:"endpoint"`      //
	Enable          bool           `json:"enable"`        // 是否监控
	PublicIP        string         `json:"public_ip"`     //
	PrivateIP       string         `json:"private_ip"`    //
	UsePublicIP     bool           `json:"use_public_ip"` // 是否使用外网IP监控
	Labels          []Label        `json:"labels"`        // 实例上的标签
}

func (Instance) TableName() string {
	return models.TableNameInstance
}

type InstanceType struct {
	models.BaseModel
	Name string `json:"name"` //
}

type InstanceGroup struct {
	models.BaseModel
	Name string `json:"name"` //
	Desc string `json:"desc"` //
}
