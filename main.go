package main

import (
	"demo/api"
	_ "demo/api/controller"
	_ "demo/api/service"
	"demo/config"
	"demo/datasource"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

func main() {
	app := newApp()

	//应用App设置
	configAction(app)

	//路由设置
	mvcHandle(app)

	initConfig := config.InitConfig()
	addr := ":" + initConfig.Port
	_ = app.Run(
		iris.Addr(addr),                               //在端口9000进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

//构建App
func newApp() *iris.Application {
	app := iris.New()
	app.Party("/")
	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		_ = context.View("index.html")
	})

	return app
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {

	//启用session
	redis := datasource.NewRedis()
	//设置session的同步位置为redis

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessionCookie",
		Expires: 24 * time.Hour,
	})
	sessManager.UseDatabase(redis)
	engine := datasource.NewPostgresEngine()
	redigoConn, _ := datasource.InitRedis()
	//TODO 实例化api模块
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://foo.com"},   //允许通过的主机名称
		AllowCredentials: true,
	})
	api.Route(sessManager, engine, redigoConn, mvc.New(app.Party("/api", crs)))

}

/**
 * 项目设置
 */
func configAction(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		_, _ = context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		_, _ = context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}
