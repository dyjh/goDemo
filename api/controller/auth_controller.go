package controller

import (
	"demo/api/service"
	"demo/model"
	"demo/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type AuthController struct {
	RedisConn redis.Conn
	Ctx iris.Context
	AuthService service.AuthService
	Session *sessions.Session
}



func (ac *AuthController) PostRegister(context iris.Context) mvc.Result {
	var registerData model.User
	if err := ac.Ctx.ReadForm(&registerData); err != nil {
		panic(err.Error())
	} else {
		errString := utils.ValidateData(registerData)
		if errString != "success" {
			return utils.Api(utils.RECODE_OK, errString, map[string]interface{}{})
		}
	}
	fmt.Println(registerData)
	result := ac.AuthService.UserRegister(registerData)
	if result != "success" {
		return utils.Api(utils.RECODE_FAIL, result, map[string]interface{}{})
	} else {
		return utils.Api(utils.RECODE_OK, result, map[string]interface{}{})
	}
}

func (ac *AuthController) PostLogin(context iris.Context) mvc.Result {
	var loginData service.LoginData
	if err := ac.Ctx.ReadForm(&loginData); err != nil {
		panic(err.Error())
	} else {
		errString := utils.ValidateData(loginData)
		if errString != "success" {
			return utils.Api(utils.RECODE_OK, errString, map[string]interface{}{})
		}
	}
	result := ac.AuthService.UserLogin(loginData, ac.RedisConn)
	fmt.Println(result.Message)
	if result.Message == "success" {
		return utils.Api(utils.RECODE_OK, result.Message, map[string]interface{}{"token":result.Data})
	}
	return utils.Api(utils.RECODE_FAIL, result.Message, map[string]interface{}{})
}


/*func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(registerData)
}*/

