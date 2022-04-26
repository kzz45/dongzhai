package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateSecret(secret k8s_model.Secret) error {
	var s k8s_model.Secret
	if !errors.Is(db.GlobalGorm.Where("name = ? AND cluster_id = ? AND namespace = ?", secret.Name, secret.ClusterId, secret.Namespace).
		First(&s).Error, gorm.ErrRecordNotFound) {
		return errors.New("secret exists")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 这里获取下命名空间
	var project k8s_model.Project
	if err := tx.Model(&k8s_model.Project{}).
		Where("id = ?", secret.ProjectId).
		Preload("Product").
		Find(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	secret.Namespace = project.Product.Name
	if err := tx.Create(&secret).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("name = ? AND cluster_id = ? AND namespace = ?", secret.Name, secret.ClusterId, secret.Namespace).
		Find(&s).Error; err != nil {
		tx.Rollback()
		return err
	}
	// if err := create_secret(s.ClusterId, s); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetSecrets(p *models.Pagination) ([]k8s_model.Secret, int64, error) {
	var secrets []k8s_model.Secret
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&secrets).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).
		Offset(offset).
		Preload("Project").Preload("Cluster").
		Find(&secrets).Error; err != nil {
		return nil, 0, err
	}
	return secrets, p.Total, nil
}

func UpdateSecret(secret k8s_model.Secret) error {
	return db.GlobalGorm.Where("id = ?", secret.ID).First(&secret).Updates(&secret).Error
}

func DeleteSecretById(id int) error {
	var secret k8s_model.Secret
	return db.GlobalGorm.Where("id = ?", id).Delete(&secret).Error
}
