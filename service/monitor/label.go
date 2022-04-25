package monitor

import (
	"dongzhai/db"
	"dongzhai/models"
	"dongzhai/models/monitor"
	"errors"

	"gorm.io/gorm"
)

func CreateLabelName(label monitor.Label) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", label.Name).First(&label).Error, gorm.ErrRecordNotFound) {
		return errors.New("label exists")
	}

	// var label_value monitor.LabelValue
	// var label_values []monitor.LabelValue
	// for _, v := range label.Values {
	// 	label_value.Name = v.Name
	// 	label_values = append(label_values, label_value)
	// }
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&label).Error; err != nil {
		tx.Callback()
		return err
	}
	// if err := tx.Model(&label).Association("Values").Append(label_values); err != nil {
	// 	tx.Callback()
	// 	return err
	// }
	return tx.Commit().Error
}

func CreateLabelValue(label monitor.LabelValue) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", label.Name).First(&label).Error, gorm.ErrRecordNotFound) {
		return errors.New("label_value exists")
	}
	return db.GlobalGorm.Create(&label).Error
}

func GetLabelName(p *models.Pagination) ([]monitor.Label, int64, error) {
	var labels []monitor.Label
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&labels).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&labels).Error; err != nil {
		return nil, 0, err
	}
	return labels, p.Total, nil
}

func GetLabelValue(p *models.Pagination) ([]monitor.LabelValue, int64, error) {
	var label_values []monitor.LabelValue
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&label_values).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Preload("Label").Find(&label_values).Error; err != nil {
		return nil, 0, err
	}
	return label_values, p.Total, nil
}
