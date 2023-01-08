package schema

import (
	"time"
)

type ActivityLog struct {
	Id          uint   `json:"id" gorm:"primaryKey;"`
	Client      string `gorm:"size:100;" json:"params_client"`
	Duration    string `gorm:"size:100;" json:"params_duration"`
	Method      string `gorm:"size:100;" json:"params_method"`
	Status      string `gorm:"size:100;" json:"params_status"`
	Path        string `gorm:"size:100;" json:"params_path"`
	ReqBody     string `gorm:"size:100;" json:"params_body"`
	PathOp      string `gorm:"size:100;" json:"params_path_op"`
	UserId      string `gorm:"size:100;" json:"params_user"`
	Source      string `gorm:"size:100;" json:"params_source"`
	Application string `gorm:"size:100;" json:"params_application"`
	Referrer    string `gorm:"size:100;" json:"params_referrer"`
	RequestId   string `gorm:"size:100;" json:"params_request_id"`
	CreatedAt   time.Time
	User        User
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
