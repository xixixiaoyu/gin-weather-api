package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"gin-weather/internal/model"
	"gin-weather/internal/service"

	"github.com/gin-gonic/gin"
)

// WeatherController 天气控制器
type WeatherController struct {
	weatherService service.WeatherService
}

// NewWeatherController 创建天气控制器实例
func NewWeatherController(weatherService service.WeatherService) *WeatherController {
	return &WeatherController{
		weatherService: weatherService,
	}
}

// GetWeather 获取天气信息
// @Summary 获取天气信息
// @Description 根据城市名称或坐标获取当前天气信息
// @Tags weather
// @Accept json
// @Produce json
// @Param city query string false "城市名称（与坐标二选一）"
// @Param lat query number false "纬度（需要与经度一起使用）"
// @Param lon query number false "经度（需要与纬度一起使用）"
// @Param units query string false "单位系统" Enums(metric, imperial, standard) default(metric)
// @Param lang query string false "语言" default(zh_cn)
// @Success 200 {object} model.APIResponse{data=model.WeatherResponse}
// @Failure 400 {object} model.APIResponse{error=model.ErrorResponse}
// @Failure 500 {object} model.APIResponse{error=model.ErrorResponse}
// @Router /api/v1/weather [get]
func (wc *WeatherController) GetWeather(c *gin.Context) {
	// 解析查询参数
	var req model.WeatherRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		wc.respondWithError(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 验证参数
	if err := wc.validateRequest(&req); err != nil {
		wc.respondWithError(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 设置默认值
	if req.Units == "" {
		req.Units = "metric"
	}
	if req.Lang == "" {
		req.Lang = "zh_cn"
	}

	var weatherResp *model.WeatherResponse
	var err error

	// 根据请求类型调用相应的服务方法
	if req.City != "" {
		weatherResp, err = wc.weatherService.GetWeatherByCity(req.City, req.Units, req.Lang)
	} else {
		weatherResp, err = wc.weatherService.GetWeatherByCoordinates(req.Lat, req.Lon, req.Units, req.Lang)
	}

	if err != nil {
		wc.respondWithError(c, http.StatusInternalServerError, "获取天气信息失败", err.Error())
		return
	}

	// 返回成功响应
	wc.respondWithSuccess(c, weatherResp)
}

// GetWeatherByCity 根据城市名称获取天气信息
// @Summary 根据城市名称获取天气信息
// @Description 根据城市名称获取当前天气信息
// @Tags weather
// @Accept json
// @Produce json
// @Param city path string true "城市名称"
// @Param units query string false "单位系统" Enums(metric, imperial, standard) default(metric)
// @Param lang query string false "语言" default(zh_cn)
// @Success 200 {object} model.APIResponse{data=model.WeatherResponse}
// @Failure 400 {object} model.APIResponse{error=model.ErrorResponse}
// @Failure 500 {object} model.APIResponse{error=model.ErrorResponse}
// @Router /api/v1/weather/city/{city} [get]
func (wc *WeatherController) GetWeatherByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		wc.respondWithError(c, http.StatusBadRequest, "参数错误", "城市名称不能为空")
		return
	}

	units := c.DefaultQuery("units", "metric")
	lang := c.DefaultQuery("lang", "zh_cn")

	weatherResp, err := wc.weatherService.GetWeatherByCity(city, units, lang)
	if err != nil {
		wc.respondWithError(c, http.StatusInternalServerError, "获取天气信息失败", err.Error())
		return
	}

	wc.respondWithSuccess(c, weatherResp)
}

// GetWeatherByCoordinates 根据坐标获取天气信息
// @Summary 根据坐标获取天气信息
// @Description 根据经纬度坐标获取当前天气信息
// @Tags weather
// @Accept json
// @Produce json
// @Param lat path number true "纬度"
// @Param lon path number true "经度"
// @Param units query string false "单位系统" Enums(metric, imperial, standard) default(metric)
// @Param lang query string false "语言" default(zh_cn)
// @Success 200 {object} model.APIResponse{data=model.WeatherResponse}
// @Failure 400 {object} model.APIResponse{error=model.ErrorResponse}
// @Failure 500 {object} model.APIResponse{error=model.ErrorResponse}
// @Router /api/v1/weather/coordinates/{lat}/{lon} [get]
func (wc *WeatherController) GetWeatherByCoordinates(c *gin.Context) {
	latStr := c.Param("lat")
	lonStr := c.Param("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		wc.respondWithError(c, http.StatusBadRequest, "参数错误", "纬度格式不正确")
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		wc.respondWithError(c, http.StatusBadRequest, "参数错误", "经度格式不正确")
		return
	}

	// 验证坐标范围
	if lat < -90 || lat > 90 {
		wc.respondWithError(c, http.StatusBadRequest, "参数错误", "纬度必须在 -90 到 90 之间")
		return
	}
	if lon < -180 || lon > 180 {
		wc.respondWithError(c, http.StatusBadRequest, "参数错误", "经度必须在 -180 到 180 之间")
		return
	}

	units := c.DefaultQuery("units", "metric")
	lang := c.DefaultQuery("lang", "zh_cn")

	weatherResp, err := wc.weatherService.GetWeatherByCoordinates(lat, lon, units, lang)
	if err != nil {
		wc.respondWithError(c, http.StatusInternalServerError, "获取天气信息失败", err.Error())
		return
	}

	wc.respondWithSuccess(c, weatherResp)
}

// HealthCheck 健康检查接口
// @Summary 健康检查
// @Description 检查服务是否正常运行
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} model.APIResponse{data=map[string]string}
// @Router /api/v1/health [get]
func (wc *WeatherController) HealthCheck(c *gin.Context) {
	wc.respondWithSuccess(c, gin.H{
		"status":  "ok",
		"service": "gin-weather",
		"version": "1.0.0",
	})
}

// validateRequest 验证请求参数
func (wc *WeatherController) validateRequest(req *model.WeatherRequest) error {
	// 城市名称和坐标必须提供其中一个
	if req.City == "" && (req.Lat == 0 && req.Lon == 0) {
		return fmt.Errorf("必须提供城市名称或坐标信息")
	}

	// 如果提供了坐标，需要验证范围
	if req.Lat != 0 || req.Lon != 0 {
		if req.Lat < -90 || req.Lat > 90 {
			return fmt.Errorf("纬度必须在 -90 到 90 之间")
		}
		if req.Lon < -180 || req.Lon > 180 {
			return fmt.Errorf("经度必须在 -180 到 180 之间")
		}
	}

	return nil
}

// respondWithError 返回错误响应
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

// respondWithSuccess 返回成功响应
func (wc *WeatherController) respondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Data:    data,
	})
}
