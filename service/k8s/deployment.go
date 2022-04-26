package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateDeployment(deployment k8s_model.Deployment) error {
	var deploy k8s_model.Deployment
	if !errors.Is(db.GlobalGorm.Where("name = ? AND cluster_id = ? AND namespace = ?", deployment.Name, deployment.ClusterId, deployment.Namespace).
		First(&deploy).Error, gorm.ErrRecordNotFound) {
		return errors.New("deployment exist")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 这里获取下命名空间
	var project k8s_model.Project
	if err := tx.Model(&k8s_model.Project{}).Where("id = ?", deployment.ProjectId).
		Preload("Product").
		Find(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	deployment.Namespace = project.Product.Name
	for _, container := range deployment.Containers {
		container.Namespace = project.Product.Name
	}
	// 数据库创建
	if err := tx.Create(&deployment).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("name = ? AND cluster_id = ? AND namespace = ?", deployment.Name, deployment.ClusterId, deployment.Namespace).
		Preload("Containers").
		Preload("Containers.ContainerPorts").
		Find(&deploy).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 关联创建Pod ?
	// var pod models.Pod
	// pod.Namespace = deploy.Namespace
	// pod.ClusterId = &deploy.ClusterId
	// pod.DeploymentId = &deploy.ID
	// if err := tx.Create(&pod).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// k8s创建
	// if err := apply_deployment(deploy.ClusterId, deploy); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetDeployment(p *models.Pagination) ([]k8s_model.Deployment, int64, error) {
	var deploys []k8s_model.Deployment
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&deploys).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).
		Offset(offset).
		Preload("Project").Preload("Cluster").Preload("Containers").Preload("Containers.ContainerPorts").
		Find(&deploys).Error; err != nil {
		return nil, 0, err
	}
	return deploys, p.Total, nil
}

func GetDeploymentWithAppId(p *models.Pagination, app_id int) ([]k8s_model.Deployment, int64, error) {
	var deploys []k8s_model.Deployment
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Model(&k8s_model.Deployment{}).
		Where("project_id = ?", app_id).
		Find(&deploys).
		Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Model(&k8s_model.Deployment{}).
		Where("project_id = ?", app_id).
		Limit(p.Size).
		Offset(offset).
		Preload("Project").Preload("Cluster").Preload("Containers").Preload("Containers.ContainerPorts").
		Find(&deploys).Error; err != nil {
		return nil, 0, err
	}
	return deploys, p.Total, nil
}

func UpdateDeployment(deployment k8s_model.Deployment) error {
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 数据库更新
	if err := tx.Where("id = ?", deployment.ID).First(&k8s_model.Deployment{}).Updates(&deployment).Error; err != nil {
		tx.Rollback()
		return err
	}
	// k8s更新
	// if err := apply_deployment(deployment.ClusterId, deployment); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}
