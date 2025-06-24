package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config 应用配置结构体
type Config struct {
	Server ServerConfig `json:"server"`
	Weather WeatherConfig `json:"weather"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
	Mode string `json:"mode"` // debug, release, test
}

// WeatherConfig 天气服务配置
type WeatherConfig struct {
	APIKey   string `json:"api_key"`
	BaseURL  string `json:"base_url"`
	Timeout  int    `json:"timeout"` // 请求超时时间（秒）
	Provider string `json:"provider"` // 天气服务提供商
}

// Load 从环境变量加载配置
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

	// 验证必需的配置项
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("配置验证失败: %w", err)
	}

	return config, nil
}

// validate 验证配置的有效性
func (c *Config) validate() error {
	if c.Weather.APIKey == "" {
		return fmt.Errorf("WEATHER_API_KEY 环境变量不能为空")
	}

	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("服务器端口必须在 1-65535 范围内")
	}

	if c.Weather.Timeout <= 0 {
		return fmt.Errorf("天气 API 超时时间必须大于 0")
	}

	return nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 获取环境变量并转换为整数，如果不存在或转换失败则返回默认值
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
