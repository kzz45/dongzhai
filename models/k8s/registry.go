package k8s

import "dongzhai/models"

type Registry struct {
	models.BaseModel
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	Addr     string `json:"addr" binding:"required"`
	Status   int    `json:"status"`
	Sign     string `json:"sign" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Registry) TableName() string {
	return models.TableNameRegistry
}

type HarborRepo struct {
	RepoName string `json:"repository_name"`
}

type HarborRepos struct {
	Repos []HarborRepo `json:"repository"`
}
