package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"gin-weather/internal/config"
	"gin-weather/internal/model"
)

// OpenWeatherMapService OpenWeatherMap 天气服务实现
type OpenWeatherMapService struct {
	config *config.WeatherConfig
	client *http.Client
}

// NewOpenWeatherMapService 创建 OpenWeatherMap 服务实例
func NewOpenWeatherMapService(cfg *config.WeatherConfig) *OpenWeatherMapService {
	return &OpenWeatherMapService{
		config: cfg,
		client: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
	}
}

// GetWeatherByCity 根据城市名称获取天气信息
func (s *OpenWeatherMapService) GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error) {
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", s.config.APIKey)
	params.Add("units", s.getUnits(units))
	params.Add("lang", s.getLang(lang))

	return s.fetchWeather(params)
}

// GetWeatherByCoordinates 根据坐标获取天气信息
func (s *OpenWeatherMapService) GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error) {
	params := url.Values{}
	params.Add("lat", strconv.FormatFloat(lat, 'f', 6, 64))
	params.Add("lon", strconv.FormatFloat(lon, 'f', 6, 64))
	params.Add("appid", s.config.APIKey)
	params.Add("units", s.getUnits(units))
	params.Add("lang", s.getLang(lang))

	return s.fetchWeather(params)
}

// fetchWeather 发起天气 API 请求
func (s *OpenWeatherMapService) fetchWeather(params url.Values) (*model.WeatherResponse, error) {
	// 构建请求 URL
	requestURL := fmt.Sprintf("%s/weather?%s", s.config.BaseURL, params.Encode())

	// 发起 HTTP 请求
	resp, err := s.client.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("请求天气 API 失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		var errorResp OpenWeatherMapError
		if err := json.Unmarshal(body, &errorResp); err == nil {
			return nil, fmt.Errorf("天气 API 错误 [%d]: %s", errorResp.Cod, errorResp.Message)
		}
		return nil, fmt.Errorf("天气 API 请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析响应数据
	var owmResp OpenWeatherMapResponse
	if err := json.Unmarshal(body, &owmResp); err != nil {
		return nil, fmt.Errorf("解析天气数据失败: %w", err)
	}

	// 转换为标准格式
	return s.convertToStandardFormat(&owmResp), nil
}

// getUnits 获取单位系统，默认为 metric
func (s *OpenWeatherMapService) getUnits(units string) string {
	if units == "" {
		return "metric"
	}
	return units
}

// getLang 获取语言设置，默认为 zh_cn
func (s *OpenWeatherMapService) getLang(lang string) string {
	if lang == "" {
		return "zh_cn"
	}
	return lang
}

// convertToStandardFormat 将 OpenWeatherMap 响应转换为标准格式
func (s *OpenWeatherMapService) convertToStandardFormat(owm *OpenWeatherMapResponse) *model.WeatherResponse {
	weather := make([]model.Weather, len(owm.Weather))
	for i, w := range owm.Weather {
		weather[i] = model.Weather{
			ID:          w.ID,
			Main:        w.Main,
			Description: w.Description,
			Icon:        w.Icon,
		}
	}

	var rain *model.Rain
	if owm.Rain != nil {
		rain = &model.Rain{
			OneHour:   owm.Rain.OneHour,
			ThreeHour: owm.Rain.ThreeHour,
		}
	}

	var snow *model.Snow
	if owm.Snow != nil {
		snow = &model.Snow{
			OneHour:   owm.Snow.OneHour,
			ThreeHour: owm.Snow.ThreeHour,
		}
	}

	return &model.WeatherResponse{
		Location: model.Location{
			Name:      owm.Name,
			Country:   owm.Sys.Country,
			Latitude:  owm.Coord.Lat,
			Longitude: owm.Coord.Lon,
			Timezone:  owm.Timezone,
		},
		Current: model.Current{
			Temperature: owm.Main.Temp,
			FeelsLike:   owm.Main.FeelsLike,
			TempMin:     owm.Main.TempMin,
			TempMax:     owm.Main.TempMax,
			Pressure:    owm.Main.Pressure,
			Humidity:    owm.Main.Humidity,
			Visibility:  owm.Visibility,
			Weather:     weather,
			Wind: model.Wind{
				Speed:     owm.Wind.Speed,
				Direction: owm.Wind.Deg,
				Gust:      owm.Wind.Gust,
			},
			Clouds: model.Clouds{
				All: owm.Clouds.All,
			},
			Rain:      rain,
			Snow:      snow,
			Sunrise:   int64(owm.Sys.Sunrise),
			Sunset:    int64(owm.Sys.Sunset),
			UpdatedAt: time.Unix(int64(owm.Dt), 0),
		},
		Timestamp: time.Now().Unix(),
		Provider:  "openweathermap",
	}
}

// OpenWeatherMap API 响应结构体定义

// OpenWeatherMapResponse OpenWeatherMap API 响应结构体
type OpenWeatherMapResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []OWMWeather `json:"weather"`
	Base       string    `json:"base"`
	Main       OWMMain   `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       OWMWind   `json:"wind"`
	Clouds     OWMClouds `json:"clouds"`
	Rain       *OWMRain  `json:"rain,omitempty"`
	Snow       *OWMSnow  `json:"snow,omitempty"`
	Dt         int       `json:"dt"`
	Sys        OWMSys    `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

// Coord 坐标信息
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// OWMWeather 天气状况
type OWMWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// OWMMain 主要天气数据
type OWMMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// OWMWind 风力信息
type OWMWind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust,omitempty"`
}

// OWMClouds 云量信息
type OWMClouds struct {
	All int `json:"all"`
}

// OWMRain 降雨信息
type OWMRain struct {
	OneHour   float64 `json:"1h,omitempty"`
	ThreeHour float64 `json:"3h,omitempty"`
}

// OWMSnow 降雪信息
type OWMSnow struct {
	OneHour   float64 `json:"1h,omitempty"`
	ThreeHour float64 `json:"3h,omitempty"`
}

// OWMSys 系统信息
type OWMSys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

// OpenWeatherMapError 错误响应结构体
type OpenWeatherMapError struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}
