# Gin Weather API 接口文档

## 概述

Gin Weather API 是一个基于 Go 语言和 Gin 框架的天气查询代理服务，提供标准化的天气数据查询接口。

- **版本**: v1.0.0
- **基础 URL**: `http://localhost:8080/api/v1`
- **协议**: HTTP/HTTPS
- **数据格式**: JSON
- **字符编码**: UTF-8

## 认证

当前版本不需要认证，但建议在生产环境中添加 API 密钥认证。

## 通用响应格式

所有 API 接口都遵循统一的响应格式：

### 成功响应

```json
{
  "success": true,
  "data": {
    // 具体的响应数据
  }
}
```

### 错误响应

```json
{
  "success": false,
  "error": {
    "error": "错误类型",
    "code": 400,
    "message": "详细错误信息"
  }
}
```

## 状态码

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 500 | 服务器内部错误 |

## 接口列表

### 1. 健康检查

检查服务是否正常运行。

**请求**

```http
GET /api/v1/health
```

**响应**

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

### 2. 通用天气查询

支持通过城市名称或地理坐标查询天气信息。

**请求**

```http
GET /api/v1/weather
```

**查询参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| city | string | 否* | 城市名称（与坐标二选一） |
| lat | float | 否* | 纬度（需要与经度一起使用） |
| lon | float | 否* | 经度（需要与纬度一起使用） |
| units | string | 否 | 单位系统：metric（默认）、imperial、standard |
| lang | string | 否 | 语言代码，默认 zh_cn |

*注：city 和 (lat, lon) 必须提供其中一组

**示例请求**

```bash
# 通过城市名称查询
curl "http://localhost:8080/api/v1/weather?city=Beijing&units=metric&lang=zh_cn"

# 通过坐标查询
curl "http://localhost:8080/api/v1/weather?lat=39.9042&lon=116.4074&units=metric&lang=zh_cn"
```

**响应**

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
      "rain": {
        "1h": 0,
        "3h": 0
      },
      "snow": {
        "1h": 0,
        "3h": 0
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

### 3. 根据城市查询天气

通过城市名称查询天气信息的便捷接口。

**请求**

```http
GET /api/v1/weather/city/{city}
```

**路径参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| city | string | 是 | 城市名称 |

**查询参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| units | string | 否 | 单位系统：metric（默认）、imperial、standard |
| lang | string | 否 | 语言代码，默认 zh_cn |

**示例请求**

```bash
curl "http://localhost:8080/api/v1/weather/city/Shanghai?units=metric&lang=zh_cn"
```

**响应**

响应格式与通用天气查询接口相同。

### 4. 根据坐标查询天气

通过地理坐标查询天气信息的便捷接口。

**请求**

```http
GET /api/v1/weather/coordinates/{lat}/{lon}
```

**路径参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| lat | float | 是 | 纬度（-90 到 90） |
| lon | float | 是 | 经度（-180 到 180） |

**查询参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| units | string | 否 | 单位系统：metric（默认）、imperial、standard |
| lang | string | 否 | 语言代码，默认 zh_cn |

**示例请求**

```bash
curl "http://localhost:8080/api/v1/weather/coordinates/31.2304/121.4737?units=metric&lang=zh_cn"
```

**响应**

响应格式与通用天气查询接口相同。

## 数据字段说明

### Location（位置信息）

| 字段 | 类型 | 说明 |
|------|------|------|
| name | string | 城市名称 |
| country | string | 国家代码 |
| latitude | float | 纬度 |
| longitude | float | 经度 |
| timezone | int | 时区偏移（秒） |

### Current（当前天气）

| 字段 | 类型 | 说明 |
|------|------|------|
| temperature | float | 当前温度 |
| feels_like | float | 体感温度 |
| temp_min | float | 最低温度 |
| temp_max | float | 最高温度 |
| pressure | int | 大气压力（hPa） |
| humidity | int | 湿度（%） |
| visibility | int | 能见度（米） |
| uv_index | float | 紫外线指数 |
| weather | array | 天气状况数组 |
| wind | object | 风力信息 |
| clouds | object | 云量信息 |
| rain | object | 降雨信息（可选） |
| snow | object | 降雪信息（可选） |
| sunrise | int | 日出时间戳 |
| sunset | int | 日落时间戳 |
| updated_at | string | 数据更新时间 |

### Weather（天气状况）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int | 天气状况 ID |
| main | string | 天气主要状况 |
| description | string | 天气详细描述 |
| icon | string | 天气图标代码 |

### Wind（风力信息）

| 字段 | 类型 | 说明 |
|------|------|------|
| speed | float | 风速 |
| direction | int | 风向（度） |
| gust | float | 阵风速度 |

### Clouds（云量信息）

| 字段 | 类型 | 说明 |
|------|------|------|
| all | int | 云量百分比 |

### Rain/Snow（降水信息）

| 字段 | 类型 | 说明 |
|------|------|------|
| 1h | float | 过去1小时降水量（mm） |
| 3h | float | 过去3小时降水量（mm） |

## 单位系统

### metric（公制，默认）
- 温度：摄氏度（°C）
- 风速：米/秒（m/s）
- 压力：百帕（hPa）

### imperial（英制）
- 温度：华氏度（°F）
- 风速：英里/小时（mph）
- 压力：百帕（hPa）

### standard（标准）
- 温度：开尔文（K）
- 风速：米/秒（m/s）
- 压力：百帕（hPa）

## 语言支持

支持的语言代码包括但不限于：

- `zh_cn` - 简体中文（默认）
- `en` - 英语
- `zh_tw` - 繁体中文
- `ja` - 日语
- `ko` - 韩语
- `fr` - 法语
- `de` - 德语
- `es` - 西班牙语

## 错误处理

### 常见错误

| 错误代码 | 错误信息 | 说明 |
|----------|----------|------|
| 400 | 参数验证失败 | 请求参数格式错误或缺失 |
| 400 | 城市名称不能为空 | 城市参数为空 |
| 400 | 纬度格式不正确 | 纬度参数格式错误 |
| 400 | 经度格式不正确 | 经度参数格式错误 |
| 400 | 纬度必须在 -90 到 90 之间 | 纬度超出有效范围 |
| 400 | 经度必须在 -180 到 180 之间 | 经度超出有效范围 |
| 500 | 获取天气信息失败 | 第三方 API 调用失败 |

### 错误响应示例

```json
{
  "success": false,
  "error": {
    "error": "参数验证失败",
    "code": 400,
    "message": "纬度必须在 -90 到 90 之间"
  }
}
```

## 限制说明

- 基于 OpenWeatherMap 免费账户限制：
  - 每分钟最多 60 次请求
  - 每月最多 1,000,000 次请求
- 建议在生产环境中实现请求限流和缓存机制

## 更新日志

### v1.0.0
- 初始版本发布
- 支持城市名称和坐标查询
- 集成 OpenWeatherMap API
- 提供标准化响应格式
