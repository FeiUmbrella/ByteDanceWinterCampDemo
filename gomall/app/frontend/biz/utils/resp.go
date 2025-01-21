package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"gomall/app/frontend/middleware"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// test: 向后端返回至前端的rep中添加 user_id 字段
// 看前端主页面是否隐藏sign in 按钮，显示user图标（逻辑在前端文件header.tmpl）
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// 如果能从cookie中拿到userid，就会到有user图标的主页面，否则到有sign in按钮的主页面

	content["user_id"] = ctx.Value(middleware.SessionUserId)
	return content
}
