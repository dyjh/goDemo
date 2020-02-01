package model

import "time"

type DrawUserLog struct {
	Id       int64   `xorm:"pk autoincr" json:"id"`
	UserId   int64   `xorm:"int(11)" json:"user_id"`
	ActiveId int64   `xorm:"int(11)" json:"active_id"`
	GoodsId  int64   `xorm:"int(11)" json:"goods_id"`
	CreateTime time.Time   `xorm:"create" json:"create_time"`
}