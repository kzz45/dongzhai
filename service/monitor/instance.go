package monitor

import (
	"dongzhai/db"
	"dongzhai/models/monitor"
	"errors"

	"gorm.io/gorm"
)

func CreateInstance(instance monitor.Instance) error {
	if !errors.Is(db.GlobalGorm.Where("uuid = ?", instance.UUID).First(&instance).Error, gorm.ErrRecordNotFound) {
		return errors.New("server already exist")
	}
	return db.GlobalGorm.Create(&instance).Error
}
