package monitor

import "dongzhai/models"

type Receiver struct {
	models.BaseModel
	Name    string `json:"name"`    //
	Desc    string `json:"desc"`    //
	Channel string `json:"channel"` //
	WebHook string `json:"webhook"` //
}
