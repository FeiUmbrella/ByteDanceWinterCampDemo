package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth1 "gomall/app/frontend/hertz_gen/frontend/auth"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth1.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO: USER SVC API 用户登录验证

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	if err := session.Save(); err != nil {
		return "", err
	}
	redirect = "/" // 重定向的网页
	if req.Next != "" {
		redirect = req.Next
	}
	return req.Next, err
}
