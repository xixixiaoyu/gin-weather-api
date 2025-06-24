package service

import (
	"gin-weather/internal/model"
)

// WeatherService 天气服务接口
type WeatherService interface {
	// GetWeatherByCity 根据城市名称获取天气信息
	GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error)
	
	// GetWeatherByCoordinates 根据坐标获取天气信息
	GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error)
}
