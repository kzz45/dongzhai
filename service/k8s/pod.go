package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
)

func GetPods(p *models.Pagination) ([]k8s_model.Pod, int64, error) {
	var pods []k8s_model.Pod
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&pods).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Preload("Cluster").Preload("Deployment").
		Find(&pods).Error; err != nil {
		return nil, 0, err
	}
	return pods, p.Total, nil
}
