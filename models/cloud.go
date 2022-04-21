package models

type Cloud struct {
	BaseModel
	Name      string `json:"name"`       //
	AccessKey string `json:"access_key"` //
	SecretKey string `json:"secret_key"` //
}

func (Cloud) TableName() string {
	return TableNameCloud
}
