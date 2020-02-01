package controller

import (
	"demo/api/service"
	"demo/router/middleware"
	"demo/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

type UserController struct {
	Ctx iris.Context
	AuthService service.AuthService
	Session *sessions.Session
}

func (uc *UserController) GetInfo(cxt iris.Context) mvc.Result {
	userInfo := middleware.ParseClaimsConfig(cxt)

	id := strconv.FormatInt(userInfo.Id, 10)

	return utils.Api(utils.RECODE_OK, "用户id" + id, map[string]interface{}{})
}