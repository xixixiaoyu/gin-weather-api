# Gin Weather API 架构设计文档

## 概述

本文档详细描述了 Gin Weather API 项目的架构设计，包括系统架构、模块设计、数据流向、设计模式等内容。

## 系统架构

### 整体架构图

```
┌─────────────────────────────────────────────────────────────┐
│                        Client Layer                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Web Browser │  │ Mobile App  │  │ Third-party Service │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────┬───────────────────────────────────────┘
                      │ HTTP/HTTPS
┌─────────────────────▼───────────────────────────────────────┐
│                   API Gateway Layer                         │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │              Gin Web Framework                          │ │
│  │  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐ │ │
│  │  │ CORS        │ │ Logging     │ │ Request ID          │ │ │
│  │  │ Middleware  │ │ Middleware  │ │ Middleware          │ │ │
│  │  └─────────────┘ └─────────────┘ └─────────────────────┘ │ │
│  └─────────────────────────────────────────────────────────┘ │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                 Application Layer                           │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │                Controller Layer                         │ │
│  │  ┌─────────────────┐  ┌─────────────────────────────────┐ │ │
│  │  │ Weather         │  │ Health Check                    │ │ │
│  │  │ Controller      │  │ Controller                      │ │ │
│  │  └─────────────────┘  └─────────────────────────────────┘ │ │
│  └─────────────────────┬───────────────────────────────────┘ │
│                        │                                     │
│  ┌─────────────────────▼───────────────────────────────────┐ │
│  │                Service Layer                            │ │
│  │  ┌─────────────────┐  ┌─────────────────────────────────┐ │ │
│  │  │ Weather Service │  │ Configuration                   │ │ │
│  │  │ Interface       │  │ Service                         │ │ │
│  │  └─────────────────┘  └─────────────────────────────────┘ │ │
│  │  ┌─────────────────┐                                     │ │
│  │  │ OpenWeatherMap  │                                     │ │
│  │  │ Service Impl    │                                     │ │
│  │  └─────────────────┘                                     │ │
│  └─────────────────────┬───────────────────────────────────┘ │
└─────────────────────────┼───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────────┐
│                   External Services                             │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐  │
│  │ OpenWeatherMap  │  │ Other Weather   │  │ Future Services │  │
│  │ API             │  │ APIs            │  │                 │  │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

### 架构特点

1. **分层架构**：清晰的层次划分，职责分离
2. **微服务友好**：易于拆分为微服务
3. **可扩展性**：支持多种天气服务提供商
4. **高可用性**：无状态设计，易于水平扩展

## 模块设计

### 1. 配置管理模块 (Config)

```go
// 配置结构
type Config struct {
    Server  ServerConfig
    Weather WeatherConfig
}

// 职责
- 环境变量读取
- 配置验证
- 默认值管理
- 类型转换
```

**设计原则**：
- 配置外部化
- 类型安全
- 验证机制
- 默认值合理

### 2. 数据模型模块 (Model)

```go
// 核心模型
- WeatherRequest    // 请求模型
- WeatherResponse   // 响应模型
- APIResponse       // 通用响应
- ErrorResponse     // 错误响应
```

**设计原则**：
- 数据标准化
- 类型安全
- JSON 友好
- 验证支持

### 3. 服务层模块 (Service)

```go
// 接口设计
type WeatherService interface {
    GetWeatherByCity(city, units, lang string) (*WeatherResponse, error)
    GetWeatherByCoordinates(lat, lon float64, units, lang string) (*WeatherResponse, error)
}

// 实现
- OpenWeatherMapService  // OpenWeatherMap 实现
- 未来可扩展其他服务商
```

**设计模式**：
- 策略模式
- 依赖注入
- 接口隔离
- 工厂模式

### 4. 控制器模块 (Controller)

```go
// 控制器职责
- HTTP 请求处理
- 参数验证
- 响应格式化
- 错误处理
```

**设计原则**：
- 单一职责
- 依赖注入
- 统一响应
- 错误处理

### 5. 路由模块 (Router)

```go
// 路由组织
/api/v1/
├── health              // 健康检查
└── weather/
    ├── ""              // 通用查询
    ├── city/:city      // 城市查询
    └── coordinates/:lat/:lon  // 坐标查询
```

**设计特点**：
- RESTful 设计
- 版本化 API
- 中间件支持
- 语义化路径

## 数据流向

### 请求处理流程

```
1. HTTP Request
   ↓
2. Gin Router
   ↓
3. Middleware Chain
   ├── CORS Middleware
   ├── Logging Middleware
   ├── Recovery Middleware
   └── Request ID Middleware
   ↓
4. Controller
   ├── Parameter Binding
   ├── Validation
   └── Business Logic Call
   ↓
5. Service Layer
   ├── Business Logic
   ├── External API Call
   └── Data Transformation
   ↓
6. External API
   ├── OpenWeatherMap API
   └── HTTP Request/Response
   ↓
7. Response Processing
   ├── Data Conversion
   ├── Error Handling
   └── Response Formatting
   ↓
8. HTTP Response
```

### 错误处理流程

```
Error Occurrence
   ↓
Error Classification
   ├── Validation Error (400)
   ├── External API Error (500)
   ├── Configuration Error (500)
   └── Unknown Error (500)
   ↓
Error Logging
   ├── Request ID
   ├── Error Details
   └── Stack Trace
   ↓
Error Response
   ├── Standard Format
   ├── User-friendly Message
   └── Error Code
```

## 设计模式

### 1. 依赖注入模式

```go
// 构造函数注入
func NewWeatherController(weatherService service.WeatherService) *WeatherController {
    return &WeatherController{
        weatherService: weatherService,
    }
}
```

**优势**：
- 降低耦合度
- 便于单元测试
- 提高可维护性

### 2. 策略模式

```go
// 天气服务策略
type WeatherService interface {
    GetWeatherByCity(city, units, lang string) (*WeatherResponse, error)
}

// 不同实现
- OpenWeatherMapService
- AccuWeatherService (未来扩展)
- WeatherAPIService (未来扩展)
```

**优势**：
- 算法可替换
- 易于扩展
- 符合开闭原则

### 3. 工厂模式

```go
// 服务工厂
func CreateWeatherService(provider string, config *WeatherConfig) WeatherService {
    switch provider {
    case "openweathermap":
        return NewOpenWeatherMapService(config)
    case "accuweather":
        return NewAccuWeatherService(config)
    default:
        panic("unsupported provider")
    }
}
```

**优势**：
- 创建逻辑集中
- 易于管理
- 支持配置化

### 4. 适配器模式

```go
// 将第三方 API 响应适配为标准格式
func (s *OpenWeatherMapService) convertToStandardFormat(owm *OpenWeatherMapResponse) *WeatherResponse {
    // 数据转换逻辑
}
```

**优势**：
- 接口统一
- 屏蔽差异
- 易于切换

## 安全设计

### 1. 输入验证

```go
// 参数验证
type WeatherRequest struct {
    City  string  `binding:"required_without=Lat"`
    Lat   float64 `binding:"required_without=City,min=-90,max=90"`
    Lon   float64 `binding:"required_with=Lat,min=-180,max=180"`
    Units string  `binding:"omitempty,oneof=metric imperial standard"`
}
```

### 2. 错误信息安全

```go
// 避免泄露敏感信息
func (wc *WeatherController) respondWithError(c *gin.Context, statusCode int, error, message string) {
    // 不返回内部错误详情
    c.JSON(statusCode, APIResponse{
        Success: false,
        Error: &ErrorResponse{
            Error:   error,
            Code:    statusCode,
            Message: message, // 用户友好的消息
        },
    })
}
```

### 3. 配置安全

```go
// 敏感信息环境变量化
WEATHER_API_KEY=your_secret_key  // 不硬编码在代码中
```

## 性能优化

### 1. HTTP 客户端优化

```go
// 连接池和超时设置
client := &http.Client{
    Timeout: time.Duration(cfg.Timeout) * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

### 2. 响应缓存 (未来扩展)

```go
// 可以添加 Redis 缓存
type CachedWeatherService struct {
    weatherService WeatherService
    cache         Cache
    ttl           time.Duration
}
```

### 3. 并发处理

```go
// Gin 框架天然支持并发
// 每个请求在独立的 goroutine 中处理
```

## 可扩展性设计

### 1. 服务提供商扩展

```go
// 新增天气服务提供商
type AccuWeatherService struct {
    config *WeatherConfig
    client *http.Client
}

func (s *AccuWeatherService) GetWeatherByCity(city, units, lang string) (*WeatherResponse, error) {
    // AccuWeather API 实现
}
```

### 2. 功能扩展

```go
// 可以扩展的功能
- 天气预报 (Forecast)
- 历史天气 (Historical)
- 天气警报 (Alerts)
- 空气质量 (Air Quality)
```

### 3. 中间件扩展

```go
// 可以添加的中间件
- 限流中间件 (Rate Limiting)
- 认证中间件 (Authentication)
- 缓存中间件 (Caching)
- 监控中间件 (Metrics)
```

## 监控和日志

### 1. 日志设计

```go
// 结构化日志
{
    "timestamp": "2024-01-01T12:00:00Z",
    "level": "INFO",
    "request_id": "req-123456",
    "method": "GET",
    "path": "/api/v1/weather/city/Beijing",
    "status": 200,
    "duration": "150ms",
    "user_agent": "curl/7.68.0"
}
```

### 2. 健康检查

```go
// 健康检查端点
GET /api/v1/health
{
    "success": true,
    "data": {
        "status": "ok",
        "service": "gin-weather",
        "version": "1.0.0",
        "timestamp": 1640995200
    }
}
```

### 3. 指标监控 (未来扩展)

```go
// 可以集成 Prometheus 指标
- 请求总数
- 响应时间
- 错误率
- 外部 API 调用次数
```

## 部署架构

### 1. 单机部署

```
┌─────────────────┐
│   Load Balancer │
└─────────┬───────┘
          │
┌─────────▼───────┐
│  Gin Weather    │
│     Service     │
└─────────┬───────┘
          │
┌─────────▼───────┐
│ OpenWeatherMap  │
│      API        │
└─────────────────┘
```

### 2. 容器化部署

```
┌─────────────────┐
│   Docker Host   │
│  ┌─────────────┐│
│  │ gin-weather ││
│  │ Container   ││
│  └─────────────┘│
└─────────────────┘
```

### 3. 微服务部署 (未来扩展)

```
┌─────────────────┐
│   API Gateway   │
└─────────┬───────┘
          │
    ┌─────▼─────┐
    │  Weather  │
    │  Service  │
    └─────┬─────┘
          │
┌─────────▼───────┐
│ External APIs   │
└─────────────────┘
```

## 总结

这个架构设计具有以下特点：

1. **模块化**：清晰的模块划分，职责明确
2. **可扩展**：易于添加新功能和服务提供商
3. **可测试**：依赖注入，便于单元测试
4. **可维护**：代码结构清晰，遵循最佳实践
5. **高性能**：合理的并发处理和资源管理
6. **安全性**：输入验证，错误信息安全
7. **可观测**：完善的日志和监控设计

通过这种架构设计，项目具备了良好的可维护性、可扩展性和可靠性，能够满足生产环境的需求。
