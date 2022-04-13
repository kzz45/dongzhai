package models

// 顶级域名
type Domain struct {
	BaseModel
	Name          string         `json:"name"`                                    //
	Desc          string         `json:"desc"`                                    //
	DomainCerts   []DomainCert   `json:"domain_certs" gorm:"foreignKey:DomainId"` //
	DomainRecords []DomainRecord `json:"records" gorm:"foreignKey:DomainId"`      //
}

func (Domain) TableName() string {
	return TableNameDomain
}

// 域名解析记录
type DomainRecord struct {
	BaseModel
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
	return TableNameDomainRecord
}

// 域名证书
type DomainCert struct {
	BaseModel
	DomainId uint   `json:"domain_id"`                         //
	Domain   Domain `json:"domain" gorm:"foreignKey:DomainId"` //
	Name     string `json:"name"`                              //
	Desc     string `json:"desc"`                              //
}

func (DomainCert) TableName() string {
	return TableNameDomainCert
}
