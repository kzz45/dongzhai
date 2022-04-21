package service

import (
	"dongzhai/db"
	"dongzhai/models"
	"errors"

	"gorm.io/gorm"
)

func CreateCloud(cloud models.Cloud) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", cloud.Name).First(&cloud).Error, gorm.ErrRecordNotFound) {
		return errors.New("cloud factory exist")
	}
	return db.GlobalGorm.Create(&cloud).Error
}

func GetClouds(p *models.Pagination) ([]models.Cloud, int64, error) {
	var clouds []models.Cloud
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&clouds).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&clouds).Error; err != nil {
		return nil, 0, err
	}
	return clouds, p.Total, nil
}

func UpdateCloud(cloud models.Cloud) error {
	return db.GlobalGorm.Where("id = ?", cloud.ID).First(&cloud).Updates(&cloud).Error
}

func DeleteCloudById(id int) error {
	var cloud models.Cloud
	return db.GlobalGorm.Where("id = ?", id).Delete(&cloud).Error
}
