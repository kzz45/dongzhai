package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateConfigMap(configmap k8s_model.ConfigMap) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", configmap.Name).First(&configmap).Error, gorm.ErrRecordNotFound) {
		return errors.New("configmap exists")
	}
	return db.GlobalGorm.Create(&configmap).Error
}

func GetConfigMaps(p *models.Pagination) ([]k8s_model.ConfigMap, int64, error) {
	var configmaps []k8s_model.ConfigMap
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&configmaps).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&configmaps).Error; err != nil {
		return nil, 0, err
	}
	return configmaps, p.Total, nil
}

func UpdateConfigMap(configmap k8s_model.ConfigMap) error {
	return db.GlobalGorm.Where("id = ?", configmap.ID).First(&k8s_model.ConfigMap{}).Updates(&configmap).Error
}

func DeleteConfigMapById(id int) error {
	var configmap k8s_model.ConfigMap
	return db.GlobalGorm.Where("id = ?", id).Delete(&configmap).Error
}
