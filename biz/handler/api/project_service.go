// Code generated by hertz generator.

package api

import (
	"bulebook/biz/pack"
	"bulebook/pkg/errno"
	"bulebook/pkg/logger"
	service "bulebook/service/project"
	"context"

	api "bulebook/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// Publish .
// @router book/project/publish [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishRequesnt
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}

	l := service.NewProjectService(ctx)
	if err = l.CreateProject(&req); err != nil {
		logger.Errorf("project.Publish failed, err: %v", err)
		pack.RespError(c, errno.ConvertErr(err))
		return
	}
	pack.RespSuccess(c)
}
