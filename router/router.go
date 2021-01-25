package router

import (
	"blog/handler"
	"blog/utils"
	logs "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RunApp 入口
func RunApp() {
	engine := echo.New()
	engine.Renderer = initRender()                    // 初始渲染引擎
	engine.Use(midRecover, midLogger)                 // 恢复 日志记录
	engine.Use(middleware.CORSWithConfig(crosConfig)) // 跨域设置
	engine.HideBanner = true                          // 不显示横幅
	engine.HTTPErrorHandler = HTTPErrorHandler        // 自定义错误处理
	engine.Debug = utils.Conf.IsDev()                 // 运行模式 - echo框架好像没怎么使用这个
	RegDocs(engine)                                   // 注册文档
	engine.Static(`/dist`, "dist")                    // 静态目录 - 后端专用
	engine.Static(`/static`, "static")                // 静态目录
	engine.File(`/favicon.ico`, "favicon.ico")        // ico
	engine.File("/dashboard*", "dist/index.html")     // 前后端分离页面

	//--- 页面 -- start
	engine.GET(`/`, handler.IndexView)                 // 首页
	engine.GET(`/archives`, handler.ArchivesView)      // 归档
	engine.GET(`/archives.json`, handler.ArchivesJson) // 归档 json
	engine.GET(`/tags`, handler.TagsView)              // 标签
	engine.GET(`/tags.json`, handler.TagsJson)         // 标签 json
	engine.GET(`/tag/:tag`, handler.TagPostView)       // 具体某个标签
	engine.GET(`/cate/:cate`, handler.CatePostView)    // 分类
	engine.GET(`/about`, handler.AboutView)            // 关于
	engine.GET(`/links`, handler.LinksView)            // 友链
	engine.GET(`/post/*`, handler.PostView)            // 具体某个文章
	engine.GET(`/page/*`, handler.PageView)            // 具体某个页面
	//--- 页面 -- end

	api := engine.Group("/api")         // api/
	apiRouter(api)                      // 注册分组路由
	adm := engine.Group("/adm", midJwt) // adm/ 需要登陆才能访问
	admRouter(adm)                      // 注册分组路由
	err := engine.Start(utils.Conf.App.Addr)
	if err != nil {
		logs.Fatalln("run error :", err)
	}
}
