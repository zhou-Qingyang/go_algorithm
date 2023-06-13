package system

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"column:name;` // 用户登录名
	CreateTime  time.Time // 创建时间
	UpdatedTime time.Time // 更新时间

}

func (Category) TableName() string {
	return "category"
}
