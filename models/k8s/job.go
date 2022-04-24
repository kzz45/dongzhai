package k8s

import "dongzhai/models"

type Job struct {
	models.BaseModel
	ClusterId     uint        `json:"cluster_id"`                         // 所属集群
	Cluster       Cluster     `json:"cluster"`                            //
	ProjectId     uint        `json:"project_id"`                         // 所属项目
	Project       Project     `json:"project"`                            //
	Name          string      `json:"name"`                               // 任务名称
	Namespace     string      `json:"namespace"`                          // 任务命名空间
	Type          string      `json:"type"`                               // job/cronjob
	RetryNum      int32       `json:"retry_num"`                          // 重试次数(将任务标记为失败前的最大重试次数)
	Parallel      int32       `json:"parallel"`                           // 并行数量
	TimeOut       int64       `json:"timeout"`                            // 最大运行时间
	Completed     int32       `json:"completed"`                          // 完成数量(将任务标记为完成所需成功运行的容器组数量)
	Schedule      string      `json:"schedule"`                           // crontab定时规则
	Desire        int32       `json:"desire"`                             // 期望的数量
	Status        string      `json:"status"`                             // 任务状态
	Concurrency   string      `json:"concurrency"`                        // 并发策略 Allow/Forbid/Replace
	SuccessNum    int32       `json:"success_num"`                        // 成功任务保留数量
	FailedNum     int32       `json:"failed_num"`                         // 失败任务保留数量
	RestartPolicy string      `json:"restart_policy"`                     // 重启策略Never/OnFailure
	Labels        MapString   `json:"labels"`                             // 标签
	Annotation    MapString   `json:"annotations"`                        // 注解
	Containers    []Container `json:"containers" gorm:"foreignKey:JobId"` // 包含的容器
}

func (Job) TableName() string {
	return models.TableNameJob
}
