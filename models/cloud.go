package models

type Cloud struct {
	BaseModel
	Name      string `json:"name"`       //
	Desc      string `json:"desc"`       //
	Sign      string `json:"sign"`       //
	Status    bool   `json:"status"`     //
	AccessKey string `json:"access_key"` //
	SecretKey string `json:"secret_key"` //
}

func (Cloud) TableName() string {
	return TableNameCloud
}
