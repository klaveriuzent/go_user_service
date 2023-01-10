package schema

import (
	"time"
)

type ActivityLog struct {
	Id          uint   `json:"id" gorm:"primaryKey;"`
	Client      string `gorm:"size:100;" json:"client"`
	Duration    string `gorm:"size:100;" json:"duration"`
	Method      string `gorm:"size:100;" json:"method"`
	Status      string `gorm:"size:100;" json:"status"`
	Path        string `gorm:"size:100;" json:"path"`
	ReqBody     string `gorm:"size:100;" json:"body"`
	PathOp      string `gorm:"size:100;" json:"path_op"`
	UserId      string `gorm:"size:100;" json:"user"`
	Source      string `gorm:"size:100;" json:"source"`
	Application string `gorm:"size:100;" json:"application"`
	Referrer    string `gorm:"size:100;" json:"referrer"`
	RequestId   string `gorm:"size:100;" json:"request_id"`
	CreatedAt   time.Time
	User        User
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
