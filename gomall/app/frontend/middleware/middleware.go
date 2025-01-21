package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

// 一键初始化的时候注册
func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}
