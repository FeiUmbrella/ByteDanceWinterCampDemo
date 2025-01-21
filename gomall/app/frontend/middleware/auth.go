package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "userId"

// 处理方法用于获取相关用户身份信息，返回值为一个函数
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从session中获取用户信息放到ctx中
		// 以后其他业务逻辑需要用户信息就可以从context中取
		// 不用再利用session函数从session中取
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))

		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")

		// 获取不到用户信息则跳转到登录页面
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
