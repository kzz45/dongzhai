package models

// 顶级域名
type Domain struct {
	BaseModel
	Name        string       `json:"name"`                                    //
	Desc        string       `json:"desc"`                                    //
	Records     []Record     `json:"records" gorm:"foreignKey:DomainId"`      //
	DomainCerts []DomainCert `json:"domain_certs" gorm:"foreignKey:DomainId"` //
}

// 域名解析记录
type Record struct {
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

// 域名证书
type DomainCert struct {
	BaseModel
	DomainId uint   `json:"domain_id"`                         //
	Domain   Domain `json:"domain" gorm:"foreignKey:DomainId"` //
	Name     string `json:"name"`                              //
	Desc     string `json:"desc"`                              //
}
