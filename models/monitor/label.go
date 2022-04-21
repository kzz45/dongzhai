package monitor

import "dongzhai/models"

// 标签名称

type Label struct {
	models.BaseModel
	Name   string       `json:"name"`   //
	Values []LabelValue `json:"values"` //
}

// 标签值

type LabelValue struct {
	models.BaseModel
	LabelId uint   `json:"label_id"` //
	Label   Label  `json:"label"`    //
	Name    string `json:"name"`     //
}
