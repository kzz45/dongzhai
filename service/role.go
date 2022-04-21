package service

import (
	"dongzhai/db"
	"dongzhai/models"
	"errors"

	"gorm.io/gorm"
)

func CreateRole(role models.Role) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", role.Name).
		First(&role).Error, gorm.ErrRecordNotFound) {
		return errors.New("role exist")
	}
	var user models.User
	var users []models.User
	for _, g := range role.Users {
		user.ID = g.ID
		users = append(users, user)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&role).Error; err != nil {
		tx.Callback()
		return err
	}
	if err := tx.Model(&role).Association("Users").Append(users); err != nil {
		tx.Callback()
		return err
	}
	return tx.Commit().Error
}

func GetRoles(p *models.Pagination) ([]models.Role, int64, error) {
	var roles []models.Role
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&roles).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	return roles, p.Total, nil
}

func UpdateRole(role models.Role) error {
	var user models.UserGroup
	var users []models.UserGroup
	for _, g := range role.Users {
		user.ID = g.ID
		users = append(users, user)
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", role.ID).First(&role).Updates(&role).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&role).Association("Users").Replace(users); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteRoleById(id int) error {
	var role models.Role
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", id).First(&role).Delete(&role).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&role).Association("Users").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
