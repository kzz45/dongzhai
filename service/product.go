package service

import (
	"dongzhai/db"
	"dongzhai/models"
	"errors"

	"gorm.io/gorm"
)

func CreateProduct(product models.Product) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", product.Name).First(&product).Error, gorm.ErrRecordNotFound) {
		return errors.New("product exist")
	}
	return db.GlobalGorm.Create(&product).Error
}

func GetProducts(p *models.Pagination) ([]models.Product, int64, error) {
	var products []models.Product
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&products).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, p.Total, nil
}

func UpdateProduct(product models.Product) error {
	return db.GlobalGorm.Where("id = ?", product.ID).First(&product).Updates(&product).Error
}

func DeleteProductById(id int) error {
	var product models.Product
	return db.GlobalGorm.Where("id = ?", id).Delete(&product).Error
}
