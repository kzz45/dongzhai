package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableNameDomain       = "domain"
	TableNameDomainCert   = "domain_cert"
	TableNameDomainRecord = "domain_record"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`    //
	CreatedAt time.Time      `json:"created_at"`              //
	UpdatedAt time.Time      `json:"updated_at"`              //
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` //
}
