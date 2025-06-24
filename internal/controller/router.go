package controller

import (
	"fmt"
	"time"

	"gin-weather/internal/config"
	"gin-weather/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(cfg *config.Config, weatherService service.WeatherService) *gin.Engine {
	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 创建 Gin 引擎
	router := gin.New()

	// 添加中间件
	setupMiddleware(router)

	// 创建控制器实例
	weatherController := NewWeatherController(weatherService)

	// 设置路由组
	setupRoutes(router, weatherController)

	return router
}

// setupMiddleware 设置中间件
func setupMiddleware(router *gin.Engine) {
	// 日志中间件
	router.Use(gin.Logger())

	// 恢复中间件（处理 panic）
	router.Use(gin.Recovery())

	// CORS 中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 生产环境中应该限制具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 自定义中间件：请求 ID
	router.Use(RequestIDMiddleware())

	// 自定义中间件：限流（可选）
	// router.Use(RateLimitMiddleware())
}

// setupRoutes 设置路由
func setupRoutes(router *gin.Engine, weatherController *WeatherController) {
	// API 版本 1
	v1 := router.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", weatherController.HealthCheck)

		// 天气相关路由
		weather := v1.Group("/weather")
		{
			// 通用天气查询接口（支持城市名称或坐标）
			weather.GET("", weatherController.GetWeather)

			// 根据城市名称查询天气
			weather.GET("/city/:city", weatherController.GetWeatherByCity)

			// 根据坐标查询天气
			weather.GET("/coordinates/:lat/:lon", weatherController.GetWeatherByCoordinates)
		}
	}

	// 根路径重定向到 API 文档或健康检查
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin Weather API",
			"version": "1.0.0",
			"docs":    "/api/v1/health",
		})
	})
}

// RequestIDMiddleware 请求 ID 中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	}
}

// generateRequestID 生成请求 ID
func generateRequestID() string {
	// 简单的请求 ID 生成，生产环境可以使用 UUID
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
