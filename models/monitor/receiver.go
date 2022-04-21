package monitor

import "dongzhai/models"

type Receiver struct {
	models.BaseModel
	AlertRouteId *uint      `json:"alert_route_id"`                             //
	AlertRoute   AlertRoute `json:"alert_route" gorm:"foreignKey:AlertRouteId"` //
	Name         string     `json:"name"`                                       //
	Desc         string     `json:"desc"`                                       //
	Channel      string     `json:"channel"`                                    //
	WebHook      string     `json:"webhook"`                                    //
}

func (Receiver) TableName() string {
	return models.TableNameReceiver
}
