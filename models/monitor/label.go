package monitor

import "dongzhai/models"

// 标签名称
type Label struct {
	models.BaseModel
	Name   string       `json:"name"`                             // 标签名称
	Values []LabelValue `json:"values" gorm:"foreignKey:LabelId"` // 标签值
}

func (Label) TableName() string {
	return models.TableNameLabel
}

// 标签值
type LabelValue struct {
	models.BaseModel
	LabelId uint   `json:"label_id" `                       //
	Label   Label  `json:"label" gorm:"foreignKey:LabelId"` //
	Name    string `json:"name"`                            // 标签值Value
}

func (LabelValue) TableName() string {
	return models.TableNameLabelValue
}
