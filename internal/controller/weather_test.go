package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gin-weather/internal/model"

	"github.com/gin-gonic/gin"
)

// MockWeatherService 模拟天气服务
type MockWeatherService struct{}

func (m *MockWeatherService) GetWeatherByCity(city, units, lang string) (*model.WeatherResponse, error) {
	return &model.WeatherResponse{
		Location: model.Location{
			Name:      city,
			Country:   "CN",
			Latitude:  39.9042,
			Longitude: 116.4074,
			Timezone:  28800,
		},
		Current: model.Current{
			Temperature: 25.5,
			FeelsLike:   27.0,
			TempMin:     20.0,
			TempMax:     30.0,
			Pressure:    1013,
			Humidity:    60,
			Visibility:  10000,
			Weather: []model.Weather{
				{
					ID:          800,
					Main:        "Clear",
					Description: "晴天",
					Icon:        "01d",
				},
			},
			Wind: model.Wind{
				Speed:     3.5,
				Direction: 180,
				Gust:      5.0,
			},
			Clouds: model.Clouds{
				All: 0,
			},
			Sunrise:   1640995200,
			Sunset:    1641031200,
			UpdatedAt: time.Now(),
		},
		Timestamp: time.Now().Unix(),
		Provider:  "openweathermap",
	}, nil
}

func (m *MockWeatherService) GetWeatherByCoordinates(lat, lon float64, units, lang string) (*model.WeatherResponse, error) {
	return m.GetWeatherByCity("Test City", units, lang)
}

func TestWeatherController_GetWeatherByCity(t *testing.T) {
	// 设置 Gin 为测试模式
	gin.SetMode(gin.TestMode)

	// 创建模拟服务和控制器
	mockService := &MockWeatherService{}
	controller := NewWeatherController(mockService)

	// 创建路由
	router := gin.New()
	router.GET("/weather/city/:city", controller.GetWeatherByCity)

	// 创建测试请求
	req, _ := http.NewRequest("GET", "/weather/city/Beijing", nil)
	w := httptest.NewRecorder()

	// 执行请求
	router.ServeHTTP(w, req)

	// 验证响应
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，实际为 %d", w.Code)
	}

	var response model.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	if !response.Success {
		t.Error("期望请求成功")
	}

	if response.Data == nil {
		t.Error("期望返回天气数据")
	}
}

func TestWeatherController_GetWeatherByCoordinates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockWeatherService{}
	controller := NewWeatherController(mockService)

	router := gin.New()
	router.GET("/weather/coordinates/:lat/:lon", controller.GetWeatherByCoordinates)

	req, _ := http.NewRequest("GET", "/weather/coordinates/39.9042/116.4074", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，实际为 %d", w.Code)
	}

	var response model.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	if !response.Success {
		t.Error("期望请求成功")
	}
}

func TestWeatherController_GetWeatherByCoordinatesInvalidLat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockWeatherService{}
	controller := NewWeatherController(mockService)

	router := gin.New()
	router.GET("/weather/coordinates/:lat/:lon", controller.GetWeatherByCoordinates)

	// 测试无效纬度
	req, _ := http.NewRequest("GET", "/weather/coordinates/invalid/116.4074", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400，实际为 %d", w.Code)
	}
}

func TestWeatherController_GetWeatherByCoordinatesOutOfRange(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockWeatherService{}
	controller := NewWeatherController(mockService)

	router := gin.New()
	router.GET("/weather/coordinates/:lat/:lon", controller.GetWeatherByCoordinates)

	// 测试超出范围的纬度
	req, _ := http.NewRequest("GET", "/weather/coordinates/91/116.4074", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("期望状态码 400，实际为 %d", w.Code)
	}
}

func TestWeatherController_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockWeatherService{}
	controller := NewWeatherController(mockService)

	router := gin.New()
	router.GET("/health", controller.HealthCheck)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，实际为 %d", w.Code)
	}

	var response model.APIResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	if !response.Success {
		t.Error("期望健康检查成功")
	}
}
