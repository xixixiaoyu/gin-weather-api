#!/bin/bash

# Gin Weather API 测试脚本
# 使用方法: ./test_api.sh [API_KEY]

set -e

# 配置
BASE_URL="http://localhost:8080/api/v1"
API_KEY=${1:-"your_api_key_here"}

echo "🌤️  Gin Weather API 测试脚本"
echo "================================"
echo "Base URL: $BASE_URL"
echo ""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 测试函数
test_endpoint() {
    local name="$1"
    local url="$2"
    local expected_status="$3"
    
    echo -e "${BLUE}测试: $name${NC}"
    echo "URL: $url"
    
    response=$(curl -s -w "\n%{http_code}" "$url")
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    if [ "$http_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ 状态码: $http_code (期望: $expected_status)${NC}"
        if [ "$expected_status" -eq 200 ]; then
            echo "响应数据:"
            echo "$body" | jq '.' 2>/dev/null || echo "$body"
        fi
    else
        echo -e "${RED}❌ 状态码: $http_code (期望: $expected_status)${NC}"
        echo "响应内容: $body"
    fi
    echo ""
}

# 检查服务是否运行
echo -e "${YELLOW}检查服务状态...${NC}"
if ! curl -s "$BASE_URL/health" > /dev/null; then
    echo -e "${RED}❌ 服务未运行，请先启动服务:${NC}"
    echo "   go run cmd/server/main.go"
    echo "   或者"
    echo "   ./bin/gin-weather"
    exit 1
fi
echo -e "${GREEN}✅ 服务正在运行${NC}"
echo ""

# 1. 健康检查
test_endpoint "健康检查" "$BASE_URL/health" 200

# 2. 根据城市查询天气（需要有效的 API Key）
if [ "$API_KEY" != "your_api_key_here" ]; then
    echo -e "${YELLOW}使用 API Key 测试天气查询...${NC}"
    
    # 设置环境变量
    export WEATHER_API_KEY="$API_KEY"
    
    test_endpoint "北京天气查询" "$BASE_URL/weather/city/Beijing?units=metric&lang=zh_cn" 200
    test_endpoint "上海天气查询" "$BASE_URL/weather/city/Shanghai?units=metric&lang=en" 200
    test_endpoint "坐标天气查询" "$BASE_URL/weather/coordinates/39.9042/116.4074?units=metric" 200
    test_endpoint "通用查询（城市）" "$BASE_URL/weather?city=Guangzhou&units=metric&lang=zh_cn" 200
    test_endpoint "通用查询（坐标）" "$BASE_URL/weather?lat=31.2304&lon=121.4737&units=imperial" 200
else
    echo -e "${YELLOW}⚠️  跳过天气查询测试（需要有效的 API Key）${NC}"
    echo "使用方法: $0 YOUR_API_KEY"
    echo ""
fi

# 3. 错误情况测试
echo -e "${YELLOW}测试错误处理...${NC}"
test_endpoint "无效城市参数" "$BASE_URL/weather/city/" 404
test_endpoint "无效坐标（纬度超范围）" "$BASE_URL/weather/coordinates/91/116.4074" 400
test_endpoint "无效坐标（经度超范围）" "$BASE_URL/weather/coordinates/39.9042/181" 400
test_endpoint "无效坐标格式" "$BASE_URL/weather/coordinates/invalid/116.4074" 400
test_endpoint "缺少参数" "$BASE_URL/weather" 400

# 4. 性能测试（简单）
echo -e "${YELLOW}简单性能测试...${NC}"
echo "发送 10 个并发请求到健康检查接口..."

start_time=$(date +%s.%N)
for i in {1..10}; do
    curl -s "$BASE_URL/health" > /dev/null &
done
wait
end_time=$(date +%s.%N)

duration=$(echo "$end_time - $start_time" | bc)
echo -e "${GREEN}✅ 10 个并发请求完成，耗时: ${duration}s${NC}"
echo ""

echo -e "${GREEN}🎉 测试完成！${NC}"
echo ""
echo "💡 提示:"
echo "1. 获取免费 API Key: https://openweathermap.org/api"
echo "2. 设置环境变量: export WEATHER_API_KEY=your_key"
echo "3. 查看完整 API 文档: docs/API.md"
