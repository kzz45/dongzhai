package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateJob(job k8s_model.Job) error {
	var j k8s_model.Job
	if !errors.Is(db.GlobalGorm.Where("name = ? AND cluster_id = ?", job.Name, job.ClusterId).
		First(&j).Error, gorm.ErrRecordNotFound) {
		return errors.New("job exist")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 这里获取下命名空间
	var project k8s_model.Project
	if err := tx.Model(&k8s_model.Project{}).Where("id = ?", job.ProjectId).Preload("Product").
		Find(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	job.Namespace = project.Product.Name
	for _, container := range job.Containers {
		container.Namespace = project.Product.Name
	}
	// 数据库创建
	if err := tx.Create(&job).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("name = ? AND cluster_id = ? AND namespace = ?", job.Name, job.ClusterId, job.Namespace).
		Preload("Containers").
		Preload("Containers.ContainerPorts").
		Find(&j).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 关联创建Pod ?
	// var pod models.Pod
	// pod.Namespace = j.Namespace
	// pod.ClusterId = j.ClusterId
	// pod.DeploymentId = j.ID
	// if err := tx.Create(&pod).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// k8s创建

	// k8s创建
	// if err := apply_job(j.ClusterId, j); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetJobs(p *models.Pagination) ([]k8s_model.Job, int64, error) {
	var jobs []k8s_model.Job
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&jobs).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).
		Offset(offset).
		Preload("Project").Preload("Cluster").Preload("Containers").Preload("Containers.ContainerPorts").
		Find(&jobs).Error; err != nil {
		return nil, 0, err
	}
	return jobs, p.Total, nil
}
