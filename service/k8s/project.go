package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	"dongzhai/models/k8s"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateProject(project k8s.Project) error {
	var p k8s.Project
	if !errors.Is(db.GlobalGorm.Where("name = ?", project.Name).First(&project).Error, gorm.ErrRecordNotFound) {
		return errors.New("project exist")
	}
	tx := db.GlobalGorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&project).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("product_id = ?", project.ProductId).Preload("Product").Find(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	logrus.Infof("create namespace: %s", p.Product.Name)
	// if err := CreateNamespace(p.ClusterId, p); err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	return tx.Commit().Error
}

func GetProjects(p *models.Pagination) ([]k8s.Project, int64, error) {
	var projects []k8s.Project
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&projects).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	// var groups []models.Group
	// err := db.GORM.Preload("Users").Where("id IN (SELECT group_id FROM user_group WHERE user_id = ?)", 2).Find(&groups).Error
	// logrus.Errorf("%v", err)
	// logrus.Errorln(groups)
	// group_ids := make([]uint, 0)
	// for _, group := range groups {
	// 	group_ids = append(group_ids, group.ID)
	// }

	// var xxx []models.Project
	// err = db.GORM.Preload("Product").Preload("Cluster").Where("user_group_id IN ?", group_ids).Find(&xxx).Error
	// logrus.Errorf("%v", err)
	// logrus.Errorln(xxx)

	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).
		Preload("Product").
		Preload("Cluster").
		Preload("UserGroup").
		Find(&projects).Error; err != nil {
		return nil, 0, err
	}
	return projects, p.Total, nil
}

func UpdateProject(project k8s.Project) error {
	return db.GlobalGorm.Where("id = ?", project.ID).First(&project).Updates(&project).Error
}

func DeleteProjectById(id int) error {
	var project k8s.Project
	return db.GlobalGorm.Where("id = ?", id).Delete(&project).Error
}
