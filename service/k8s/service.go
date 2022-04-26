package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateService(service k8s_model.Service) error {
	var svc k8s_model.Service
	if !errors.Is(db.GlobalGorm.Where("name = ?", service.Name).First(&svc).Error, gorm.ErrRecordNotFound) {
		return errors.New("service exist")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 这里获取下命名空间
	var project k8s_model.Project
	if err := tx.Model(&k8s_model.Project{}).Where("id = ?", service.ProjectId).
		Preload("Product").
		Find(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	service.Namespace = project.Product.Name
	// 数据库创建
	if err := tx.Create(&service).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("name = ? AND cluster_id = ?", service.Name, service.ClusterId).
		Preload("Deployment").Preload("ServicePorts").
		Find(&svc).Error; err != nil {
		tx.Rollback()
		return err
	}
	// k8s创建
	// if err := apply_service(svc.ClusterId, svc); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetServices(p *models.Pagination) ([]k8s_model.Service, int64, error) {
	var services []k8s_model.Service
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&services).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).
		Preload("Cluster").
		Preload("Project").
		Preload("ServicePorts").
		Find(&services).Error; err != nil {
		return nil, 0, err
	}
	return services, p.Total, nil
}

func GetServicePorts(svc_id int) ([]k8s_model.ServicePort, error) {
	var svc_ports []k8s_model.ServicePort
	if err := db.GlobalGorm.Where("service_id = ?", svc_id).Find(&svc_ports).Error; err != nil {
		return nil, err
	}
	return svc_ports, nil
}
