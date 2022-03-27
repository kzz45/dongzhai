package models

type Deployment struct {
	BaseModel
	App        string     `json:"app"`
	Name       string     `json:"name"`
	Namespace  string     `json:"namespace"`
	Available  int32      `json:"available"`
	Desire     int32      `json:"desire"`
	Status     string     `json:"status"`
	Annotation Annotation `json:"annotations"`
}
