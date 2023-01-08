package schema

import (
	"time"
)

type ActivityLog struct {
	Id        uint   `json:"id" gorm:"primaryKey;"`
	ClientIp  string `gorm:"size:50;" json:"client_ip"`
	Duration  string `gorm:"size:10;" json:"duration"`
	Method    string `gorm:"size:100;" json:"method"`
	Path      string `gorm:"type:text" json:"path"`
	Status    string `gorm:"size:5;" json:"status"`
	UserId    string `json:"user_id"`
	Referrer  string `json:"referrer"`
	RequestId string `json:"request_id"`
	CreatedAt time.Time
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
