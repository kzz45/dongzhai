package monitor

import "dongzhai/models"

type AlertRule struct {
	models.BaseModel
	ProductId *uint          `json:"product_id"` //
	Product   models.Product `json:"product"`    //
	Name      string         `json:"name"`       // 规则名称
	Desc      string         `json:"desc"`       //
	Level     string         `json:"level"`      // 规则级别 (普通/提醒/严重)
	Type      string         `json:"type"`       // 规则类型 (机器/K8S/MySQL/Redis/)
	Interval  int            `json:"interval"`   // 持续时间
	PromQL    string         `json:"promql"`     // 告警规则PromQL语句
	OP        string         `json:"op"`         // 操作符
	Value     int            `json:"value"`      // 值
	Summary   string         `json:"summary"`    // 告警信息
	Enable    bool           `json:"enable"`     // 是否启用
}

func (AlertRule) TableName() string {
	return models.TableNameAlertRule
}
