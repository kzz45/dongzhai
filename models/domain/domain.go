package domain

import "dongzhai/models"

// 顶级域名
type Domain struct {
	models.BaseModel
	Name          string         `json:"name"`                                    //
	Desc          string         `json:"desc"`                                    //
	DomainCerts   []DomainCert   `json:"domain_certs" gorm:"foreignKey:DomainId"` //
	DomainRecords []DomainRecord `json:"records" gorm:"foreignKey:DomainId"`      //
}

func (Domain) TableName() string {
	return models.TableNameDomain
}

// 域名解析记录
type DomainRecord struct {
	models.BaseModel
	DomainId uint   `json:"domain_id"`                         //
	Domain   Domain `json:"domain" gorm:"foreignKey:DomainId"` //
	Name     string `json:"name"`                              //
	Desc     string `json:"desc"`                              //
	Type     string `json:"type"`                              //
	Value    string `json:"value"`                             //
	TTL      int    `json:"ttl"`                               //
	Status   string `json:"status"`                            //
}

func (DomainRecord) TableName() string {
	return models.TableNameDomainRecord
}

// 域名证书
type DomainCert struct {
	models.BaseModel
	DomainId uint   `json:"domain_id"`                         //
	Domain   Domain `json:"domain" gorm:"foreignKey:DomainId"` //
	Name     string `json:"name"`                              //
	Desc     string `json:"desc"`                              //
}

func (DomainCert) TableName() string {
	return models.TableNameDomainCert
}
