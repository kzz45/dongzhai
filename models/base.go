package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableNameCloud        = "cloud"
	TableNameDomain       = "domain"
	TableNameDomainCert   = "domain_cert"
	TableNameDomainRecord = "domain_record"
	TableNameTask         = "monitor_task"
	TableNameGroup        = "monitor_group"
	TableNameServer       = "monitor_server"
	TableNameProduct      = "monitor_product"
	TableNameInstance     = "monitor_instance"
	TableNameAlertRule    = "monitor_alert_rule"
	TableNameAlertRoute   = "monitor_alert_route"
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
