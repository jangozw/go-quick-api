package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-quick-api/erron"
	"github.com/go-quick-api/pkg/app"
	"github.com/go-quick-api/service"
	"github.com/go-quick-api/types"
)

func CheckBaseRequest(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		app.TestInstance()
		c.Next()
	}
}

func CheckTaskRequest(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		taskType := c.GetHeader(types.TaskTypeKey)
		allowTasks := ctx.Value(types.KeyAllowTasks)
		var allow bool
		if allowTasks != nil {
			if mp, ok := allowTasks.(map[string]bool); ok {
				if v, ok := mp[taskType]; ok && v == true {
					allow = true
				}
			}
		}
		if !allow {
			app.AbortJSON(c, app.ResponseFail(erron.New(erron.ErrRequestParam, "invalid task_type: "+taskType)))
			return
		}
		_, err := service.GetTaskHandler(taskType)
		if err != nil {
			app.AbortJSON(c, app.ResponseFail(erron.New(erron.ErrRequestParam, "couldn't match handlers for task_type: "+taskType)))
			return
		}
		// 继续下一步
		c.Next()
	}
}
