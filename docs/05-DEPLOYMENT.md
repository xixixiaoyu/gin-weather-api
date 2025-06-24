# Gin Weather API 部署指南

本文档提供了 Gin Weather API 在不同环境下的详细部署指南，包括本地开发、Docker 容器、云服务器等多种部署方式。

## 部署概述

### 支持的部署方式

1. **本地开发部署** - 开发和测试环境
2. **Docker 容器部署** - 标准化容器部署
3. **云服务器部署** - 生产环境部署
4. **Kubernetes 部署** - 容器编排部署
5. **Serverless 部署** - 无服务器部署

### 系统要求

| 环境 | CPU | 内存 | 磁盘 | 网络 |
|------|-----|------|------|------|
| 开发环境 | 1 核 | 1GB | 5GB | 宽带 |
| 测试环境 | 1 核 | 2GB | 10GB | 宽带 |
| 生产环境 | 2 核 | 4GB | 20GB | 高速网络 |

## 本地开发部署

### 1. 环境准备

```bash
# 检查 Go 版本
go version  # 需要 1.21+

# 克隆项目
git clone <repository-url>
cd gin-weather

# 安装依赖
go mod download
```

### 2. 配置环境变量

```bash
# 复制环境变量文件
cp .env.example .env

# 编辑配置文件
vim .env
```

必需配置：
```env
# 天气 API 配置
WEATHER_API_KEY=your_openweathermap_api_key_here

# 服务器配置（可选）
SERVER_HOST=localhost
SERVER_PORT=8080
GIN_MODE=debug
```

### 3. 启动服务

```bash
# 方法 1: 直接运行
go run cmd/server/main.go

# 方法 2: 使用 Makefile
make run

# 方法 3: 构建后运行
make build
./bin/gin-weather
```

### 4. 验证部署

```bash
# 健康检查
curl http://localhost:8080/api/v1/health

# 测试 API
curl "http://localhost:8080/api/v1/weather/city/Beijing"
```

## Docker 容器部署

### 1. 构建 Docker 镜像

```bash
# 使用 Dockerfile 构建
docker build -t gin-weather:latest .

# 或使用 Makefile
make docker-build
```

### 2. 运行容器

```bash
# 基本运行
docker run -p 8080:8080 \
  -e WEATHER_API_KEY=your_api_key \
  gin-weather:latest

# 后台运行
docker run -d \
  --name gin-weather \
  -p 8080:8080 \
  -e WEATHER_API_KEY=your_api_key \
  --restart unless-stopped \
  gin-weather:latest
```

### 3. 使用 Docker Compose

```bash
# 创建 .env 文件
echo "WEATHER_API_KEY=your_api_key" > .env

# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### 4. Docker 镜像优化

#### 多阶段构建 Dockerfile

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-weather cmd/server/main.go

# 运行阶段
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/gin-weather .
RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1
CMD ["./gin-weather"]
```

## 云服务器部署

### 1. 服务器准备

#### Ubuntu/Debian 系统

```bash
# 更新系统
sudo apt update && sudo apt upgrade -y

# 安装必要工具
sudo apt install -y curl wget git build-essential

# 安装 Go
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### CentOS/RHEL 系统

```bash
# 更新系统
sudo yum update -y

# 安装必要工具
sudo yum install -y curl wget git gcc

# 安装 Go
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### 2. 应用部署

```bash
# 创建应用目录
sudo mkdir -p /opt/gin-weather
sudo chown $USER:$USER /opt/gin-weather
cd /opt/gin-weather

# 克隆代码
git clone <repository-url> .

# 构建应用
go build -o gin-weather cmd/server/main.go

# 创建配置文件
sudo mkdir -p /etc/gin-weather
sudo cp .env.example /etc/gin-weather/.env
sudo vim /etc/gin-weather/.env
```

### 3. 系统服务配置

#### 创建 systemd 服务文件

```bash
sudo vim /etc/systemd/system/gin-weather.service
```

服务文件内容：
```ini
[Unit]
Description=Gin Weather API Service
After=network.target

[Service]
Type=simple
User=gin-weather
Group=gin-weather
WorkingDirectory=/opt/gin-weather
ExecStart=/opt/gin-weather/gin-weather
EnvironmentFile=/etc/gin-weather/.env
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal
SyslogIdentifier=gin-weather

# 安全设置
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/gin-weather

[Install]
WantedBy=multi-user.target
```

#### 创建专用用户

```bash
# 创建系统用户
sudo useradd --system --shell /bin/false gin-weather

# 设置文件权限
sudo chown -R gin-weather:gin-weather /opt/gin-weather
sudo chmod +x /opt/gin-weather/gin-weather
```

#### 启动服务

```bash
# 重新加载 systemd
sudo systemctl daemon-reload

# 启用服务
sudo systemctl enable gin-weather

# 启动服务
sudo systemctl start gin-weather

# 检查状态
sudo systemctl status gin-weather

# 查看日志
sudo journalctl -u gin-weather -f
```

### 4. 反向代理配置

#### Nginx 配置

```bash
# 安装 Nginx
sudo apt install -y nginx  # Ubuntu/Debian
sudo yum install -y nginx  # CentOS/RHEL

# 创建配置文件
sudo vim /etc/nginx/sites-available/gin-weather
```

Nginx 配置内容：
```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 重定向到 HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;

    # SSL 证书配置
    ssl_certificate /path/to/your/certificate.crt;
    ssl_certificate_key /path/to/your/private.key;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384;
    ssl_prefer_server_ciphers off;

    # 安全头
    add_header X-Frame-Options DENY;
    add_header X-Content-Type-Options nosniff;
    add_header X-XSS-Protection "1; mode=block";
    add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload";

    # 代理配置
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
        
        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 健康检查
    location /health {
        access_log off;
        proxy_pass http://127.0.0.1:8080/api/v1/health;
    }

    # 日志配置
    access_log /var/log/nginx/gin-weather.access.log;
    error_log /var/log/nginx/gin-weather.error.log;
}
```

启用配置：
```bash
# 创建软链接
sudo ln -s /etc/nginx/sites-available/gin-weather /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重启 Nginx
sudo systemctl restart nginx
```

## Kubernetes 部署

### 1. 创建 Kubernetes 配置文件

#### Deployment 配置

```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gin-weather
  labels:
    app: gin-weather
spec:
  replicas: 3
  selector:
    matchLabels:
      app: gin-weather
  template:
    metadata:
      labels:
        app: gin-weather
    spec:
      containers:
      - name: gin-weather
        image: gin-weather:latest
        ports:
        - containerPort: 8080
        env:
        - name: WEATHER_API_KEY
          valueFrom:
            secretKeyRef:
              name: gin-weather-secret
              key: weather-api-key
        - name: SERVER_HOST
          value: "0.0.0.0"
        - name: SERVER_PORT
          value: "8080"
        - name: GIN_MODE
          value: "release"
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "200m"
        livenessProbe:
          httpGet:
            path: /api/v1/health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /api/v1/health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

#### Service 配置

```yaml
# k8s/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: gin-weather-service
spec:
  selector:
    app: gin-weather
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: ClusterIP
```

#### Secret 配置

```yaml
# k8s/secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: gin-weather-secret
type: Opaque
data:
  weather-api-key: <base64-encoded-api-key>
```

#### Ingress 配置

```yaml
# k8s/ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: gin-weather-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - your-domain.com
    secretName: gin-weather-tls
  rules:
  - host: your-domain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: gin-weather-service
            port:
              number: 80
```

### 2. 部署到 Kubernetes

```bash
# 创建 Secret
echo -n 'your_api_key' | base64
kubectl apply -f k8s/secret.yaml

# 部署应用
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml

# 检查部署状态
kubectl get pods -l app=gin-weather
kubectl get services
kubectl get ingress

# 查看日志
kubectl logs -l app=gin-weather -f
```

## 监控和日志

### 1. 日志配置

#### 结构化日志

```go
// 在应用中添加结构化日志
import (
    "github.com/sirupsen/logrus"
)

func setupLogger() *logrus.Logger {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
    logger.SetLevel(logrus.InfoLevel)
    return logger
}
```

#### 日志收集

```bash
# 使用 rsyslog 收集日志
sudo vim /etc/rsyslog.d/gin-weather.conf
```

配置内容：
```
# Gin Weather API 日志
:programname, isequal, "gin-weather" /var/log/gin-weather.log
& stop
```

### 2. 监控配置

#### Prometheus 监控

```go
// 添加 Prometheus 指标
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration in seconds",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
}

// 在路由中添加指标端点
router.GET("/metrics", gin.WrapH(promhttp.Handler()))
```

#### Grafana 仪表板

```json
{
  "dashboard": {
    "title": "Gin Weather API",
    "panels": [
      {
        "title": "Request Rate",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])",
            "legendFormat": "{{method}} {{endpoint}}"
          }
        ]
      },
      {
        "title": "Response Time",
        "type": "graph",
        "targets": [
          {
            "expr": "histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))",
            "legendFormat": "95th percentile"
          }
        ]
      }
    ]
  }
}
```

## 安全配置

### 1. 防火墙配置

```bash
# Ubuntu/Debian (ufw)
sudo ufw allow ssh
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable

# CentOS/RHEL (firewalld)
sudo firewall-cmd --permanent --add-service=ssh
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload
```

### 2. SSL/TLS 证书

#### 使用 Let's Encrypt

```bash
# 安装 Certbot
sudo apt install -y certbot python3-certbot-nginx  # Ubuntu/Debian
sudo yum install -y certbot python3-certbot-nginx  # CentOS/RHEL

# 获取证书
sudo certbot --nginx -d your-domain.com

# 自动续期
sudo crontab -e
# 添加以下行
0 12 * * * /usr/bin/certbot renew --quiet
```

### 3. 安全加固

```bash
# 禁用 root SSH 登录
sudo vim /etc/ssh/sshd_config
# 设置 PermitRootLogin no

# 配置 fail2ban
sudo apt install -y fail2ban
sudo systemctl enable fail2ban
sudo systemctl start fail2ban

# 定期更新系统
sudo apt update && sudo apt upgrade -y  # Ubuntu/Debian
sudo yum update -y  # CentOS/RHEL
```

## 备份和恢复

### 1. 配置备份

```bash
#!/bin/bash
# backup.sh
BACKUP_DIR="/backup/gin-weather"
DATE=$(date +%Y%m%d_%H%M%S)

# 创建备份目录
mkdir -p $BACKUP_DIR

# 备份配置文件
cp /etc/gin-weather/.env $BACKUP_DIR/env_$DATE
cp /etc/systemd/system/gin-weather.service $BACKUP_DIR/service_$DATE
cp /etc/nginx/sites-available/gin-weather $BACKUP_DIR/nginx_$DATE

# 备份应用文件
tar -czf $BACKUP_DIR/app_$DATE.tar.gz /opt/gin-weather

echo "备份完成: $BACKUP_DIR"
```

### 2. 自动备份

```bash
# 添加到 crontab
crontab -e
# 每天凌晨 2 点备份
0 2 * * * /path/to/backup.sh
```

## 故障排查

### 1. 常见问题

#### 服务无法启动
```bash
# 检查服务状态
sudo systemctl status gin-weather

# 查看详细日志
sudo journalctl -u gin-weather -n 50

# 检查配置文件
sudo -u gin-weather /opt/gin-weather/gin-weather
```

#### 网络连接问题
```bash
# 检查端口监听
sudo netstat -tlnp | grep :8080

# 检查防火墙
sudo ufw status
sudo firewall-cmd --list-all

# 测试网络连接
curl -v http://localhost:8080/api/v1/health
```

#### 性能问题
```bash
# 检查系统资源
top
htop
free -h
df -h

# 检查应用性能
curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/api/v1/health"
```

### 2. 日志分析

```bash
# 查看错误日志
sudo journalctl -u gin-weather | grep ERROR

# 统计请求状态
sudo journalctl -u gin-weather | grep "status=" | awk '{print $NF}' | sort | uniq -c

# 查看慢请求
sudo journalctl -u gin-weather | grep "duration" | awk '$NF > 1000'
```

通过遵循这个部署指南，您可以在各种环境中成功部署 Gin Weather API 服务，并确保其稳定、安全地运行。
