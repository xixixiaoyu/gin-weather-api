# Gin Weather API

基于 Go 语言和 Gin 框架的天气查询 API 代理服务，提供标准化的天气数据查询接口。

## 功能特性

### 🖥️ 前端界面
- 🎨 **现代化 UI**：基于 Vue.js 3 + TypeScript 的响应式界面
- 🔍 **智能搜索**：支持城市名称搜索和地理位置定位
- 📱 **移动适配**：完美支持手机、平板和桌面设备
- ⚙️ **个性化设置**：温度单位、语言、自动定位等配置
- 🎯 **用户体验**：加载动画、错误提示、搜索历史

### 🔧 后端 API
- 🌤️ **多种查询方式**：支持城市名称和地理坐标查询
- 🔄 **API 代理**：作为第三方天气 API 的代理服务
- 📊 **标准化响应**：统一的 JSON 响应格式
- 🛡️ **参数验证**：完善的请求参数验证
- 🚀 **高性能**：基于 Gin 框架，性能优异
- 📝 **完整日志**：详细的请求和错误日志
- 🔧 **配置灵活**：支持环境变量配置
- ✅ **单元测试**：完整的测试覆盖

## 技术栈

### 前端技术
- **框架**：Vue.js 3 (Composition API)
- **语言**：TypeScript
- **构建工具**：Vite
- **状态管理**：Pinia
- **样式**：UnoCSS
- **图标**：Lucide Vue Next
- **HTTP 客户端**：Axios

### 后端技术
- **语言**：Go 1.21+
- **框架**：Gin Web Framework
- **第三方 API**：OpenWeatherMap
- **配置管理**：环境变量
- **测试**：Go 内置测试框架

### 部署技术
- **容器化**：Docker + Docker Compose
- **Web 服务器**：Nginx (前端)
- **反向代理**：Nginx → Gin

## 快速开始

### 1. 克隆项目

```bash
git clone <repository-url>
cd gin-weather
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置环境变量

复制环境变量示例文件：

```bash
cp .env.example .env
```

编辑 `.env` 文件，填入你的 OpenWeatherMap API 密钥：

```env
WEATHER_API_KEY=your_openweathermap_api_key_here
```

> 💡 **获取 API 密钥**：访问 [OpenWeatherMap](https://openweathermap.org/api) 注册账户并获取免费 API 密钥

### 4. 启动后端服务

```bash
go run cmd/server/main.go
```

后端服务将在 `http://localhost:8080` 启动。

### 5. 启动前端服务

```bash
cd frontend
npm install
npm run dev
```

前端应用将在 `http://localhost:5173` 启动。

### 6. 访问应用

- **前端界面**: http://localhost:5173
- **后端 API**: http://localhost:8080

### 7. 测试 API

访问健康检查接口：

```bash
curl http://localhost:8080/api/v1/health
```

查询北京天气：

```bash
curl "http://localhost:8080/api/v1/weather/city/Beijing"
```

## API 文档

### 基础信息

- **Base URL**: `http://localhost:8080/api/v1`
- **响应格式**: JSON
- **字符编码**: UTF-8

### 通用响应格式

```json
{
  "success": true,
  "data": {
    // 响应数据
  },
  "error": {
    "error": "错误类型",
    "code": 400,
    "message": "详细错误信息"
  }
}
```

### 接口列表

#### 1. 健康检查

```http
GET /api/v1/health
```

**响应示例**：
```json
{
  "success": true,
  "data": {
    "status": "ok",
    "service": "gin-weather",
    "version": "1.0.0"
  }
}
```

#### 2. 通用天气查询

```http
GET /api/v1/weather?city=Beijing&units=metric&lang=zh_cn
GET /api/v1/weather?lat=39.9042&lon=116.4074&units=metric&lang=zh_cn
```

**查询参数**：
- `city` (string): 城市名称（与坐标二选一）
- `lat` (float): 纬度（需要与经度一起使用）
- `lon` (float): 经度（需要与纬度一起使用）
- `units` (string): 单位系统，可选值：`metric`（默认）、`imperial`、`standard`
- `lang` (string): 语言，默认 `zh_cn`

#### 3. 根据城市查询

```http
GET /api/v1/weather/city/{city}?units=metric&lang=zh_cn
```

**路径参数**：
- `city` (string): 城市名称

#### 4. 根据坐标查询

```http
GET /api/v1/weather/coordinates/{lat}/{lon}?units=metric&lang=zh_cn
```

**路径参数**：
- `lat` (float): 纬度 (-90 到 90)
- `lon` (float): 经度 (-180 到 180)

### 响应数据结构

```json
{
  "success": true,
  "data": {
    "location": {
      "name": "Beijing",
      "country": "CN",
      "latitude": 39.9042,
      "longitude": 116.4074,
      "timezone": 28800
    },
    "current": {
      "temperature": 25.5,
      "feels_like": 27.0,
      "temp_min": 20.0,
      "temp_max": 30.0,
      "pressure": 1013,
      "humidity": 60,
      "visibility": 10000,
      "uv_index": 0,
      "weather": [
        {
          "id": 800,
          "main": "Clear",
          "description": "晴天",
          "icon": "01d"
        }
      ],
      "wind": {
        "speed": 3.5,
        "direction": 180,
        "gust": 5.0
      },
      "clouds": {
        "all": 0
      },
      "sunrise": 1640995200,
      "sunset": 1641031200,
      "updated_at": "2024-01-01T12:00:00Z"
    },
    "timestamp": 1640995200,
    "provider": "openweathermap"
  }
}
```

## 配置说明

### 环境变量

| 变量名 | 描述 | 默认值 | 必需 |
|--------|------|--------|------|
| `SERVER_HOST` | 服务器监听地址 | `0.0.0.0` | 否 |
| `SERVER_PORT` | 服务器端口 | `8080` | 否 |
| `GIN_MODE` | Gin 运行模式 | `debug` | 否 |
| `WEATHER_API_KEY` | 天气 API 密钥 | - | 是 |
| `WEATHER_BASE_URL` | 天气 API 基础 URL | `https://api.openweathermap.org/data/2.5` | 否 |
| `WEATHER_TIMEOUT` | API 请求超时时间（秒） | `10` | 否 |
| `WEATHER_PROVIDER` | 天气服务提供商 | `openweathermap` | 否 |

## 开发指南

### 项目结构

```
gin-weather/
├── cmd/server/          # 主程序入口
├── internal/
│   ├── config/         # 配置管理
│   ├── controller/     # 控制器层
│   ├── model/          # 数据模型
│   └── service/        # 服务层
├── pkg/                # 公共包
├── configs/            # 配置文件
├── docs/               # 文档
├── .env.example        # 环境变量示例
├── go.mod              # Go 模块文件
└── README.md           # 项目说明
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/config

# 运行测试并显示覆盖率
go test -cover ./...
```

### 构建项目

```bash
# 构建可执行文件
go build -o bin/gin-weather cmd/server/main.go

# 交叉编译（Linux）
GOOS=linux GOARCH=amd64 go build -o bin/gin-weather-linux cmd/server/main.go
```

## 部署

### Docker 部署

创建 `Dockerfile`：

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o gin-weather cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/gin-weather .
EXPOSE 8080
CMD ["./gin-weather"]
```

构建和运行：

```bash
docker build -t gin-weather .
docker run -p 8080:8080 -e WEATHER_API_KEY=your_api_key gin-weather
```

### 使用 Docker Compose（推荐）

一键启动前后端服务：

```bash
docker-compose up -d
```

这将同时启动：
- **前端服务**: http://localhost:3000
- **后端服务**: http://localhost:8080

停止服务：

```bash
docker-compose down
```

## 📚 完整学习文档

我们为您准备了详细的学习文档，建议按以下顺序阅读：

1. **[docs/00-README.md](docs/00-README.md)** - 📖 文档导航中心
2. **[docs/01-TUTORIAL.md](docs/01-TUTORIAL.md)** - 🎓 项目实现教程（从零开始学习）
3. **[docs/02-ARCHITECTURE.md](docs/02-ARCHITECTURE.md)** - 🏗️ 架构设计文档
4. **[docs/03-DEVELOPMENT.md](docs/03-DEVELOPMENT.md)** - 💻 开发指南
5. **[docs/04-API.md](docs/04-API.md)** - 📡 API 接口文档
6. **[docs/05-DEPLOYMENT.md](docs/05-DEPLOYMENT.md)** - 🚀 部署指南
7. **[docs/06-USER-GUIDE.md](docs/06-USER-GUIDE.md)** - 👤 用户使用指南
8. **[docs/07-FRONTEND-TUTORIAL.md](docs/07-FRONTEND-TUTORIAL.md)** - 🎨 前端实现教程

### 🎯 快速导航

- **初学者** → 从 [01-TUTORIAL.md](docs/01-TUTORIAL.md) 开始
- **有经验的开发者** → 直接查看 [02-ARCHITECTURE.md](docs/02-ARCHITECTURE.md)
- **前端开发者** → 查看 [07-FRONTEND-TUTORIAL.md](docs/07-FRONTEND-TUTORIAL.md)
- **API 使用者** → 查看 [04-API.md](docs/04-API.md)
- **运维人员** → 查看 [05-DEPLOYMENT.md](docs/05-DEPLOYMENT.md)

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 联系方式

如有问题或建议，请提交 Issue 或联系维护者。
