package model

import "time"

// WeatherRequest 天气查询请求结构体
type WeatherRequest struct {
	City      string  `json:"city" form:"city" binding:"required_without=Lat"`           // 城市名称
	Lat       float64 `json:"lat" form:"lat" binding:"required_without=City"`            // 纬度
	Lon       float64 `json:"lon" form:"lon" binding:"required_with=Lat"`                // 经度
	Units     string  `json:"units" form:"units" binding:"omitempty,oneof=metric imperial standard"` // 单位系统
	Lang      string  `json:"lang" form:"lang" binding:"omitempty"`                      // 语言
}

// WeatherResponse 标准化的天气响应结构体
type WeatherResponse struct {
	Location    Location    `json:"location"`    // 位置信息
	Current     Current     `json:"current"`     // 当前天气
	Timestamp   int64       `json:"timestamp"`   // 响应时间戳
	Provider    string      `json:"provider"`    // 数据提供商
}

// Location 位置信息
type Location struct {
	Name      string  `json:"name"`      // 城市名称
	Country   string  `json:"country"`   // 国家代码
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
	Timezone  int     `json:"timezone"`  // 时区偏移（秒）
}

// Current 当前天气信息
type Current struct {
	Temperature     float64     `json:"temperature"`      // 当前温度
	FeelsLike       float64     `json:"feels_like"`       // 体感温度
	TempMin         float64     `json:"temp_min"`         // 最低温度
	TempMax         float64     `json:"temp_max"`         // 最高温度
	Pressure        int         `json:"pressure"`         // 大气压力（hPa）
	Humidity        int         `json:"humidity"`         // 湿度（%）
	Visibility      int         `json:"visibility"`       // 能见度（米）
	UVIndex         float64     `json:"uv_index"`         // 紫外线指数
	Weather         []Weather   `json:"weather"`          // 天气状况
	Wind            Wind        `json:"wind"`             // 风力信息
	Clouds          Clouds      `json:"clouds"`           // 云量信息
	Rain            *Rain       `json:"rain,omitempty"`   // 降雨信息
	Snow            *Snow       `json:"snow,omitempty"`   // 降雪信息
	Sunrise         int64       `json:"sunrise"`          // 日出时间戳
	Sunset          int64       `json:"sunset"`           // 日落时间戳
	UpdatedAt       time.Time   `json:"updated_at"`       // 数据更新时间
}

// Weather 天气状况
type Weather struct {
	ID          int    `json:"id"`          // 天气状况 ID
	Main        string `json:"main"`        // 天气主要状况
	Description string `json:"description"` // 天气详细描述
	Icon        string `json:"icon"`        // 天气图标代码
}

// Wind 风力信息
type Wind struct {
	Speed     float64 `json:"speed"`     // 风速
	Direction int     `json:"direction"` // 风向（度）
	Gust      float64 `json:"gust"`      // 阵风速度
}

// Clouds 云量信息
type Clouds struct {
	All int `json:"all"` // 云量百分比
}

// Rain 降雨信息
type Rain struct {
	OneHour   float64 `json:"1h,omitempty"`   // 过去1小时降雨量（mm）
	ThreeHour float64 `json:"3h,omitempty"`   // 过去3小时降雨量（mm）
}

// Snow 降雪信息
type Snow struct {
	OneHour   float64 `json:"1h,omitempty"`   // 过去1小时降雪量（mm）
	ThreeHour float64 `json:"3h,omitempty"`   // 过去3小时降雪量（mm）
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Error   string `json:"error"`             // 错误信息
	Code    int    `json:"code"`              // 错误代码
	Message string `json:"message,omitempty"` // 详细错误信息
}

// APIResponse 通用 API 响应结构体
type APIResponse struct {
	Success bool        `json:"success"`           // 请求是否成功
	Data    interface{} `json:"data,omitempty"`    // 响应数据
	Error   *ErrorResponse `json:"error,omitempty"` // 错误信息
}
