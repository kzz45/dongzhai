package models

type Product struct {
	BaseModel
	Name string `json:"name"` //
	Desc string `json:"desc"` //
}

func (Product) TableName() string {
	return TableNameProduct
}
