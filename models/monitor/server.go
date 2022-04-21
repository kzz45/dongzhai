package monitor

import "dongzhai/models"

// Prometheus Server

type Server struct {
	models.BaseModel
	Name      string     `json:"name"`      //
	Desc      string     `json:"desc"`      //
	UUID      string     `json:"uuid"`      // 唯一ID
	Sign      string     `json:"sign"`      // 标识
	Addr      string     `json:"addr"`      // Prometheus Server地址
	Tasks     []Task     `json:"tasks"`     //
	Instances []Instance `json:"instances"` //
}

func (Server) TableName() string {
	return models.TableNameServer
}
