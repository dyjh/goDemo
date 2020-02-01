package model

import (
	"demo/utils"
	"time"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type LuckyDrawActive struct {
	Id          int64     `xorm:"pk autoincr" json:"id"`        //主键 用户ID
	Title       string    `xorm:"varchar(30) notnull" json:"title" validate:"required"`  //用户名称
	RuleText    string    `xorm:"longtext" json:"rule_text" validate:"required"`
	Background  string    `xorm:"longtext" json:"background" validate:"required"`
	StartTime   time.Time `xorm:"" json:"start_time" validate:"required"`
	EndTime     time.Time `xorm:"" json:"end_time" validate:"required"`
	CreateTime  time.Time `xorm:"create" json:"create_time"`
	DeleteTime  time.Time `xorm:"deleted" json:"delete_time"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (active *LuckyDrawActive) LuckyDrawActiveToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           active.Id,
		"title":        active.Title,
		"rule":         active.RuleText,
		"start":        utils.FormatDatetime(active.StartTime),
		"end":          utils.FormatDatetime(active.EndTime),
		"background":   active.Background,
	}
	return respInfo
}
