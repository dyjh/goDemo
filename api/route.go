package api

import (
	"demo/api/controller"
	"demo/api/service"
	"demo/router/middleware"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func Route(sessManager *sessions.Sessions, engine *xorm.Engine, conn redis.Conn, mvc *mvc.Application) {
	authService := service.NewAuthService(engine)
	auth := mvc.Party("/auth")
	auth.Register(
		conn,
		authService,
		sessManager.Start,
	)
	auth.Handle(new(controller.AuthController))

	userService := service.NewUserService(engine)

	//user := mvc.Party("/user", middleware.RedisHandler)     //redis方式鉴权
	user := mvc.Party("/user", middleware.JwtHandler().Serve)
	user.Register(
		conn,
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.UserController))
}
