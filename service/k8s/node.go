package k8s

import (
	"dongzhai/db"
	"dongzhai/models"
	k8smodel "dongzhai/models/k8s"
)

func GetClusterNodes(p *models.Pagination, cluster_id int) ([]k8smodel.Node, int64, error) {
	var nodes []k8smodel.Node
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Model(&k8smodel.Node{}).
		Where("cluster_id = ?", cluster_id).
		Find(&nodes).
		Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Model(&k8smodel.Node{}).
		Where("cluster_id = ?", cluster_id).
		Limit(p.Size).
		Offset(offset).
		Preload("Cluster").
		Find(&nodes).Error; err != nil {
		return nil, 0, err
	}
	return nodes, p.Total, nil
}
