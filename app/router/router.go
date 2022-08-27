package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
	"main/app/common/log"
	"main/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	logger := log.GetSugaredLogger()

	r := gin.Default()

	// 使用zap接收gin框架默认的日志并配置日志归档
	r.Use(middleware.ZapLogger(logger), middleware.ZapRecovery(logger, true))

	// 跨域
	r.Use(middleware.Cors()) // 直接放行全部跨域请求

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/parse", api.Parse)
		apiGroup.POST("/gen", api.Generate)
	}

	return r
}
