package monitor

import "dongzhai/models"

type AlertRoute struct {
	models.BaseModel
	ProductId      *uint          `json:"product_id"`      //
	Product        models.Product `json:"product"`         //
	Name           string         `json:"name"`            //
	Desc           string         `json:"desc"`            //
	GroupBy        int            `json:"group_by"`        //
	GroupWait      int            `json:"group_wait"`      //
	GroupInterval  int            `json:"group_interval"`  //
	RepeatInterval int            `json:"repeat_interval"` //
	Enable         bool           `json:"enable"`          //
	Receiver       []Receiver     `json:"receiver"`        // 接收者
}

func (AlertRoute) TableName() string {
	return models.TableNameAlertRoute
}
