#!/bin/bash

# Gin Weather API 演示脚本
# 用于快速启动和演示项目功能

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# 显示标题
show_title() {
    echo -e "${CYAN}"
    echo "╔══════════════════════════════════════════════════════════════╗"
    echo "║                    Gin Weather API 演示                      ║"
    echo "║                                                              ║"
    echo "║  🌤️  现代化天气查询应用                                        ║"
    echo "║  🚀  Go + Vue.js + TypeScript + Docker                      ║"
    echo "╚══════════════════════════════════════════════════════════════╝"
    echo -e "${NC}"
}

# 检查依赖
check_dependencies() {
    echo -e "${YELLOW}🔍 检查依赖...${NC}"
    
    # 检查 Docker
    if ! command -v docker &> /dev/null; then
        echo -e "${RED}❌ Docker 未安装，请先安装 Docker${NC}"
        exit 1
    fi
    
    # 检查 Docker Compose
    if ! command -v docker-compose &> /dev/null; then
        echo -e "${RED}❌ Docker Compose 未安装，请先安装 Docker Compose${NC}"
        exit 1
    fi
    
    # 检查 Go（可选）
    if command -v go &> /dev/null; then
        echo -e "${GREEN}✅ Go $(go version | cut -d' ' -f3)${NC}"
    else
        echo -e "${YELLOW}⚠️  Go 未安装（Docker 模式不需要）${NC}"
    fi
    
    # 检查 Node.js（可选）
    if command -v node &> /dev/null; then
        echo -e "${GREEN}✅ Node.js $(node --version)${NC}"
    else
        echo -e "${YELLOW}⚠️  Node.js 未安装（Docker 模式不需要）${NC}"
    fi
    
    echo -e "${GREEN}✅ Docker $(docker --version | cut -d' ' -f3 | tr -d ',')${NC}"
    echo -e "${GREEN}✅ Docker Compose $(docker-compose --version | cut -d' ' -f3 | tr -d ',')${NC}"
}

# 检查环境变量
check_env() {
    echo -e "${YELLOW}🔧 检查环境配置...${NC}"
    
    if [ ! -f "$PROJECT_ROOT/.env" ]; then
        echo -e "${YELLOW}⚠️  .env 文件不存在，创建示例文件...${NC}"
        cp "$PROJECT_ROOT/.env.example" "$PROJECT_ROOT/.env"
        echo -e "${RED}❗ 请编辑 .env 文件，设置您的 OpenWeatherMap API 密钥${NC}"
        echo -e "${BLUE}💡 获取 API 密钥: https://openweathermap.org/api${NC}"
        read -p "按 Enter 继续（确保已设置 API 密钥）..."
    fi
    
    # 检查 API 密钥
    if grep -q "your_openweathermap_api_key_here" "$PROJECT_ROOT/.env"; then
        echo -e "${RED}❗ 请在 .env 文件中设置真实的 API 密钥${NC}"
        echo -e "${BLUE}💡 编辑命令: nano .env${NC}"
        read -p "按 Enter 继续..."
    else
        echo -e "${GREEN}✅ API 密钥已配置${NC}"
    fi
}

# Docker 模式启动
start_docker() {
    echo -e "${BLUE}🐳 使用 Docker Compose 启动服务...${NC}"
    
    cd "$PROJECT_ROOT"
    
    # 停止可能存在的容器
    docker-compose down 2>/dev/null || true
    
    # 构建并启动服务
    echo -e "${YELLOW}📦 构建镜像...${NC}"
    docker-compose build
    
    echo -e "${YELLOW}🚀 启动服务...${NC}"
    docker-compose up -d
    
    # 等待服务启动
    echo -e "${YELLOW}⏳ 等待服务启动...${NC}"
    sleep 10
    
    # 检查服务状态
    check_services_docker
}

# 开发模式启动
start_dev() {
    echo -e "${BLUE}💻 使用开发模式启动服务...${NC}"
    
    # 检查 Go 和 Node.js
    if ! command -v go &> /dev/null; then
        echo -e "${RED}❌ 开发模式需要安装 Go${NC}"
        exit 1
    fi
    
    if ! command -v node &> /dev/null; then
        echo -e "${RED}❌ 开发模式需要安装 Node.js${NC}"
        exit 1
    fi
    
    cd "$PROJECT_ROOT"
    
    # 启动后端
    echo -e "${YELLOW}🔧 启动后端服务...${NC}"
    source .env
    export WEATHER_API_KEY
    go run cmd/server/main.go &
    BACKEND_PID=$!
    
    # 等待后端启动
    sleep 3
    
    # 启动前端
    echo -e "${YELLOW}🎨 启动前端服务...${NC}"
    cd frontend
    if [ ! -d "node_modules" ]; then
        npm install
    fi
    npm run dev &
    FRONTEND_PID=$!
    
    # 等待前端启动
    sleep 5
    
    # 检查服务状态
    check_services_dev
    
    # 保存 PID 以便后续清理
    echo $BACKEND_PID > /tmp/gin-weather-backend.pid
    echo $FRONTEND_PID > /tmp/gin-weather-frontend.pid
}

# 检查 Docker 服务状态
check_services_docker() {
    echo -e "${YELLOW}🔍 检查服务状态...${NC}"
    
    # 检查后端
    if curl -s http://localhost:8080/api/v1/health > /dev/null; then
        echo -e "${GREEN}✅ 后端服务正常 (http://localhost:8080)${NC}"
    else
        echo -e "${RED}❌ 后端服务异常${NC}"
        docker-compose logs gin-weather-backend
        return 1
    fi
    
    # 检查前端
    if curl -s http://localhost:3000/health > /dev/null; then
        echo -e "${GREEN}✅ 前端服务正常 (http://localhost:3000)${NC}"
    else
        echo -e "${RED}❌ 前端服务异常${NC}"
        docker-compose logs gin-weather-frontend
        return 1
    fi
}

# 检查开发服务状态
check_services_dev() {
    echo -e "${YELLOW}🔍 检查服务状态...${NC}"
    
    # 检查后端
    if curl -s http://localhost:8080/api/v1/health > /dev/null; then
        echo -e "${GREEN}✅ 后端服务正常 (http://localhost:8080)${NC}"
    else
        echo -e "${RED}❌ 后端服务异常${NC}"
        return 1
    fi
    
    # 检查前端
    if curl -s http://localhost:5173 > /dev/null; then
        echo -e "${GREEN}✅ 前端服务正常 (http://localhost:5173)${NC}"
    else
        echo -e "${RED}❌ 前端服务异常${NC}"
        return 1
    fi
}

# 演示 API 功能
demo_api() {
    echo -e "${PURPLE}🎯 演示 API 功能...${NC}"
    
    local base_url="http://localhost:8080/api/v1"
    
    echo -e "${CYAN}1. 健康检查${NC}"
    curl -s "$base_url/health" | jq '.' || echo "请求失败"
    echo ""
    
    echo -e "${CYAN}2. 查询北京天气${NC}"
    curl -s "$base_url/weather/city/Beijing" | jq '.data.location, .data.current.temperature' || echo "请求失败"
    echo ""
    
    echo -e "${CYAN}3. 查询上海天气${NC}"
    curl -s "$base_url/weather/city/Shanghai" | jq '.data.location, .data.current.temperature' || echo "请求失败"
    echo ""
}

# 显示访问信息
show_access_info() {
    echo -e "${GREEN}"
    echo "🎉 服务启动成功！"
    echo ""
    echo "📱 前端界面:"
    if [ "$1" = "docker" ]; then
        echo "   http://localhost:3000"
    else
        echo "   http://localhost:5173"
    fi
    echo ""
    echo "🔧 后端 API:"
    echo "   http://localhost:8080"
    echo "   健康检查: http://localhost:8080/api/v1/health"
    echo ""
    echo "📚 文档:"
    echo "   项目文档: docs/"
    echo "   API 文档: docs/04-API.md"
    echo "   用户指南: docs/06-USER-GUIDE.md"
    echo ""
    echo "🛑 停止服务:"
    if [ "$1" = "docker" ]; then
        echo "   docker-compose down"
    else
        echo "   Ctrl+C 或运行 ./scripts/stop.sh"
    fi
    echo -e "${NC}"
}

# 主菜单
show_menu() {
    echo -e "${YELLOW}请选择启动模式:${NC}"
    echo "1) 🐳 Docker 模式 (推荐)"
    echo "2) 💻 开发模式"
    echo "3) 🎯 仅演示 API"
    echo "4) ❌ 退出"
    echo ""
    read -p "请输入选择 (1-4): " choice
    
    case $choice in
        1)
            start_docker
            demo_api
            show_access_info "docker"
            ;;
        2)
            start_dev
            demo_api
            show_access_info "dev"
            ;;
        3)
            echo -e "${YELLOW}请确保服务已启动${NC}"
            demo_api
            ;;
        4)
            echo -e "${GREEN}👋 再见！${NC}"
            exit 0
            ;;
        *)
            echo -e "${RED}❌ 无效选择${NC}"
            show_menu
            ;;
    esac
}

# 主函数
main() {
    show_title
    check_dependencies
    check_env
    show_menu
}

# 运行主函数
main "$@"
