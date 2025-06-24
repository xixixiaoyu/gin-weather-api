package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// 设置测试环境变量
	os.Setenv("WEATHER_API_KEY", "test_api_key")
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("WEATHER_TIMEOUT", "15")
	defer func() {
		os.Unsetenv("WEATHER_API_KEY")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("WEATHER_TIMEOUT")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证配置值
	if cfg.Weather.APIKey != "test_api_key" {
		t.Errorf("期望 API Key 为 'test_api_key'，实际为 '%s'", cfg.Weather.APIKey)
	}

	if cfg.Server.Port != 9090 {
		t.Errorf("期望端口为 9090，实际为 %d", cfg.Server.Port)
	}

	if cfg.Weather.Timeout != 15 {
		t.Errorf("期望超时时间为 15，实际为 %d", cfg.Weather.Timeout)
	}
}

func TestLoadWithoutAPIKey(t *testing.T) {
	// 确保没有设置 API Key
	os.Unsetenv("WEATHER_API_KEY")

	_, err := Load()
	if err == nil {
		t.Error("期望在没有 API Key 时返回错误")
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV", "test_value")
	defer os.Unsetenv("TEST_ENV")

	value := getEnv("TEST_ENV", "default")
	if value != "test_value" {
		t.Errorf("期望值为 'test_value'，实际为 '%s'", value)
	}

	value = getEnv("NON_EXISTENT_ENV", "default")
	if value != "default" {
		t.Errorf("期望默认值为 'default'，实际为 '%s'", value)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	defer os.Unsetenv("TEST_INT")

	value := getEnvAsInt("TEST_INT", 456)
	if value != 123 {
		t.Errorf("期望值为 123，实际为 %d", value)
	}

	value = getEnvAsInt("NON_EXISTENT_INT", 456)
	if value != 456 {
		t.Errorf("期望默认值为 456，实际为 %d", value)
	}

	// 测试无效的整数值
	os.Setenv("INVALID_INT", "not_a_number")
	defer os.Unsetenv("INVALID_INT")

	value = getEnvAsInt("INVALID_INT", 789)
	if value != 789 {
		t.Errorf("期望默认值为 789，实际为 %d", value)
	}
}
