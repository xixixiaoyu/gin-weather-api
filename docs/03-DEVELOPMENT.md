# Gin Weather API 开发指南

本文档为开发者提供详细的开发指南，包括开发环境搭建、代码规范、调试技巧、扩展指南等内容。

## 开发环境搭建

### 系统要求

- **Go 版本**: 1.21 或更高版本
- **操作系统**: Linux, macOS, Windows
- **内存**: 至少 2GB RAM
- **磁盘空间**: 至少 1GB 可用空间

### 必需工具

```bash
# 1. 安装 Go
# 从 https://golang.org/dl/ 下载并安装

# 2. 验证 Go 安装
go version

# 3. 设置 Go 环境变量
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 4. 安装开发工具
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 推荐工具

```bash
# 代码编辑器
- Visual Studio Code (推荐)
- GoLand
- Vim/Neovim

# VS Code 扩展
- Go (官方扩展)
- REST Client
- Docker
- GitLens

# 调试工具
- Delve (Go 调试器)
- Postman/Insomnia (API 测试)
- curl (命令行 HTTP 客户端)
```

## 项目设置

### 1. 克隆项目

```bash
git clone <repository-url>
cd gin-weather
```

### 2. 安装依赖

```bash
# 下载依赖
go mod download

# 整理依赖
go mod tidy

# 验证依赖
go mod verify
```

### 3. 环境配置

```bash
# 复制环境变量文件
cp .env.example .env

# 编辑环境变量
vim .env
```

必需的环境变量：
```env
WEATHER_API_KEY=your_openweathermap_api_key_here
```

可选的环境变量：
```env
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
GIN_MODE=debug
WEATHER_TIMEOUT=10
```

### 4. 获取 API 密钥

1. 访问 [OpenWeatherMap](https://openweathermap.org/api)
2. 注册免费账户
3. 获取 API 密钥
4. 将密钥添加到 `.env` 文件

## 开发工作流

### 1. 启动开发服务器

```bash
# 方法 1: 直接运行
go run cmd/server/main.go

# 方法 2: 使用 Makefile
make run

# 方法 3: 构建后运行
make build
./bin/gin-weather
```

### 2. 代码修改流程

```bash
# 1. 创建功能分支
git checkout -b feature/new-feature

# 2. 编写代码
# ... 修改代码 ...

# 3. 格式化代码
make fmt

# 4. 代码检查
make vet
make lint

# 5. 运行测试
make test

# 6. 提交代码
git add .
git commit -m "feat: add new feature"

# 7. 推送分支
git push origin feature/new-feature
```

### 3. 测试流程

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/config

# 运行测试并显示覆盖率
go test -cover ./...

# 生成覆盖率报告
make coverage
```

## 代码规范

### 1. 命名规范

```go
// 包名：小写，简短，有意义
package config

// 结构体：大驼峰命名
type WeatherService struct {}

// 接口：大驼峰命名，通常以 -er 结尾
type WeatherProvider interface {}

// 函数：大驼峰命名（公开），小驼峰命名（私有）
func GetWeather() {}
func validateRequest() {}

// 变量：小驼峰命名
var weatherService WeatherService

// 常量：大驼峰命名或全大写
const DefaultTimeout = 10
const API_VERSION = "v1"
```

### 2. 注释规范

```go
// Package config 提供应用程序配置管理功能
package config

// Config 应用配置结构体
// 包含服务器配置和天气服务配置
type Config struct {
    Server  ServerConfig  `json:"server"`  // 服务器配置
    Weather WeatherConfig `json:"weather"` // 天气服务配置
}

// Load 从环境变量加载配置
// 返回配置实例和可能的错误
func Load() (*Config, error) {
    // 实现逻辑...
}
```

### 3. 错误处理

```go
// 好的错误处理
func GetWeather(city string) (*WeatherResponse, error) {
    if city == "" {
        return nil, fmt.Errorf("城市名称不能为空")
    }
    
    resp, err := api.Call(city)
    if err != nil {
        return nil, fmt.Errorf("调用天气 API 失败: %w", err)
    }
    
    return resp, nil
}

// 避免的错误处理
func GetWeather(city string) *WeatherResponse {
    resp, _ := api.Call(city) // 忽略错误
    return resp
}
```

### 4. 结构体设计

```go
// 好的结构体设计
type WeatherService struct {
    config *Config      // 依赖注入
    client *http.Client // 可配置的客户端
    logger Logger       // 日志记录器
}

// 构造函数
func NewWeatherService(config *Config, logger Logger) *WeatherService {
    return &WeatherService{
        config: config,
        client: &http.Client{
            Timeout: time.Duration(config.Timeout) * time.Second,
        },
        logger: logger,
    }
}
```

## 调试技巧

### 1. 使用 Delve 调试器

```bash
# 安装 Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试
dlv debug cmd/server/main.go

# 调试命令
(dlv) break main.main    # 设置断点
(dlv) continue           # 继续执行
(dlv) next              # 单步执行
(dlv) print variable    # 打印变量
(dlv) stack             # 查看调用栈
```

### 2. 日志调试

```go
// 添加调试日志
func (s *OpenWeatherMapService) GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error) {
    log.Printf("开始查询城市天气: city=%s, units=%s, lang=%s", city, units, lang)
    
    // API 调用
    resp, err := s.client.Get(url)
    if err != nil {
        log.Printf("API 调用失败: %v", err)
        return nil, err
    }
    
    log.Printf("API 调用成功: status=%d", resp.StatusCode)
    return result, nil
}
```

### 3. 单元测试调试

```go
func TestWeatherService(t *testing.T) {
    // 使用 t.Logf 输出调试信息
    t.Logf("测试开始: %s", time.Now())
    
    service := NewWeatherService(config)
    result, err := service.GetWeatherByCity("Beijing", "metric", "zh_cn")
    
    if err != nil {
        t.Fatalf("测试失败: %v", err)
    }
    
    t.Logf("测试结果: %+v", result)
}
```

### 4. HTTP 请求调试

```bash
# 使用 curl 测试 API
curl -v "http://localhost:8080/api/v1/weather/city/Beijing"

# 使用 httpie (更友好的输出)
http GET localhost:8080/api/v1/weather/city/Beijing

# 测试错误情况
curl -v "http://localhost:8080/api/v1/weather/coordinates/91/181"
```

## 扩展指南

### 1. 添加新的天气服务提供商

#### 步骤 1: 创建服务实现

```go
// internal/service/accuweather.go
type AccuWeatherService struct {
    config *config.WeatherConfig
    client *http.Client
}

func NewAccuWeatherService(cfg *config.WeatherConfig) *AccuWeatherService {
    return &AccuWeatherService{
        config: cfg,
        client: &http.Client{
            Timeout: time.Duration(cfg.Timeout) * time.Second,
        },
    }
}

func (s *AccuWeatherService) GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error) {
    // 实现 AccuWeather API 调用逻辑
}

func (s *AccuWeatherService) GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error) {
    // 实现坐标查询逻辑
}
```

#### 步骤 2: 更新工厂方法

```go
// cmd/server/main.go
func createWeatherService(cfg *config.Config) service.WeatherService {
    switch cfg.Weather.Provider {
    case "openweathermap":
        return service.NewOpenWeatherMapService(&cfg.Weather)
    case "accuweather":
        return service.NewAccuWeatherService(&cfg.Weather)
    default:
        log.Fatalf("不支持的天气服务提供商: %s", cfg.Weather.Provider)
        return nil
    }
}
```

#### 步骤 3: 添加配置支持

```go
// internal/config/config.go
// 在 WeatherConfig 中添加新的配置项
type WeatherConfig struct {
    // ... 现有字段
    AccuWeatherAPIKey string `json:"accuweather_api_key"`
}
```

#### 步骤 4: 编写测试

```go
// internal/service/accuweather_test.go
func TestAccuWeatherService_GetWeatherByCity(t *testing.T) {
    // 测试实现
}
```

### 2. 添加新的 API 端点

#### 步骤 1: 扩展服务接口

```go
// internal/service/weather.go
type WeatherService interface {
    GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error)
    GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error)
    // 新增方法
    GetWeatherForecast(city string, days int, units, lang string) (*model.ForecastResponse, error)
}
```

#### 步骤 2: 添加数据模型

```go
// internal/model/forecast.go
type ForecastResponse struct {
    Location Location        `json:"location"`
    Forecast []DailyForecast `json:"forecast"`
    Provider string          `json:"provider"`
}

type DailyForecast struct {
    Date        string  `json:"date"`
    TempMax     float64 `json:"temp_max"`
    TempMin     float64 `json:"temp_min"`
    Description string  `json:"description"`
    // ... 其他字段
}
```

#### 步骤 3: 实现控制器方法

```go
// internal/controller/weather.go
func (wc *WeatherController) GetWeatherForecast(c *gin.Context) {
    city := c.Param("city")
    days := c.DefaultQuery("days", "5")
    
    // 参数验证和业务逻辑
    // ...
}
```

#### 步骤 4: 添加路由

```go
// internal/controller/router.go
weather.GET("/forecast/:city", weatherController.GetWeatherForecast)
```

### 3. 添加中间件

#### 限流中间件示例

```go
// internal/middleware/ratelimit.go
func RateLimitMiddleware(requests int, window time.Duration) gin.HandlerFunc {
    // 使用 sync.Map 存储客户端请求计数
    clients := sync.Map{}
    
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        
        // 检查请求频率
        if isRateLimited(clientIP, requests, window, &clients) {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error": "请求过于频繁，请稍后再试",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

#### 认证中间件示例

```go
// internal/middleware/auth.go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("X-API-Key")
        
        if !isValidAPIKey(apiKey) {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "无效的 API 密钥",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

## 性能优化

### 1. HTTP 客户端优化

```go
// 优化 HTTP 客户端配置
func createOptimizedHTTPClient(timeout int) *http.Client {
    transport := &http.Transport{
        MaxIdleConns:        100,              // 最大空闲连接数
        MaxIdleConnsPerHost: 10,               // 每个主机最大空闲连接数
        IdleConnTimeout:     90 * time.Second, // 空闲连接超时
        DisableCompression:  false,            // 启用压缩
        ForceAttemptHTTP2:   true,            // 强制使用 HTTP/2
    }
    
    return &http.Client{
        Timeout:   time.Duration(timeout) * time.Second,
        Transport: transport,
    }
}
```

### 2. 响应缓存

```go
// 简单的内存缓存实现
type Cache struct {
    data map[string]CacheItem
    mu   sync.RWMutex
}

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, exists := c.data[key]
    if !exists || time.Now().After(item.ExpiresAt) {
        return nil, false
    }
    
    return item.Value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    c.data[key] = CacheItem{
        Value:     value,
        ExpiresAt: time.Now().Add(ttl),
    }
}
```

### 3. 并发处理

```go
// 并发调用多个 API
func (s *WeatherService) GetWeatherFromMultipleSources(city string) (*WeatherResponse, error) {
    type result struct {
        response *WeatherResponse
        err      error
        source   string
    }
    
    results := make(chan result, 2)
    
    // 并发调用
    go func() {
        resp, err := s.openWeatherMap.GetWeatherByCity(city, "metric", "zh_cn")
        results <- result{resp, err, "openweathermap"}
    }()
    
    go func() {
        resp, err := s.accuWeather.GetWeatherByCity(city, "metric", "zh_cn")
        results <- result{resp, err, "accuweather"}
    }()
    
    // 等待第一个成功的响应
    for i := 0; i < 2; i++ {
        select {
        case res := <-results:
            if res.err == nil {
                return res.response, nil
            }
        case <-time.After(5 * time.Second):
            return nil, fmt.Errorf("请求超时")
        }
    }
    
    return nil, fmt.Errorf("所有数据源都失败")
}
```

## 部署指南

### 1. 本地部署

```bash
# 构建应用
make build

# 设置环境变量
export WEATHER_API_KEY=your_api_key

# 运行应用
./bin/gin-weather
```

### 2. Docker 部署

```bash
# 构建镜像
make docker-build

# 运行容器
docker run -p 8080:8080 -e WEATHER_API_KEY=your_api_key gin-weather
```

### 3. 生产环境部署

```bash
# 1. 构建生产版本
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-weather cmd/server/main.go

# 2. 创建系统服务文件
sudo vim /etc/systemd/system/gin-weather.service

# 3. 启动服务
sudo systemctl enable gin-weather
sudo systemctl start gin-weather
```

## 故障排查

### 1. 常见问题

#### 问题：服务启动失败
```bash
# 检查配置
go run cmd/server/main.go

# 检查环境变量
env | grep WEATHER

# 检查端口占用
lsof -i :8080
```

#### 问题：API 调用失败
```bash
# 检查网络连接
curl -v https://api.openweathermap.org/data/2.5/weather?q=Beijing&appid=your_key

# 检查 API 密钥
echo $WEATHER_API_KEY

# 查看日志
tail -f /var/log/gin-weather.log
```

#### 问题：测试失败
```bash
# 运行特定测试
go test -v ./internal/config

# 查看测试覆盖率
go test -cover ./...

# 运行基准测试
go test -bench=. ./...
```

### 2. 日志分析

```bash
# 查看错误日志
grep "ERROR" /var/log/gin-weather.log

# 查看慢请求
grep "duration.*[5-9][0-9][0-9]ms" /var/log/gin-weather.log

# 统计请求状态码
awk '{print $9}' /var/log/gin-weather.log | sort | uniq -c
```

## 贡献指南

### 1. 代码贡献流程

1. Fork 项目
2. 创建功能分支
3. 编写代码和测试
4. 提交 Pull Request
5. 代码审查
6. 合并代码

### 2. 提交信息规范

```bash
# 格式: type(scope): description
feat(weather): 添加新的天气服务提供商
fix(config): 修复配置加载错误
docs(api): 更新 API 文档
test(service): 添加服务层测试
refactor(controller): 重构控制器代码
```

### 3. 代码审查清单

- [ ] 代码符合项目规范
- [ ] 包含必要的测试
- [ ] 文档已更新
- [ ] 没有引入安全问题
- [ ] 性能影响可接受
- [ ] 向后兼容

通过遵循这个开发指南，您可以高效地开发和维护 Gin Weather API 项目。
