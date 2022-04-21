package models

type User struct {
	BaseModel
	RoleId   *uint       `json:"role_id"`                             //
	Role     Role        `json:"role" gorm:"foreignKey:RoleId;"`      // 角色
	Username string      `json:"username"`                            // 用户名
	Nickname string      `json:"nickname"`                            // 别名
	Password string      `json:"password"`                            // 密码
	Status   bool        `json:"status" gorm:"default: true;"`        // 状态
	Admin    bool        `json:"admin" gorm:"default: false;"`        // 管理员
	Groups   []UserGroup `json:"groups" gorm:"many2many:user_group;"` // 用户分组
}

func (User) TableName() string {
	return TableNameUser
}

type UserGroup struct {
	BaseModel
	Name  string `json:"name"`                               // 组名
	Desc  string `json:"desc"`                               // 描述
	Users []User `json:"users" gorm:"many2many:user_group;"` // 组员
}

func (UserGroup) TableName() string {
	return TableNameUserGroup
}
