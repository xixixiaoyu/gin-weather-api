# Gin Weather API 项目实现教程

本教程将详细介绍如何从零开始构建一个基于 Go 语言和 Gin 框架的天气查询 API 代理服务。

## 目录

1. [项目概述](#项目概述)
2. [技术选型](#技术选型)
3. [项目初始化](#项目初始化)
4. [项目架构设计](#项目架构设计)
5. [配置管理实现](#配置管理实现)
6. [数据模型设计](#数据模型设计)
7. [服务层实现](#服务层实现)
8. [控制器层实现](#控制器层实现)
9. [路由和中间件](#路由和中间件)
10. [主程序入口](#主程序入口)
11. [单元测试](#单元测试)
12. [Docker 化](#docker-化)
13. [文档和工具](#文档和工具)
14. [最佳实践总结](#最佳实践总结)

## 项目概述

### 需求分析

我们要构建一个天气查询 API 代理服务，具备以下功能：

- 🌤️ 支持多种查询方式（城市名称、地理坐标）
- 🔄 作为第三方天气 API 的代理服务
- 📊 提供标准化的 JSON 响应格式
- 🛡️ 完善的参数验证和错误处理
- 🚀 高性能和可扩展性
- 📝 完整的日志记录
- 🔧 灵活的配置管理

### 架构设计原则

- **分层架构**：控制器、服务、模型分离
- **依赖注入**：便于测试和扩展
- **接口设计**：面向接口编程
- **配置外部化**：使用环境变量
- **错误处理**：统一的错误响应格式
- **可测试性**：完整的单元测试覆盖

## 技术选型

### 核心技术栈

| 技术 | 版本 | 用途 | 选择理由 |
|------|------|------|----------|
| Go | 1.21+ | 编程语言 | 高性能、并发支持、简洁语法 |
| Gin | v1.10+ | Web 框架 | 轻量级、高性能、中间件丰富 |
| OpenWeatherMap | v2.5 | 天气数据源 | 免费额度、数据准确、API 稳定 |

### 依赖包选择

```go
// 核心依赖
github.com/gin-gonic/gin           // Web 框架
github.com/gin-contrib/cors        // CORS 中间件

// 标准库
encoding/json    // JSON 处理
net/http        // HTTP 客户端
os              // 环境变量
time            // 时间处理
```

## 项目初始化

### 第一步：创建 Go 模块

```bash
# 创建项目目录
mkdir gin-weather
cd gin-weather

# 初始化 Go 模块
go mod init gin-weather
```

这一步创建了 `go.mod` 文件，定义了模块名称和 Go 版本要求。

### 第二步：创建目录结构

```bash
mkdir -p cmd/server internal/{config,controller,service,model} pkg/weather configs docs examples
```

目录结构说明：

- `cmd/server/` - 主程序入口
- `internal/` - 内部包，不对外暴露
  - `config/` - 配置管理
  - `controller/` - HTTP 控制器
  - `service/` - 业务逻辑服务
  - `model/` - 数据模型
- `pkg/` - 可对外暴露的公共包
- `configs/` - 配置文件
- `docs/` - 文档
- `examples/` - 使用示例

这种结构遵循了 Go 项目的标准布局，便于代码组织和维护。

## 项目架构设计

### 分层架构图

```
┌─────────────────┐
│   HTTP Client   │
└─────────┬───────┘
          │
┌─────────▼───────┐
│   Controller    │  ← HTTP 请求处理、参数验证
└─────────┬───────┘
          │
┌─────────▼───────┐
│    Service      │  ← 业务逻辑、第三方 API 调用
└─────────┬───────┘
          │
┌─────────▼───────┐
│ External API    │  ← OpenWeatherMap API
└─────────────────┘
```

### 数据流向

1. **请求接收**：Gin 路由接收 HTTP 请求
2. **参数验证**：Controller 验证请求参数
3. **业务处理**：Service 处理业务逻辑
4. **外部调用**：Service 调用第三方 API
5. **数据转换**：将外部 API 响应转换为标准格式
6. **响应返回**：Controller 返回标准化 JSON 响应

## 配置管理实现

### 设计思路

配置管理需要解决以下问题：

- 环境变量读取和类型转换
- 默认值设置
- 配置验证
- 结构化配置

### 实现步骤

#### 1. 定义配置结构体

```go
// internal/config/config.go
type Config struct {
    Server  ServerConfig  `json:"server"`
    Weather WeatherConfig `json:"weather"`
}

type ServerConfig struct {
    Port int    `json:"port"`
    Host string `json:"host"`
    Mode string `json:"mode"`
}

type WeatherConfig struct {
    APIKey   string `json:"api_key"`
    BaseURL  string `json:"base_url"`
    Timeout  int    `json:"timeout"`
    Provider string `json:"provider"`
}
```

**设计要点**：
- 使用嵌套结构体组织配置
- JSON 标签便于序列化
- 字段命名清晰明确

#### 2. 实现配置加载函数

```go
func Load() (*Config, error) {
    config := &Config{
        Server: ServerConfig{
            Port: getEnvAsInt("SERVER_PORT", 8080),
            Host: getEnv("SERVER_HOST", "0.0.0.0"),
            Mode: getEnv("GIN_MODE", "debug"),
        },
        Weather: WeatherConfig{
            APIKey:   getEnv("WEATHER_API_KEY", ""),
            BaseURL:  getEnv("WEATHER_BASE_URL", "https://api.openweathermap.org/data/2.5"),
            Timeout:  getEnvAsInt("WEATHER_TIMEOUT", 10),
            Provider: getEnv("WEATHER_PROVIDER", "openweathermap"),
        },
    }

    if err := config.validate(); err != nil {
        return nil, fmt.Errorf("配置验证失败: %w", err)
    }

    return config, nil
}
```

**关键特性**：
- 环境变量优先，有默认值
- 类型安全的转换
- 配置验证

#### 3. 辅助函数实现

```go
func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}
```

**设计优势**：
- 统一的环境变量处理逻辑
- 类型转换错误处理
- 代码复用

## 数据模型设计

### 设计原则

1. **标准化**：统一的响应格式
2. **完整性**：包含所有必要的天气信息
3. **扩展性**：便于添加新字段
4. **类型安全**：使用合适的 Go 类型

### 核心模型

#### 1. 请求模型

```go
type WeatherRequest struct {
    City  string  `json:"city" form:"city" binding:"required_without=Lat"`
    Lat   float64 `json:"lat" form:"lat" binding:"required_without=City"`
    Lon   float64 `json:"lon" form:"lon" binding:"required_with=Lat"`
    Units string  `json:"units" form:"units" binding:"omitempty,oneof=metric imperial standard"`
    Lang  string  `json:"lang" form:"lang" binding:"omitempty"`
}
```

**验证规则**：
- `required_without=Lat`：城市名称在没有纬度时必需
- `required_with=Lat`：经度在有纬度时必需
- `oneof`：单位只能是指定值之一

#### 2. 响应模型

```go
type WeatherResponse struct {
    Location  Location  `json:"location"`
    Current   Current   `json:"current"`
    Timestamp int64     `json:"timestamp"`
    Provider  string    `json:"provider"`
}
```

**设计考虑**：
- 嵌套结构组织相关数据
- 时间戳便于缓存管理
- 提供商信息便于追踪数据源

#### 3. 通用响应格式

```go
type APIResponse struct {
    Success bool           `json:"success"`
    Data    interface{}    `json:"data,omitempty"`
    Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
    Error   string `json:"error"`
    Code    int    `json:"code"`
    Message string `json:"message,omitempty"`
}
```

**统一响应的好处**：
- 客户端处理简化
- 错误信息标准化
- API 一致性

## 服务层实现

### 设计模式

使用**策略模式**和**接口设计**：

```go
type WeatherService interface {
    GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error)
    GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error)
}
```

**优势**：
- 便于扩展其他天气服务提供商
- 便于单元测试（可以 mock）
- 符合依赖倒置原则

### OpenWeatherMap 服务实现

#### 1. 结构体定义

```go
type OpenWeatherMapService struct {
    config *config.WeatherConfig
    client *http.Client
}

func NewOpenWeatherMapService(cfg *config.WeatherConfig) *OpenWeatherMapService {
    return &OpenWeatherMapService{
        config: cfg,
        client: &http.Client{
            Timeout: time.Duration(cfg.Timeout) * time.Second,
        },
    }
}
```

**设计要点**：
- 依赖注入配置
- 自定义 HTTP 客户端
- 超时控制

#### 2. API 调用实现

```go
func (s *OpenWeatherMapService) fetchWeather(params url.Values) (*model.WeatherResponse, error) {
    requestURL := fmt.Sprintf("%s/weather?%s", s.config.BaseURL, params.Encode())
    
    resp, err := s.client.Get(requestURL)
    if err != nil {
        return nil, fmt.Errorf("请求天气 API 失败: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("读取响应体失败: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        // 错误处理逻辑
    }

    // JSON 解析和数据转换
}
```

**错误处理策略**：
- 网络错误包装
- HTTP 状态码检查
- JSON 解析错误处理
- 错误信息本地化

#### 3. 数据转换

```go
func (s *OpenWeatherMapService) convertToStandardFormat(owm *OpenWeatherMapResponse) *model.WeatherResponse {
    // 将 OpenWeatherMap 的响应格式转换为我们的标准格式
    return &model.WeatherResponse{
        Location: model.Location{
            Name:      owm.Name,
            Country:   owm.Sys.Country,
            Latitude:  owm.Coord.Lat,
            Longitude: owm.Coord.Lon,
            Timezone:  owm.Timezone,
        },
        // ... 其他字段转换
    }
}
```

**转换的必要性**：
- 屏蔽第三方 API 的变化
- 提供一致的数据格式
- 便于切换服务提供商

## 控制器层实现

### 职责分离

控制器层专注于：
- HTTP 请求处理
- 参数验证
- 响应格式化
- 错误处理

### 实现步骤

#### 1. 控制器结构体

```go
type WeatherController struct {
    weatherService service.WeatherService
}

func NewWeatherController(weatherService service.WeatherService) *WeatherController {
    return &WeatherController{
        weatherService: weatherService,
    }
}
```

**依赖注入的好处**：
- 便于单元测试
- 降低耦合度
- 符合 SOLID 原则

#### 2. 请求处理方法

```go
func (wc *WeatherController) GetWeather(c *gin.Context) {
    var req model.WeatherRequest
    if err := c.ShouldBindQuery(&req); err != nil {
        wc.respondWithError(c, http.StatusBadRequest, "参数验证失败", err.Error())
        return
    }

    // 参数验证
    if err := wc.validateRequest(&req); err != nil {
        wc.respondWithError(c, http.StatusBadRequest, "参数验证失败", err.Error())
        return
    }

    // 业务逻辑调用
    var weatherResp *model.WeatherResponse
    var err error
    
    if req.City != "" {
        weatherResp, err = wc.weatherService.GetWeatherByCity(req.City, req.Units, req.Lang)
    } else {
        weatherResp, err = wc.weatherService.GetWeatherByCoordinates(req.Lat, req.Lon, req.Units, req.Lang)
    }

    if err != nil {
        wc.respondWithError(c, http.StatusInternalServerError, "获取天气信息失败", err.Error())
        return
    }

    wc.respondWithSuccess(c, weatherResp)
}
```

**处理流程**：
1. 参数绑定和验证
2. 业务逻辑调用
3. 错误处理
4. 响应返回

#### 3. 响应辅助方法

```go
func (wc *WeatherController) respondWithError(c *gin.Context, statusCode int, error, message string) {
    c.JSON(statusCode, model.APIResponse{
        Success: false,
        Error: &model.ErrorResponse{
            Error:   error,
            Code:    statusCode,
            Message: message,
        },
    })
}

func (wc *WeatherController) respondWithSuccess(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, model.APIResponse{
        Success: true,
        Data:    data,
    })
}
```

**统一响应的优势**：
- 减少重复代码
- 保证响应格式一致
- 便于维护

## 路由和中间件

### 路由设计

#### 1. RESTful API 设计

```go
func setupRoutes(router *gin.Engine, weatherController *WeatherController) {
    v1 := router.Group("/api/v1")
    {
        v1.GET("/health", weatherController.HealthCheck)
        
        weather := v1.Group("/weather")
        {
            weather.GET("", weatherController.GetWeather)
            weather.GET("/city/:city", weatherController.GetWeatherByCity)
            weather.GET("/coordinates/:lat/:lon", weatherController.GetWeatherByCoordinates)
        }
    }
}
```

**路由设计原则**：
- 版本化 API（/api/v1）
- 资源导向的 URL 设计
- 语义化的路径参数

#### 2. 中间件配置

```go
func setupMiddleware(router *gin.Engine) {
    // 日志中间件
    router.Use(gin.Logger())
    
    // 恢复中间件
    router.Use(gin.Recovery())
    
    // CORS 中间件
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }))
    
    // 自定义中间件
    router.Use(RequestIDMiddleware())
}
```

**中间件的作用**：
- **日志记录**：记录请求信息
- **错误恢复**：处理 panic
- **CORS 支持**：跨域请求处理
- **请求 ID**：便于请求追踪

#### 3. 自定义中间件

```go
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
```

**请求 ID 的用途**：
- 请求追踪
- 日志关联
- 问题排查

## 主程序入口

### 程序启动流程

```go
func main() {
    // 1. 加载配置
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }

    // 2. 创建服务实例
    var weatherService service.WeatherService
    switch cfg.Weather.Provider {
    case "openweathermap":
        weatherService = service.NewOpenWeatherMapService(&cfg.Weather)
    default:
        log.Fatalf("不支持的天气服务提供商: %s", cfg.Weather.Provider)
    }

    // 3. 设置路由
    router := controller.SetupRouter(cfg, weatherService)

    // 4. 创建 HTTP 服务器
    server := &http.Server{
        Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
        Handler: router,
    }

    // 5. 启动服务器
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("服务器启动失败: %v", err)
        }
    }()

    // 6. 优雅关闭
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("服务器强制关闭: %v", err)
    }
}
```

### 优雅关闭

**为什么需要优雅关闭**：
- 完成正在处理的请求
- 释放资源
- 避免数据丢失

**实现要点**：
- 监听系统信号
- 设置关闭超时
- 等待请求完成

## 单元测试

### 测试策略

1. **配置测试**：验证环境变量加载
2. **控制器测试**：HTTP 请求处理
3. **服务测试**：业务逻辑验证

### 测试实现

#### 1. 配置测试

```go
func TestLoad(t *testing.T) {
    os.Setenv("WEATHER_API_KEY", "test_api_key")
    os.Setenv("SERVER_PORT", "9090")
    defer func() {
        os.Unsetenv("WEATHER_API_KEY")
        os.Unsetenv("SERVER_PORT")
    }()

    cfg, err := Load()
    if err != nil {
        t.Fatalf("加载配置失败: %v", err)
    }

    if cfg.Weather.APIKey != "test_api_key" {
        t.Errorf("期望 API Key 为 'test_api_key'，实际为 '%s'", cfg.Weather.APIKey)
    }
}
```

#### 2. 控制器测试

```go
func TestWeatherController_GetWeatherByCity(t *testing.T) {
    gin.SetMode(gin.TestMode)
    
    mockService := &MockWeatherService{}
    controller := NewWeatherController(mockService)
    
    router := gin.New()
    router.GET("/weather/city/:city", controller.GetWeatherByCity)
    
    req, _ := http.NewRequest("GET", "/weather/city/Beijing", nil)
    w := httptest.NewRecorder()
    
    router.ServeHTTP(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("期望状态码 200，实际为 %d", w.Code)
    }
}
```

**测试要点**：
- 使用 Mock 服务
- 测试各种场景
- 验证响应格式

### Mock 服务

```go
type MockWeatherService struct{}

func (m *MockWeatherService) GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error) {
    return &model.WeatherResponse{
        Location: model.Location{Name: city},
        // ... 模拟数据
    }, nil
}
```

**Mock 的好处**：
- 隔离外部依赖
- 控制测试数据
- 提高测试速度

## Docker 化

### Dockerfile 设计

```dockerfile
# 多阶段构建
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gin-weather cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/gin-weather .
EXPOSE 8080
CMD ["./gin-weather"]
```

**多阶段构建的优势**：
- 减小镜像大小
- 提高安全性
- 分离构建和运行环境

### Docker Compose

```yaml
version: '3.8'
services:
  gin-weather:
    build: .
    ports:
      - "8080:8080"
    environment:
      - WEATHER_API_KEY=${WEATHER_API_KEY}
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--spider", "http://localhost:8080/api/v1/health"]
      interval: 30s
      timeout: 10s
      retries: 3
```

**健康检查的重要性**：
- 监控服务状态
- 自动重启故障服务
- 负载均衡器集成

## 文档和工具

### Makefile 自动化

```makefile
.PHONY: build test run docker-build

build:
	go build -o bin/gin-weather cmd/server/main.go

test:
	go test -v ./...

run:
	go run cmd/server/main.go

docker-build:
	docker build -t gin-weather .
```

**自动化的好处**：
- 统一构建流程
- 减少人为错误
- 提高开发效率

### API 文档

详细的 API 文档包括：
- 接口说明
- 参数定义
- 响应格式
- 错误代码
- 使用示例

## 最佳实践总结

### 代码组织

1. **分层架构**：清晰的职责分离
2. **依赖注入**：便于测试和扩展
3. **接口设计**：面向接口编程
4. **错误处理**：统一的错误处理策略

### 配置管理

1. **环境变量**：敏感信息外部化
2. **默认值**：合理的默认配置
3. **验证机制**：启动时配置验证
4. **类型安全**：强类型配置结构

### 错误处理

1. **错误包装**：使用 `fmt.Errorf` 包装错误
2. **错误分类**：区分不同类型的错误
3. **日志记录**：记录详细的错误信息
4. **用户友好**：返回易懂的错误消息

### 测试策略

1. **单元测试**：测试单个组件
2. **集成测试**：测试组件交互
3. **Mock 使用**：隔离外部依赖
4. **覆盖率**：保证测试覆盖率

### 部署和运维

1. **容器化**：使用 Docker 部署
2. **健康检查**：监控服务状态
3. **优雅关闭**：正确处理服务停止
4. **日志管理**：结构化日志输出

### 安全考虑

1. **输入验证**：严格的参数验证
2. **错误信息**：避免泄露敏感信息
3. **HTTPS**：生产环境使用 HTTPS
4. **限流**：防止 API 滥用

这个教程展示了如何从零开始构建一个生产级别的 Go Web 服务，涵盖了架构设计、代码实现、测试、部署等各个方面。通过学习这个项目，您可以掌握 Go 语言 Web 开发的最佳实践。
