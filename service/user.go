package service

import (
	"dongzhai/db"
	"dongzhai/models"
	"dongzhai/utils"
	"errors"

	"gorm.io/gorm"
)

func CreateUser(user models.User) error {
	if !errors.Is(db.GlobalGorm.Where("username = ? AND nickname = ?", user.Username, user.Nickname).
		First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("user exist")
	}
	var user_group models.UserGroup
	var user_groups []models.UserGroup
	for _, g := range user.Groups {
		user_group.ID = g.ID
		user_groups = append(user_groups, user_group)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&user).Error; err != nil {
		tx.Callback()
		return err
	}
	if err := tx.Model(&user).Association("Groups").Append(user_groups); err != nil {
		tx.Callback()
		return err
	}
	return tx.Commit().Error
}

func GetUsers(p *models.Pagination) ([]models.User, int64, error) {
	var users []models.User
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&users).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Preload("Role").Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, p.Total, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := db.GlobalGorm.Where("id = ?", id).Preload("Role").First(&user).Error
	return user, err
}

func UpdateUser(user models.User) error {
	var user_group models.UserGroup
	var user_groups []models.UserGroup
	for _, g := range user.Groups {
		user_group.ID = g.ID
		user_groups = append(user_groups, user_group)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", user.ID).First(&user).Updates(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&user).Association("Groups").Replace(user_groups); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteUserById(id int) error {
	var user models.User
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", id).First(&user).Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&user).Association("Groups").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func UserLogin(u models.UserLogin) (models.UserLoginResp, error) {
	var user models.User
	var user_resp models.UserLoginResp
	err := db.GlobalGorm.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return user_resp, err
	}
	token, err := utils.CreateToken(user.ID, user.Username, user.Password, user.RoleId, user.Admin)
	if err != nil {
		return user_resp, err
	}
	user_resp.Id = user.ID
	user_resp.Username = user.Username
	user_resp.Nickname = user.Nickname
	user_resp.Token = token
	return user_resp, nil
}
