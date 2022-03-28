package models

type Cluster struct {
	BaseModel
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Sign       string `json:"sign"`
	Addr       string `json:"addr"`
	Token      string `json:"token"`
	Endpoint   string `json:"endpoint"`
	KubeConfig string `json:"kubeconfig"`
}

func (Cluster) TableName() string {
	return "cluster"
}
