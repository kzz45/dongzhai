package monitor

import "dongzhai/models"

type Task struct {
	models.BaseModel
	ServerId  *uint      `json:"server_id"`                                 // 关联的Prometheus Server
	Server    Server     `json:"server" gorm:"foreignKey:ServerId"`         //
	Name      string     `json:"name"`                                      // 任务名称
	Interval  int        `json:"interval"`                                  // 抓取周期
	Timeout   int        `json:"timeout"`                                   // 抓取超时时间
	URL       string     `json:"url"`                                       // 任务URL
	Instances []Instance `json:"instances" gorm:"many2many:task_instance;"` // 任务下的实例
}

func (Task) TableName() string {
	return models.TableNameTask
}
