package models

import (
	"time"
)

type MediaAppAdsenseShield struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('主键') INT"`
	Name123Abc string `json:"name_123_abc" xorm:"name_123_abc comment('屏蔽名称') VARCHAR(64)"`

	CreateAt time.Time `json:"create_at" xorm:"comment('创建时间1') DATETIME"`

	Version    int       `json:"version" xorm:"version comment('乐观锁') INT"`
	CreateTime time.Time `json:"create_time" xorm:"created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"updated comment('修改时间') DATETIME"`
	DeleteTime time.Time `json:"delete_time" xorm:"deleted comment('删除时间') DATETIME"`
}
