package monitor

import (
	"dongzhai/db"
	"dongzhai/models"
	"dongzhai/models/monitor"
	"errors"

	"gorm.io/gorm"
)

func CreateServer(server monitor.Server) error {
	if !errors.Is(db.GlobalGorm.Where("name = ?", server.Name).First(&server).Error, gorm.ErrRecordNotFound) {
		return errors.New("server already exist")
	}
	return db.GlobalGorm.Create(&server).Error
}

func GetServers(p *models.Pagination) ([]monitor.Server, int64, error) {
	var servers []monitor.Server
	if p.Size < 1 {
		p.Size = 15
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if err := db.GlobalGorm.Find(&servers).Count(&p.Total).Error; err != nil {
		return nil, 0, err
	}
	offset := p.Size * (p.Page - 1)
	if err := db.GlobalGorm.Limit(p.Size).Offset(offset).Find(&servers).Error; err != nil {
		return nil, 0, err
	}
	return servers, p.Total, nil
}

func UpdateServer(server monitor.Server) error {
	return db.GlobalGorm.Where("id = ?", server.ID).Updates(&server).Error
}

func DeleteServerById(id int) error {
	var server monitor.Server
	return db.GlobalGorm.Where("id = ?", id).First(&server).Delete(&server).Error
}
