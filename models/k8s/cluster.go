package k8s

import "dongzhai/models"

type Cluster struct {
	models.BaseModel
	Name       string `json:"name"`       // 集群名称
	Desc       string `json:"desc"`       // 集群描述
	Sign       string `json:"sign"`       // 集群来自 阿里云 腾讯云 等
	Status     int    `json:"status"`     // 集群状态 正常/维护
	Addr       string `json:"addr"`       // 集群地址 apiserver地址
	Token      string `json:"token"`      // token
	KubeConfig string `json:"kubeconfig"` // 集群凭证
	Endpoints  string `json:"endpoints"`  // 访问入口
}

func (Cluster) TableName() string {
	return models.TableNameCluster
}
