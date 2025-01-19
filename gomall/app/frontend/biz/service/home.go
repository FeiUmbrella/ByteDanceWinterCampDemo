package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "gomall/app/frontend/hertz_gen/frontend/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	Item := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/T-shirt-2.jpg"},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/T-shirt-3.jpg"},

		{"Name": "Sticker-1", "Price": 130, "Picture": "/static/image/Sticker-1.jpg"},
		{"Name": "Sticker-2", "Price": 140, "Picture": "/static/image/Sticker-2.jpg"},
		{"Name": "Sticker-3", "Price": 150, "Picture": "/static/image/Sticker-3.jpg"},
	}

	rep := make(map[string]any)
	rep["Items"] = Item
	rep["Title"] = "Hot Sales"

	return rep, nil
}
