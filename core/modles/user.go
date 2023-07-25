package modles

import "time"

type UserBasic struct {
	Id          int
	Identity    string
	Name        string
	Password    string
	Email       string
	NowVolume   int64     `xorm:"now_volume"`
	TotalVolume int64     `xorm:"total_volume"`
	CreatedAt   time.Time `xorm:"create_at"`
	UpdatedAt   time.Time `xorm:"update_at"`
	DeletedAt   time.Time `xorm:"delete_at"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
