package router

import (
	"blog/handler"

	"github.com/labstack/echo/v4"
)

// admRouter 登录访问
func admRouter(adm *echo.Group) {
	adm.GET(`/sys`, handler.Sys)                      // 服务器信息
	adm.GET(`/collect`, handler.Collect)              // 统计信息
	adm.GET(`/auth`, handler.UserAuth)                // 获取当前登陆信息
	adm.POST(`/upload`, handler.Upload)               // 图片上传
	adm.POST(`/user/edit/self`, handler.UserEditSelf) // 修改自身信息
	adm.POST(`/user/pass`, handler.UserPass)          // 修改密码
	adm.GET(`/cate/drop/:id`, handler.CateDrop)       // 删除分类
	adm.POST(`/cate/add`, handler.CateAdd)            // 添加分类
	adm.POST(`/cate/edit`, handler.CateEdit)          // 编辑分类
	adm.POST(`/post/opts`, handler.PostOpts)          // 文章/页面-编辑/添加
	adm.GET(`/post/drop/:id`, handler.PostDrop)       // 删除文章/页面
	adm.GET(`/tag/drop/:id`, handler.TagDrop)         // 删除标签
	adm.POST(`/tag/add`, handler.TagAdd)              // 添加标签
	adm.POST(`/tag/edit`, handler.TagEdit)            // 编辑标签
	adm.POST(`/opts/edit`, handler.OptsEdit)          // 编辑配置项
}
