package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateIngress(ingress k8s_model.Ingress) error {
	var ing k8s_model.Ingress
	if !errors.Is(db.GlobalGorm.Where("name = ?", ingress.Name).First(&ing).Error, gorm.ErrRecordNotFound) {
		return errors.New("ingress exists")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 这里获取下命名空间
	var project k8s_model.Project
	if err := tx.Model(&k8s_model.Project{}).Where("id = ?", ingress.ProjectId).
		Preload("Product").
		Find(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	ingress.Namespace = project.Product.Name
	// 数据库创建
	if err := tx.Create(&ingress).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("name = ? AND cluster_id = ?", ingress.Name, ingress.ClusterId).
		Preload("Rules").
		Preload("Rules.Service").
		Preload("Rules.ServicePort").
		Find(&ing).Error; err != nil {
		tx.Rollback()
		return err
	}

	// k8s创建
	// if err := apply_ingress(ing.ClusterId, ing); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetIngresses(p *models.Pagination) ([]k8s_model.Ingress, int64, error) {
	var ingresses []k8s_model.Ingress
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&ingresses).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).
		Preload("Cluster").
		Preload("Project").
		Preload("Rules").
		Preload("Rules.Service").
		Preload("Rules.ServicePort").
		Find(&ingresses).Error; err != nil {
		return nil, 0, err
	}
	return ingresses, p.Total, nil
}

func UpdateIngress(ingress k8s_model.Ingress) error {
	return db.GlobalGorm.Where("id = ?", ingress.ID).First(&k8s_model.Ingress{}).Updates(&ingress).Error
}

func DeleteIngressById(id int) error {
	var ingress k8s_model.Ingress
	return db.GlobalGorm.Where("id = ?", id).Delete(&ingress).Error
}
