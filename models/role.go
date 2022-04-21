package models

type Role struct {
	BaseModel
	Name  string `json:"name"`                           //
	Desc  string `json:"desc"`                           //
	Users []User `json:"users" gorm:"foreignKey:RoleId"` //
}

func (Role) TableName() string {
	return TableNameRole
}
