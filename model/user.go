package model

import (
	"demo/utils"
	"time"
)

/**
 * 用户信息结构体,用于生成用户信息表
 */
type User struct {
	Id          int64     `xorm:"pk autoincr" json:"id"`        //主键 用户ID
	UserName    string    `xorm:"varchar(12)" json:"username" validate:"required"`  //用户名称
	RegisterTime time.Time `xorm:"" json:"register_time"`         //用户注册时间
	Mobile      string    `xorm:"varchar(11) unique" json:"mobile" validate:"required"`    //用户的移动手机号
	IsActive    int8      `json:"is_active"`                    //用户是否激活
	Balance     int64     `json:"balance"`                      //用户的账户余额（简单起见，使用int类型）
	Avatar      string    `xorm:"varchar(255)" json:"avatar"`   //用户的头像
	Pwd         string    `json:"password" validate:"required,min=6,max=20"`                     //用户的账户密码
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           user.Id,
		"user_id":      user.Id,
		"username":     user.UserName,
		"register_time": utils.FormatDatetime(user.RegisterTime),
		"mobile":       user.Mobile,
		"is_active":    user.IsActive,
		"balance":      user.Balance,
		"avatar":       user.Avatar,
	}
	return respInfo
}
