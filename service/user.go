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
	return db.GlobalGorm.Create(&user).Error
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
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, p.Total, nil
}
