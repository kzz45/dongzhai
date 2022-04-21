package service

import (
	"dongzhai/db"
	"dongzhai/models"
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
