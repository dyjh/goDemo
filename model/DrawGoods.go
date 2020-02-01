package model

import (
	"time"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type DrawGoods struct {
	Id             int64         `xorm:"pk autoincr" json:"id"`
	GoodsTable     string         `xorm:"varchar(20) null default(null)" json:"goods_table"`
	GoodsId        int64         `xorm:"int(11) null default(0)" json:"goods_id"`
	GoodsType      int8           `xorm:"int(1) default(0)" json:"goods_type"`   // 自定义 0 从某个商品表拉取 1
	Name           string         `xorm:"string(10) notnull" json:"name"`
	Thumb          string         `xorm:"longtext" json:"thumb"`
	BaseRatio      int8           `xorm:"int(3) notnull" json:"base_ratio"`
	Number         uint64          `xorm:"int(11) notnull default(0)" json:"number"`
	UserAllowTime  int64          `xorm:"int(11) default(-1)" json:"user_allow_time"` // -1  该次抽奖活动不限制单个用户中将客户
	AdditionalRule AdditionalRule `xorm:"longtext json default('{}')" json:"additional_rule"`
	CreateTime     time.Time      `xorm:"create" json:"create_time"`
	DeleteTime     time.Time      `xorm:"delete" json:"delete_time"`
}

type AdditionalRule struct {
	WinningByTotalNumber []WinningByTotalNumber `json:"winning_by_total_number"`
	WinningByOwnDrawNumber []WinningByOwnDrawNumber `json:"winning_by_own_draw_number"`
	RatioIncreaseByNoWinTimes int64 `json:"ratio_increase_by_no_win_times"`
	RatioIncreaseNumber int64 `json:"ratio_increase_number"`
}

type WinningByTotalNumber struct {
	Type     int8  `json:"type"`  //1 自定义中奖次数 2 等差计算 3 等比计算
	DiyTimes []int64 `json:"diy_times"`
	Arithmetic int64 `json:"arithmetic"`
	Geometric  int64 `json:"geometric"`
}

type WinningByOwnDrawNumber struct {
	UserId   int64 `json:"user_id"` // 0 表示应用于全部用户
	Type     int8  `json:"type"`  //1 自定义中奖次数 2 等差计算 3 等比计算
	DiyTimes []int64 `json:"diy_times"`
	Arithmetic int64 `json:"arithmetic"`
	Geometric  int64 `json:"geometric"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (goods *DrawGoods) GetGoodsInfo() interface{} {
	respInfo := map[string]interface{}{
		"id":               goods.Id,
		"name":             goods.Name,
		"thumb":            goods.Thumb,
		"base_ratio":       goods.BaseRatio,
		"number":           goods.Number,
		"additional_rule":  goods.AdditionalRule,
	}
	return respInfo
}
