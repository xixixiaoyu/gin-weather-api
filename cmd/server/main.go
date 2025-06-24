package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-weather/internal/config"
	"gin-weather/internal/controller"
	"gin-weather/internal/service"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建天气服务实例
	var weatherService service.WeatherService
	switch cfg.Weather.Provider {
	case "openweathermap":
		weatherService = service.NewOpenWeatherMapService(&cfg.Weather)
	default:
		log.Fatalf("不支持的天气服务提供商: %s", cfg.Weather.Provider)
	}

	// 设置路由
	router := controller.SetupRouter(cfg, weatherService)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	// 启动服务器的 goroutine
	go func() {
		log.Printf("服务器启动在 %s:%d", cfg.Server.Host, cfg.Server.Port)
		log.Printf("天气服务提供商: %s", cfg.Weather.Provider)
		log.Printf("运行模式: %s", cfg.Server.Mode)
		log.Printf("API 文档: http://%s:%d/api/v1/health", cfg.Server.Host, cfg.Server.Port)
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v", err)
	}

	log.Println("服务器已关闭")
}
