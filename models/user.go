package models

type User struct {
	BaseModel
	RoleId   *uint       `json:"role_id"`                             //
	Role     Role        `json:"role" gorm:"foreignKey:RoleId"`       //
	Username string      `json:"username"`                            //
	Nickname string      `json:"nickname"`                            //
	Password string      `json:"password"`                            //
	Status   bool        `json:"status"`                              //
	Admin    bool        `json:"admin"`                               //
	Groups   []UserGroup `json:"groups" gorm:"many2many:user_group;"` //
}

func (User) TableName() string {
	return TableNameUser
}

type UserGroup struct {
	BaseModel
	Name  string `json:"name"`                               //
	Desc  string `json:"desc"`                               //
	Users []User `json:"users" gorm:"many2many:user_group;"` //
}

func (UserGroup) TableName() string {
	return TableNameUserGroup
}
