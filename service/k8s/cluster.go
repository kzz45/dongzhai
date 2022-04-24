package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	"errors"

	"gorm.io/gorm"
)

func CreateCluster(cluster k8s_model.Cluster) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", cluster.Name).First(&cluster).Error, gorm.ErrRecordNotFound) {
		return errors.New("cluster exist")
	}
	return db.GlobalGorm.Create(&cluster).Error
}

func GetClusters(p *models.Pagination) ([]k8s_model.Cluster, int64, error) {
	var clusters []k8s_model.Cluster
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&clusters).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&clusters).Error; err != nil {
		return nil, 0, err
	}
	return clusters, p.Total, nil
}

func UpdateCluster(cluster k8s_model.Cluster) error {
	return db.GlobalGorm.Where("id = ?", cluster.ID).First(&k8s_model.Cluster{}).Updates(&cluster).Error
}

func DeleteClusterById(id int) error {
	var cluster k8s_model.Cluster
	return db.GlobalGorm.Where("id = ?", id).Delete(&cluster).Error
}
