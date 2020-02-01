package admin

import (
	"demo/admin/controller"
	"demo/admin/service"
	"demo/router/middleware"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func Route(sessManager *sessions.Sessions, engine *xorm.Engine, mvc *mvc.Application) {
	drawService := service.NewDrawService(engine)

	//user := mvc.Party("/user", middleware.RedisHandler)     //redis方式鉴权
	user := mvc.Party("/user", middleware.JwtHandler().Serve)
	user.Register(
		drawService,
		sessManager.Start,
	)
	user.Handle(new(controller.DrawController))
}
