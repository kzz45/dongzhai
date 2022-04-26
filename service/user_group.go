package service

import (
	"dongzhai/db"
	"dongzhai/models"
	"errors"

	"gorm.io/gorm"
)

func CreateUserGroup(group models.UserGroup) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", group.Name).First(&group).Error, gorm.ErrRecordNotFound) {
		return errors.New("user_group exists")
	}

	var user models.User
	var users []models.User
	for _, v := range group.Users {
		user.ID = uint(v.ID)
		users = append(users, user)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&group).Error; err != nil {
		tx.Callback()
		return err
	}
	if err := tx.Model(&group).Association("Users").Append(users); err != nil {
		tx.Callback()
		return err
	}
	return tx.Commit().Error
}

func GetUserGroup(p *models.Pagination) ([]models.UserGroup, int64, error) {
	var ugs []models.UserGroup
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&ugs).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	err := db.GlobalGorm.Limit(p.Size).
		Offset(offset).Preload("Users").Find(&ugs).Error
	return ugs, p.Total, err
}

func UpdateUserGroup(group models.UserGroup) error {
	var user models.User
	var users []models.User
	for _, v := range group.Users {
		user.ID = v.ID
		users = append(users, user)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", group.ID).First(&models.UserGroup{}).Updates(&group).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&group).Association("Users").Replace(users); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteUserGroup(id int) error {
	var group models.UserGroup
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", id).First(&group).Delete(&group).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&group).Association("Users").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
