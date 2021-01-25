package router

import (
	"blog/handler"

	"github.com/labstack/echo/v4"
)

// apiRouter 通用访问
// 服务几个实体 用户、分类、文章、页面、配置
func apiRouter(api *echo.Group) {
	api.GET(`/user/exist/:num`, handler.UserExist)   // 判断账号是否存在
	api.POST(`/login`, handler.UserLogin)            // 登陆
	api.GET("/vcode", handler.Vcode)                 // 验证码
	api.POST(`/logout`, handler.UserLogout)          // 注销
	api.GET(`/cate/all`, handler.CateAll)            // 分类列表
	api.GET(`/post/tag/get/:id`, handler.PostTagGet) // 通过分类查询文章
	api.GET(`/post/get/:id`, handler.PostGet)        // 文章
	api.GET(`/cate/post/:cid`, handler.CatePost)     // 通过分类查询文章
	api.GET(`/opts/:key`, handler.OptsGet)           // 获取配置项
	api.GET(`/page/all`, handler.PostPageAll)        // 页面
	api.GET(`/tag/all`, handler.TagAll)              // 标签列表
	api.GET(`/opts/base`, handler.OptsBase)          // 配置
}
