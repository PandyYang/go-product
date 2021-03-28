package main

import "github.com/kataras/iris"

func main() {
	//1.创建iris实例
	app := iris.New()
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	//注册模板
	app.RegisterView(template)
	//设置模板目标
	app. HandleDir("/assets","./backend/web/assets")
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message",
			context.Values().GetStringDefault("message","访问的页面出错!"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	//注册控制器 实现路由

	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
