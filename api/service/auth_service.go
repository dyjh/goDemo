package service

import (
	"demo/model"
	"demo/router/middleware"
	"demo/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"time"
)

type AuthService interface {
	UserRegister(data model.User) string
	UserLogin(data LoginData, conn redis.Conn) utils.ReturnData
}

func NewAuthService(engine *xorm.Engine) AuthService {
	return &authService{
		Engine: engine,
	}
}

type authService struct {
	Engine *xorm.Engine
}

func (ac *authService) UserLogin(data LoginData, conn redis.Conn) utils.ReturnData {
	var user model.User
	result, _ := ac.Engine.Where(" mobile = ?", data.Mobile).Get(&user)
	fmt.Println(user)
	user.Balance = 0
	_, _ = ac.Engine.Where("id = %d", user.Id).Update(&user)
	//ac.Engine.
	_,_ = ac.Engine.Where(" mobile = ?", data.Mobile).Get(&user)
	if !result {
		return utils.ReturnData{
			Message:"账号不存在",
			Data:result,
			Code:101,
		}
	}
	fmt.Println(user)
	compareRes := utils.ComparePasswords(user.Pwd, data.Pwd)
	if compareRes != true {
		return utils.ReturnData{
			Message:"密码错误",
			Data:result,
			Code:102,
		}
	}
	token := middleware.GetJwtToken(user.Id)
	//token := middleware.GetRedisToken(user.Id)
	//_, _ = redis.String(conn.Do("SET", token, user.Id, "EX", 7200))
	return utils.ReturnData{
		Message:"success",
		Data:token,
		Code:0,
	}
}

func (ac *authService) UserRegister(data model.User) string {
	data.Balance = 0
	data.RegisterTime = time.Now()
	data.Pwd = utils.HashAndSalt(data.Pwd)
	rowNum, err := ac.Engine.Insert(&data)
	if err != nil {
		panic(err.Error())
	}
	if rowNum == 0 {
		return "注册失败！"
	}
	return "success"
}


type LoginData struct {
	Pwd     string `json:"pwd" validate:"required,min=6,max=20"`
	Mobile   string `json:"mobile" validate:"required"`
}

