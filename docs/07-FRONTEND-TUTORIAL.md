# Gin Weather API 前端实现教程

本教程将详细介绍如何从零开始构建 Gin Weather API 的现代化前端界面，包括技术选型、项目搭建、组件开发、状态管理等各个方面。

## 目录

1. [项目概述和需求分析](#项目概述和需求分析)
2. [技术选型和架构设计](#技术选型和架构设计)
3. [项目初始化和环境搭建](#项目初始化和环境搭建)
4. [基础配置和工具链](#基础配置和工具链)
5. [类型定义和 API 服务](#类型定义和-api-服务)
6. [状态管理实现](#状态管理实现)
7. [核心组件开发](#核心组件开发)
8. [页面布局和路由](#页面布局和路由)
9. [样式设计和响应式](#样式设计和响应式)
10. [测试和质量保证](#测试和质量保证)
11. [构建和部署](#构建和部署)
12. [性能优化和最佳实践](#性能优化和最佳实践)

## 项目概述和需求分析

### 功能需求

我们需要构建一个现代化的天气查询前端应用，具备以下功能：

1. **天气查询功能**
   - 支持城市名称搜索
   - 支持地理位置定位
   - 显示详细的天气信息

2. **用户体验功能**
   - 响应式设计，支持多设备
   - 加载状态和错误处理
   - 搜索历史管理
   - 个性化设置

3. **界面设计要求**
   - 现代化的 UI 设计
   - 直观的交互体验
   - 美观的天气图标和动效
   - 友好的错误提示

### 技术需求

1. **现代前端框架**：Vue.js 3 + TypeScript
2. **构建工具**：Vite（快速开发和构建）
3. **样式方案**：UnoCSS（原子化 CSS）
4. **状态管理**：Pinia（Vue 3 推荐）
5. **HTTP 客户端**：Axios（API 通信）
6. **测试框架**：Vitest（单元测试）

## 技术选型和架构设计

### 技术栈选择

#### 前端框架：Vue.js 3
**选择理由**：
- **Composition API**：更好的逻辑复用和类型推导
- **性能优化**：更小的包体积和更快的渲染
- **TypeScript 支持**：原生 TypeScript 支持
- **生态成熟**：丰富的生态系统和社区支持

#### 构建工具：Vite
**选择理由**：
- **快速启动**：基于 ESM 的快速冷启动
- **热更新**：极快的热模块替换
- **现代化**：原生支持 TypeScript、JSX 等
- **插件生态**：丰富的插件系统

#### 样式方案：UnoCSS
**选择理由**：
- **原子化 CSS**：高度可复用的样式类
- **按需生成**：只生成使用的样式
- **高性能**：极快的编译速度
- **灵活配置**：强大的自定义能力

#### 状态管理：Pinia
**选择理由**：
- **Vue 3 原生**：专为 Vue 3 设计
- **TypeScript 友好**：完整的类型推导
- **简洁 API**：比 Vuex 更简单的 API
- **开发工具**：优秀的开发者体验

### 架构设计

```
┌─────────────────────────────────────────────────────────────┐
│                        前端架构                              │
├─────────────────────────────────────────────────────────────┤
│  View Layer (视图层)                                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ HomeView    │  │ Components  │  │ Router              │  │
│  │             │  │ - WeatherCard│  │ - Route Guards      │  │
│  │             │  │ - Search    │  │ - Lazy Loading      │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
├─────────────────────────────────────────────────────────────┤
│  State Layer (状态层)                                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Weather     │  │ User        │  │ App                 │  │
│  │ Store       │  │ Preferences │  │ State               │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
├─────────────────────────────────────────────────────────────┤
│  Service Layer (服务层)                                     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Weather API │  │ Geolocation │  │ Local Storage       │  │
│  │ Service     │  │ Service     │  │ Service             │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
├─────────────────────────────────────────────────────────────┤
│  Utils Layer (工具层)                                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ Weather     │  │ Validation  │  │ Format              │  │
│  │ Utils       │  │ Utils       │  │ Utils               │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## 项目初始化和环境搭建

### 第一步：创建 Vue 项目

```bash
# 创建前端目录
mkdir frontend
cd frontend

# 使用 Vue CLI 创建项目
npm create vue@latest . -- --typescript --router --pinia --vitest --eslint --prettier
```

**配置选项说明**：
- `--typescript`：启用 TypeScript 支持
- `--router`：添加 Vue Router 路由
- `--pinia`：添加 Pinia 状态管理
- `--vitest`：添加 Vitest 测试框架
- `--eslint`：添加 ESLint 代码检查
- `--prettier`：添加 Prettier 代码格式化

### 第二步：安装额外依赖

```bash
# 安装核心依赖
npm install @unocss/vite unocss axios @vueuse/core lucide-vue-next

# 安装图标依赖
npm install @iconify-json/lucide
```

**依赖包说明**：
- `@unocss/vite`：UnoCSS Vite 插件
- `unocss`：UnoCSS 核心库
- `axios`：HTTP 客户端
- `@vueuse/core`：Vue 组合式工具库
- `lucide-vue-next`：现代图标库
- `@iconify-json/lucide`：Lucide 图标数据

### 第三步：项目结构规划

```
frontend/
├── public/                 # 静态资源
├── src/
│   ├── components/        # 可复用组件
│   │   ├── WeatherCard.vue      # 天气卡片组件
│   │   ├── WeatherSearch.vue    # 搜索组件
│   │   ├── LoadingSpinner.vue   # 加载动画组件
│   │   ├── ErrorMessage.vue     # 错误提示组件
│   │   └── SettingsPanel.vue    # 设置面板组件
│   ├── services/          # API 服务层
│   │   ├── weatherApi.ts        # 天气 API 服务
│   │   └── geolocation.ts       # 地理位置服务
│   ├── stores/            # Pinia 状态管理
│   │   └── weather.ts           # 天气状态管理
│   ├── types/             # TypeScript 类型定义
│   │   └── weather.ts           # 天气相关类型
│   ├── utils/             # 工具函数
│   │   └── weather.ts           # 天气工具函数
│   ├── views/             # 页面组件
│   │   └── HomeView.vue         # 主页面
│   ├── App.vue            # 根组件
│   └── main.ts            # 应用入口
├── Dockerfile             # Docker 配置
├── nginx.conf             # Nginx 配置
├── uno.config.ts          # UnoCSS 配置
├── vite.config.ts         # Vite 配置
└── package.json           # 项目配置
```

## 基础配置和工具链

### UnoCSS 配置

创建 `uno.config.ts`：

```typescript
import { defineConfig, presetUno, presetAttributify, presetIcons } from 'unocss'

export default defineConfig({
  presets: [
    presetUno(),           // 基础工具类
    presetAttributify(),   // 属性化模式
    presetIcons({          // 图标支持
      collections: {
        lucide: () => import('@iconify-json/lucide/icons.json').then(i => i.default),
      },
    }),
  ],
  theme: {
    colors: {
      primary: {
        50: '#eff6ff',
        500: '#3b82f6',
        600: '#2563eb',
        // ... 更多颜色定义
      },
    },
  },
  shortcuts: {
    // 自定义快捷样式
    'btn': 'px-4 py-2 rounded-lg font-medium transition-all duration-200',
    'btn-primary': 'btn bg-primary-500 text-white hover:bg-primary-600',
    'card': 'bg-white rounded-xl shadow-lg border border-gray-100',
    'input': 'px-4 py-3 rounded-lg border border-gray-300 focus:border-primary-500',
  },
})
```

**配置要点**：
- **预设组合**：基础工具类 + 属性化 + 图标
- **主题定制**：自定义颜色系统
- **快捷样式**：常用组件样式封装

### Vite 配置

更新 `vite.config.ts`：

```typescript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from '@unocss/vite'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig({
  plugins: [
    vue(),
    UnoCSS(),  // 添加 UnoCSS 插件
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    proxy: {
      // 开发环境 API 代理
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

**配置要点**：
- **插件集成**：Vue + UnoCSS
- **路径别名**：`@` 指向 `src` 目录
- **开发代理**：代理 API 请求到后端

### 主入口配置

更新 `src/main.ts`：

```typescript
import './assets/main.css'
import 'uno.css'  // 引入 UnoCSS

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
```

**要点说明**：
- **样式引入**：UnoCSS 样式文件
- **插件注册**：Pinia + Router
- **应用挂载**：挂载到 DOM

## 类型定义和 API 服务

### 第一步：定义 TypeScript 类型

创建 `src/types/weather.ts`：

```typescript
// 天气请求参数类型
export interface WeatherRequest {
  city?: string
  lat?: number
  lon?: number
  units?: 'metric' | 'imperial' | 'standard'
  lang?: string
}

// 位置信息类型
export interface Location {
  name: string
  country: string
  latitude: number
  longitude: number
  timezone: number
}

// 天气状况类型
export interface Weather {
  id: number
  main: string
  description: string
  icon: string
}

// 当前天气类型
export interface Current {
  temperature: number
  feels_like: number
  temp_min: number
  temp_max: number
  pressure: number
  humidity: number
  visibility: number
  uv_index: number
  weather: Weather[]
  wind: {
    speed: number
    direction: number
    gust: number
  }
  clouds: {
    all: number
  }
  sunrise: number
  sunset: number
  updated_at: string
}

// 天气响应类型
export interface WeatherResponse {
  location: Location
  current: Current
  timestamp: number
  provider: string
}

// API 通用响应类型
export interface APIResponse<T = any> {
  success: boolean
  data?: T
  error?: {
    error: string
    code: number
    message?: string
  }
}
```

**类型设计原则**：
- **完整性**：覆盖所有 API 字段
- **可选性**：合理使用可选属性
- **泛型支持**：通用响应类型
- **嵌套结构**：逻辑清晰的嵌套

### 第二步：实现 API 服务

创建 `src/services/weatherApi.ts`：

```typescript
import axios from 'axios'
import type { WeatherRequest, WeatherResponse, APIResponse } from '@/types/weather'

// 创建 axios 实例
const api = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    console.log('发起请求:', config.method?.toUpperCase(), config.url)
    return config
  },
  (error) => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    console.log('响应成功:', response.status, response.config.url)
    return response
  },
  (error) => {
    console.error('响应错误:', error.response?.status, error.response?.data || error.message)
    return Promise.reject(error)
  }
)

export class WeatherApiService {
  // 健康检查
  static async healthCheck(): Promise<APIResponse> {
    try {
      const response = await api.get('/health')
      return response.data
    } catch (error) {
      throw this.handleError(error)
    }
  }

  // 根据城市查询天气
  static async getWeatherByCity(
    city: string,
    units: 'metric' | 'imperial' | 'standard' = 'metric',
    lang: string = 'zh_cn'
  ): Promise<WeatherResponse> {
    try {
      const response = await api.get<APIResponse<WeatherResponse>>(
        `/weather/city/${encodeURIComponent(city)}`,
        { params: { units, lang } }
      )

      if (response.data.success && response.data.data) {
        return response.data.data
      } else {
        throw new Error(response.data.error?.message || '获取天气数据失败')
      }
    } catch (error) {
      throw this.handleError(error)
    }
  }

  // 根据坐标查询天气
  static async getWeatherByCoordinates(
    lat: number,
    lon: number,
    units: 'metric' | 'imperial' | 'standard' = 'metric',
    lang: string = 'zh_cn'
  ): Promise<WeatherResponse> {
    try {
      const response = await api.get<APIResponse<WeatherResponse>>(
        `/weather/coordinates/${lat}/${lon}`,
        { params: { units, lang } }
      )

      if (response.data.success && response.data.data) {
        return response.data.data
      } else {
        throw new Error(response.data.error?.message || '获取天气数据失败')
      }
    } catch (error) {
      throw this.handleError(error)
    }
  }

  // 错误处理
  private static handleError(error: any): Error {
    if (axios.isAxiosError(error)) {
      if (error.response) {
        const errorData = error.response.data
        if (errorData?.error?.message) {
          return new Error(errorData.error.message)
        }
        return new Error(`请求失败: ${error.response.status} ${error.response.statusText}`)
      } else if (error.request) {
        return new Error('网络连接失败，请检查网络设置')
      }
    }
    return new Error(error.message || '未知错误')
  }
}
```

**API 服务设计要点**：
- **实例配置**：统一的 axios 配置
- **拦截器**：请求和响应的统一处理
- **错误处理**：完善的错误分类和处理
- **类型安全**：完整的 TypeScript 类型支持

### 第三步：地理位置服务

创建 `src/services/geolocation.ts`：

```typescript
import type { GeolocationPosition, GeolocationError } from '@/types/weather'

export class GeolocationService {
  // 获取当前位置
  static getCurrentPosition(): Promise<GeolocationPosition> {
    return new Promise((resolve, reject) => {
      if (!navigator.geolocation) {
        reject(new Error('浏览器不支持地理位置功能'))
        return
      }

      const options: PositionOptions = {
        enableHighAccuracy: true,
        timeout: 10000,
        maximumAge: 300000, // 5分钟缓存
      }

      navigator.geolocation.getCurrentPosition(
        (position) => {
          resolve({
            coords: {
              latitude: position.coords.latitude,
              longitude: position.coords.longitude,
              accuracy: position.coords.accuracy,
            },
          })
        },
        (error) => {
          reject(this.handleGeolocationError(error))
        },
        options
      )
    })
  }

  // 错误处理
  private static handleGeolocationError(error: GeolocationPositionError): GeolocationError {
    switch (error.code) {
      case error.PERMISSION_DENIED:
        return { code: error.code, message: '用户拒绝了地理位置请求' }
      case error.POSITION_UNAVAILABLE:
        return { code: error.code, message: '位置信息不可用' }
      case error.TIMEOUT:
        return { code: error.code, message: '获取位置信息超时' }
      default:
        return { code: error.code, message: '获取位置信息时发生未知错误' }
    }
  }
}
```

**地理位置服务特点**：
- **Promise 封装**：将回调 API 转换为 Promise
- **错误分类**：详细的错误类型处理
- **配置优化**：合理的精度和缓存设置

## 状态管理实现

### Pinia Store 设计

创建 `src/stores/weather.ts`：

```typescript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { WeatherResponse, LoadingState } from '@/types/weather'
import WeatherApiService from '@/services/weatherApi'
import GeolocationService from '@/services/geolocation'

export const useWeatherStore = defineStore('weather', () => {
  // 状态定义
  const currentWeather = ref<WeatherResponse | null>(null)
  const loadingState = ref<LoadingState>({
    isLoading: false,
    error: null,
  })

  // 用户偏好设置
  const preferences = ref({
    units: 'metric' as 'metric' | 'imperial' | 'standard',
    language: 'zh_cn',
    autoLocation: false,
  })

  // 搜索历史
  const searchHistory = ref<string[]>([])

  // 计算属性
  const isLoading = computed(() => loadingState.value.isLoading)
  const error = computed(() => loadingState.value.error)
  const hasWeatherData = computed(() => currentWeather.value !== null)

  // Actions
  const setLoading = (loading: boolean) => {
    loadingState.value.isLoading = loading
    if (loading) {
      loadingState.value.error = null
    }
  }

  const setError = (error: string | null) => {
    loadingState.value.error = error
    loadingState.value.isLoading = false
  }

  // 根据城市获取天气
  const fetchWeatherByCity = async (city: string) => {
    try {
      setLoading(true)
      const weather = await WeatherApiService.getWeatherByCity(
        city,
        preferences.value.units,
        preferences.value.language
      )

      currentWeather.value = weather
      addToSearchHistory(city)
      setLoading(false)

      return weather
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : '获取天气数据失败'
      setError(errorMessage)
      throw error
    }
  }

  // 根据坐标获取天气
  const fetchWeatherByCoordinates = async (lat: number, lon: number) => {
    try {
      setLoading(true)
      const weather = await WeatherApiService.getWeatherByCoordinates(
        lat,
        lon,
        preferences.value.units,
        preferences.value.language
      )

      currentWeather.value = weather
      setLoading(false)

      return weather
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : '获取天气数据失败'
      setError(errorMessage)
      throw error
    }
  }

  // 获取当前位置天气
  const fetchWeatherByCurrentLocation = async () => {
    try {
      setLoading(true)

      const position = await GeolocationService.getCurrentPosition()
      const weather = await fetchWeatherByCoordinates(
        position.coords.latitude,
        position.coords.longitude
      )

      return weather
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : '获取位置信息失败'
      setError(errorMessage)
      throw error
    }
  }

  // 搜索历史管理
  const addToSearchHistory = (city: string) => {
    const trimmedCity = city.trim()
    if (!trimmedCity) return

    const filtered = searchHistory.value.filter(item =>
      item.toLowerCase() !== trimmedCity.toLowerCase()
    )

    searchHistory.value = [trimmedCity, ...filtered].slice(0, 10)
    saveSearchHistory()
  }

  // 本地存储
  const saveSearchHistory = () => {
    localStorage.setItem('weather-search-history', JSON.stringify(searchHistory.value))
  }

  const loadSearchHistory = () => {
    try {
      const saved = localStorage.getItem('weather-search-history')
      if (saved) {
        searchHistory.value = JSON.parse(saved)
      }
    } catch (error) {
      console.error('加载搜索历史失败:', error)
    }
  }

  // 初始化
  const initialize = () => {
    loadSearchHistory()
    // 其他初始化逻辑
  }

  return {
    // 状态
    currentWeather,
    loadingState,
    preferences,
    searchHistory,

    // 计算属性
    isLoading,
    error,
    hasWeatherData,

    // 方法
    setLoading,
    setError,
    fetchWeatherByCity,
    fetchWeatherByCoordinates,
    fetchWeatherByCurrentLocation,
    addToSearchHistory,
    initialize,
  }
})
```

**Pinia Store 设计要点**：
- **Composition API 风格**：使用 `ref` 和 `computed`
- **状态分离**：数据状态、UI 状态、用户偏好分离
- **异步处理**：完善的异步操作和错误处理
- **本地存储**：搜索历史等数据持久化

## 核心组件开发

### 第一步：天气卡片组件

创建 `src/components/WeatherCard.vue`：

```vue
<template>
  <div
    class="card p-6 min-h-96 relative overflow-hidden"
    :class="backgroundClass"
  >
    <!-- 背景装饰 -->
    <div class="absolute inset-0 opacity-10">
      <div class="absolute top-4 right-4 w-32 h-32 rounded-full bg-white/20"></div>
      <div class="absolute bottom-8 left-8 w-24 h-24 rounded-full bg-white/10"></div>
    </div>

    <!-- 内容 -->
    <div class="relative z-10 text-white">
      <!-- 位置信息 -->
      <div class="mb-6">
        <h2 class="text-2xl font-bold mb-1">{{ weather.location.name }}</h2>
        <p class="text-white/80 text-sm">{{ weather.location.country }}</p>
        <p class="text-white/60 text-xs mt-1">
          {{ new Date(weather.current.updated_at).toLocaleString('zh-CN') }}
        </p>
      </div>

      <!-- 主要天气信息 -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <div class="text-5xl font-light mb-2">
            {{ formatTemperature(weather.current.temperature, units) }}
          </div>
          <div class="text-white/80 text-sm">
            体感 {{ formatTemperature(weather.current.feels_like, units) }}
          </div>
        </div>

        <div class="text-right">
          <div class="flex items-center justify-end mb-2">
            <component
              :is="weatherIcon"
              class="w-16 h-16 text-white/90"
            />
          </div>
          <div class="text-white/90 text-sm font-medium">
            {{ getWeatherDescription(weather.current.weather[0]?.description || '') }}
          </div>
        </div>
      </div>

      <!-- 详细信息网格 -->
      <div class="grid grid-cols-2 gap-4 text-sm">
        <!-- 温度范围 -->
        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Thermometer class="w-4 h-4 mr-2" />
            <span class="text-white/80">温度范围</span>
          </div>
          <div class="font-medium">
            {{ formatTemperature(weather.current.temp_min, units) }} -
            {{ formatTemperature(weather.current.temp_max, units) }}
          </div>
        </div>

        <!-- 湿度 -->
        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Droplets class="w-4 h-4 mr-2" />
            <span class="text-white/80">湿度</span>
          </div>
          <div class="font-medium">{{ formatHumidity(weather.current.humidity) }}</div>
        </div>

        <!-- 更多信息项... -->
      </div>

      <!-- 日出日落 -->
      <div class="flex justify-between mt-6 pt-4 border-t border-white/20">
        <div class="flex items-center">
          <Sunrise class="w-4 h-4 mr-2" />
          <span class="text-white/80 text-sm">日出</span>
          <span class="ml-2 font-medium">{{ formatTime(weather.current.sunrise) }}</span>
        </div>
        <div class="flex items-center">
          <Sunset class="w-4 h-4 mr-2" />
          <span class="text-white/80 text-sm">日落</span>
          <span class="ml-2 font-medium">{{ formatTime(weather.current.sunset) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import {
  Thermometer,
  Droplets,
  Sunrise,
  Sunset,
  Cloud,
  Sun,
} from 'lucide-vue-next'
import type { WeatherResponse } from '@/types/weather'
import {
  formatTemperature,
  formatHumidity,
  formatTime,
  getWeatherDescription,
  getWeatherBackground,
  getWeatherIcon,
} from '@/utils/weather'

interface Props {
  weather: WeatherResponse
  units?: 'metric' | 'imperial' | 'standard'
}

const props = withDefaults(defineProps<Props>(), {
  units: 'metric'
})

// 天气图标组件映射
const iconComponents = {
  sun: Sun,
  cloud: Cloud,
  // ... 更多图标映射
}

const weatherIcon = computed(() => {
  const iconName = getWeatherIcon(props.weather.current.weather[0]?.icon || '01d')
  return iconComponents[iconName as keyof typeof iconComponents] || Cloud
})

const backgroundClass = computed(() => {
  const weatherMain = props.weather.current.weather[0]?.main || 'clear'
  const gradient = getWeatherBackground(weatherMain)
  return `bg-gradient-to-br ${gradient}`
})
</script>
```

**组件设计要点**：
- **Props 类型定义**：完整的 TypeScript 类型支持
- **动态样式**：根据天气状况动态背景
- **响应式布局**：Grid 布局适配不同屏幕
- **图标系统**：动态图标组件映射

### 第二步：搜索组件

创建 `src/components/WeatherSearch.vue`：

```vue
<template>
  <div class="w-full max-w-md mx-auto">
    <!-- 搜索输入框 -->
    <div class="relative">
      <div class="relative">
        <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="输入城市名称..."
          class="input w-full pl-10 pr-12"
          :class="{ 'border-red-300 focus:border-red-500': hasError }"
          @keyup.enter="handleSearch"
          @input="handleInput"
          @focus="showSuggestions = true"
        />
        <button
          v-if="searchQuery"
          @click="clearSearch"
          class="absolute right-10 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
        >
          <X class="w-4 h-4" />
        </button>
        <button
          @click="handleLocationSearch"
          :disabled="isLoadingLocation"
          class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-primary-500"
          title="使用当前位置"
        >
          <MapPin v-if="!isLoadingLocation" class="w-4 h-4" />
          <Loader2 v-else class="w-4 h-4 animate-spin" />
        </button>
      </div>

      <!-- 错误提示 -->
      <div v-if="hasError" class="mt-2 text-sm text-red-600">
        {{ errorMessage }}
      </div>

      <!-- 搜索建议下拉框 -->
      <div
        v-if="showSuggestions && (filteredHistory.length > 0 || searchQuery)"
        class="absolute top-full left-0 right-0 mt-1 bg-white rounded-lg shadow-lg border z-50"
      >
        <!-- 搜索历史 -->
        <div v-if="filteredHistory.length > 0" class="p-2">
          <div class="text-xs text-gray-500 px-2 py-1 mb-1">搜索历史</div>
          <button
            v-for="city in filteredHistory"
            :key="city"
            @click="selectCity(city)"
            class="w-full text-left px-3 py-2 rounded-md hover:bg-gray-50 flex items-center justify-between group"
          >
            <div class="flex items-center">
              <Clock class="w-4 h-4 text-gray-400 mr-2" />
              <span>{{ city }}</span>
            </div>
            <button
              @click.stop="removeFromHistory(city)"
              class="opacity-0 group-hover:opacity-100 text-gray-400 hover:text-red-500"
            >
              <X class="w-3 h-3" />
            </button>
          </button>
        </div>

        <!-- 当前搜索项 -->
        <div v-if="searchQuery && !filteredHistory.includes(searchQuery)" class="p-2 border-t">
          <button
            @click="handleSearch"
            class="w-full text-left px-3 py-2 rounded-md hover:bg-gray-50 flex items-center"
          >
            <Search class="w-4 h-4 text-gray-400 mr-2" />
            <span>搜索 "{{ searchQuery }}"</span>
          </button>
        </div>
      </div>
    </div>

    <!-- 快捷操作按钮 -->
    <div class="flex gap-2 mt-4">
      <button
        @click="handleLocationSearch"
        :disabled="isLoadingLocation"
        class="btn-secondary flex items-center gap-2 flex-1"
      >
        <MapPin v-if="!isLoadingLocation" class="w-4 h-4" />
        <Loader2 v-else class="w-4 h-4 animate-spin" />
        <span>{{ isLoadingLocation ? '定位中...' : '当前位置' }}</span>
      </button>

      <button
        @click="handleSearch"
        :disabled="!searchQuery.trim() || isLoading"
        class="btn-primary flex items-center gap-2 flex-1"
      >
        <Search v-if="!isLoading" class="w-4 h-4" />
        <Loader2 v-else class="w-4 h-4 animate-spin" />
        <span>{{ isLoading ? '搜索中...' : '搜索' }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { Search, MapPin, X, Clock, Loader2 } from 'lucide-vue-next'
import { useWeatherStore } from '@/stores/weather'
import { validateCityName } from '@/utils/weather'

interface Emits {
  (e: 'search', city: string): void
  (e: 'location-search'): void
}

const emit = defineEmits<Emits>()
const weatherStore = useWeatherStore()

// 响应式数据
const searchQuery = ref('')
const showSuggestions = ref(false)
const isLoadingLocation = ref(false)
const errorMessage = ref('')

// 计算属性
const isLoading = computed(() => weatherStore.isLoading)
const searchHistory = computed(() => weatherStore.searchHistory)
const hasError = computed(() => !!errorMessage.value)

const filteredHistory = computed(() => {
  if (!searchQuery.value) {
    return searchHistory.value.slice(0, 5)
  }
  return searchHistory.value.filter(city =>
    city.toLowerCase().includes(searchQuery.value.toLowerCase())
  ).slice(0, 5)
})

// 方法
const handleInput = () => {
  errorMessage.value = ''
}

const handleSearch = () => {
  const city = searchQuery.value.trim()

  if (!city) {
    errorMessage.value = '请输入城市名称'
    return
  }

  if (!validateCityName(city)) {
    errorMessage.value = '城市名称格式不正确'
    return
  }

  emit('search', city)
  showSuggestions.value = false
}

const handleLocationSearch = async () => {
  isLoadingLocation.value = true
  try {
    emit('location-search')
  } finally {
    isLoadingLocation.value = false
  }
}

const selectCity = (city: string) => {
  searchQuery.value = city
  showSuggestions.value = false
  emit('search', city)
}

const clearSearch = () => {
  searchQuery.value = ''
  errorMessage.value = ''
}

const removeFromHistory = (city: string) => {
  weatherStore.removeFromSearchHistory(city)
}

// 点击外部关闭建议框
const handleClickOutside = (event: Event) => {
  const target = event.target as Element
  if (!target.closest('.relative')) {
    showSuggestions.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>
```

**搜索组件特点**：
- **智能建议**：搜索历史和实时建议
- **输入验证**：实时验证和错误提示
- **地理定位**：集成地理位置功能
- **用户体验**：加载状态和交互反馈

### 第三步：加载和错误组件

创建 `src/components/LoadingSpinner.vue`：

```vue
<template>
  <div class="flex flex-col items-center justify-center p-8">
    <!-- 加载动画 -->
    <div class="relative">
      <div class="w-16 h-16 border-4 border-primary-200 rounded-full animate-spin border-t-primary-500"></div>
      <div class="absolute top-2 left-2 w-12 h-12 border-4 border-primary-100 rounded-full animate-spin animate-reverse border-t-primary-400"></div>
      <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
        <Cloud class="w-6 h-6 text-primary-500 animate-pulse" />
      </div>
    </div>

    <!-- 加载文本 -->
    <div class="mt-4 text-center">
      <p class="text-gray-600 font-medium">{{ message }}</p>
      <p v-if="subMessage" class="text-gray-400 text-sm mt-1">{{ subMessage }}</p>
    </div>

    <!-- 加载进度点 -->
    <div class="flex space-x-1 mt-4">
      <div
        v-for="i in 3"
        :key="i"
        class="w-2 h-2 bg-primary-400 rounded-full animate-bounce"
        :style="{ animationDelay: `${(i - 1) * 0.2}s` }"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Cloud } from 'lucide-vue-next'

interface Props {
  message?: string
  subMessage?: string
}

withDefaults(defineProps<Props>(), {
  message: '正在加载天气数据...',
  subMessage: '请稍候'
})
</script>

<style scoped>
@keyframes reverse {
  from { transform: rotate(360deg); }
  to { transform: rotate(0deg); }
}

.animate-reverse {
  animation: reverse 1s linear infinite;
}
</style>
```

创建 `src/components/ErrorMessage.vue`：

```vue
<template>
  <div class="card p-6 text-center">
    <!-- 错误图标 -->
    <div class="flex justify-center mb-4">
      <div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center">
        <component :is="errorIcon" class="w-8 h-8 text-red-500" />
      </div>
    </div>

    <!-- 错误标题 -->
    <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ title }}</h3>

    <!-- 错误消息 -->
    <p class="text-gray-600 mb-6">{{ message }}</p>

    <!-- 操作按钮 -->
    <div class="flex flex-col sm:flex-row gap-3 justify-center">
      <button
        v-if="showRetry"
        @click="$emit('retry')"
        class="btn-primary flex items-center gap-2"
      >
        <RefreshCw class="w-4 h-4" />
        重试
      </button>

      <button
        v-if="showDismiss"
        @click="$emit('dismiss')"
        class="btn-secondary flex items-center gap-2"
      >
        <X class="w-4 h-4" />
        关闭
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { AlertTriangle, Wifi, MapPin, Search, Server, RefreshCw, X } from 'lucide-vue-next'

interface Props {
  type?: 'network' | 'location' | 'search' | 'server' | 'general'
  title?: string
  message: string
  showRetry?: boolean
  showDismiss?: boolean
}

interface Emits {
  (e: 'retry'): void
  (e: 'dismiss'): void
}

const props = withDefaults(defineProps<Props>(), {
  type: 'general',
  showRetry: true,
  showDismiss: true
})

defineEmits<Emits>()

const errorIcons = {
  network: Wifi,
  location: MapPin,
  search: Search,
  server: Server,
  general: AlertTriangle,
}

const errorIcon = computed(() => errorIcons[props.type])

const defaultTitles = {
  network: '网络连接错误',
  location: '位置获取失败',
  search: '搜索失败',
  server: '服务器错误',
  general: '出现错误',
}

const title = computed(() => props.title || defaultTitles[props.type])
</script>
```

**加载和错误组件特点**：
- **视觉反馈**：清晰的加载和错误状态
- **可配置性**：支持自定义消息和操作
- **一致性**：统一的设计风格
- **交互性**：支持重试和关闭操作

## 页面布局和路由

### 主页面实现

更新 `src/views/HomeView.vue`：

```vue
<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
    <!-- 头部 -->
    <header class="bg-white/80 backdrop-blur-sm border-b border-white/20 sticky top-0 z-40">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
              <Cloud class="w-5 h-5 text-white" />
            </div>
            <h1 class="text-xl font-bold text-gray-900">天气查询</h1>
          </div>

          <div class="flex items-center space-x-2">
            <button
              @click="refreshWeather"
              :disabled="isLoading"
              class="p-2 text-gray-600 hover:text-primary-500 disabled:opacity-50"
              title="刷新"
            >
              <RefreshCw class="w-5 h-5" :class="{ 'animate-spin': isLoading }" />
            </button>

            <button
              @click="showSettings = !showSettings"
              class="p-2 text-gray-600 hover:text-primary-500"
              title="设置"
            >
              <Settings class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- 主要内容 -->
    <main class="container mx-auto px-4 py-8">
      <div class="max-w-4xl mx-auto">
        <!-- 搜索区域 -->
        <div class="mb-8">
          <WeatherSearch
            @search="handleCitySearch"
            @location-search="handleLocationSearch"
          />
        </div>

        <!-- 内容区域 -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- 主要天气卡片 -->
          <div class="lg:col-span-2">
            <!-- 加载状态 -->
            <LoadingSpinner
              v-if="isLoading && !hasWeatherData"
              :message="loadingMessage"
              sub-message="请稍候片刻"
            />

            <!-- 错误状态 -->
            <ErrorMessage
              v-else-if="error && !hasWeatherData"
              :type="errorType"
              :message="error"
              @retry="handleRetry"
              @dismiss="clearError"
            />

            <!-- 天气数据 -->
            <WeatherCard
              v-else-if="hasWeatherData"
              :weather="currentWeather!"
              :units="preferences.units"
            />

            <!-- 欢迎界面 -->
            <div v-else class="card p-8 text-center">
              <div class="w-24 h-24 bg-gradient-to-br from-blue-400 to-purple-600 rounded-full mx-auto mb-6 flex items-center justify-center">
                <Cloud class="w-12 h-12 text-white" />
              </div>
              <h2 class="text-2xl font-bold text-gray-900 mb-4">欢迎使用天气查询</h2>
              <p class="text-gray-600 mb-6">
                输入城市名称或使用当前位置来查询实时天气信息
              </p>
              <div class="flex flex-col sm:flex-row gap-3 justify-center">
                <button
                  @click="handleLocationSearch"
                  class="btn-primary flex items-center gap-2"
                >
                  <MapPin class="w-4 h-4" />
                  使用当前位置
                </button>
                <button
                  @click="searchPopularCity"
                  class="btn-secondary flex items-center gap-2"
                >
                  <Search class="w-4 h-4" />
                  查看热门城市
                </button>
              </div>
            </div>
          </div>

          <!-- 侧边栏 -->
          <div class="space-y-6">
            <!-- 设置面板 -->
            <SettingsPanel
              v-if="showSettings"
              @close="showSettings = false"
            />

            <!-- 快速搜索 -->
            <div v-else class="card p-4">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">快速搜索</h3>
              <div class="space-y-2">
                <button
                  v-for="city in popularCities"
                  :key="city"
                  @click="handleCitySearch(city)"
                  class="w-full text-left p-2 rounded-lg hover:bg-gray-50 text-sm"
                >
                  {{ city }}
                </button>
              </div>
            </div>

            <!-- 搜索历史 -->
            <div v-if="searchHistory.length > 0 && !showSettings" class="card p-4">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">搜索历史</h3>
              <div class="space-y-2">
                <button
                  v-for="city in searchHistory.slice(0, 5)"
                  :key="city"
                  @click="handleCitySearch(city)"
                  class="w-full text-left p-2 rounded-lg hover:bg-gray-50 text-sm flex items-center justify-between group"
                >
                  <span>{{ city }}</span>
                  <button
                    @click.stop="removeFromHistory(city)"
                    class="opacity-0 group-hover:opacity-100 text-gray-400 hover:text-red-500"
                  >
                    <X class="w-3 h-3" />
                  </button>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- 页脚 -->
    <footer class="bg-white/50 backdrop-blur-sm border-t border-white/20 mt-12">
      <div class="container mx-auto px-4 py-6">
        <div class="text-center text-sm text-gray-600">
          <p>数据来源: OpenWeatherMap |
            <a href="https://github.com/your-username/gin-weather" class="text-primary-500 hover:text-primary-600">
              GitHub
            </a>
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { Cloud, RefreshCw, Settings, MapPin, Search, X } from 'lucide-vue-next'
import { useWeatherStore } from '@/stores/weather'
import WeatherCard from '@/components/WeatherCard.vue'
import WeatherSearch from '@/components/WeatherSearch.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import ErrorMessage from '@/components/ErrorMessage.vue'
import SettingsPanel from '@/components/SettingsPanel.vue'

// Store
const weatherStore = useWeatherStore()

// 响应式数据
const showSettings = ref(false)
const lastSearchType = ref<'city' | 'location' | null>(null)
const lastSearchValue = ref<string | null>(null)

// 计算属性
const isLoading = computed(() => weatherStore.isLoading)
const error = computed(() => weatherStore.error)
const hasWeatherData = computed(() => weatherStore.hasWeatherData)
const currentWeather = computed(() => weatherStore.currentWeather)
const preferences = computed(() => weatherStore.preferences)
const searchHistory = computed(() => weatherStore.searchHistory)

const loadingMessage = computed(() => {
  if (lastSearchType.value === 'location') {
    return '正在获取位置信息...'
  }
  return '正在获取天气数据...'
})

const errorType = computed(() => {
  if (!error.value) return 'general'

  const errorMsg = error.value.toLowerCase()
  if (errorMsg.includes('网络') || errorMsg.includes('连接')) return 'network'
  if (errorMsg.includes('位置') || errorMsg.includes('定位')) return 'location'
  if (errorMsg.includes('城市') || errorMsg.includes('搜索')) return 'search'
  if (errorMsg.includes('服务器') || errorMsg.includes('api')) return 'server'
  return 'general'
})

// 热门城市
const popularCities = [
  '北京', '上海', '广州', '深圳', '杭州',
  '南京', '武汉', '成都', '西安', '重庆'
]

// 方法
const handleCitySearch = async (city: string) => {
  try {
    lastSearchType.value = 'city'
    lastSearchValue.value = city
    await weatherStore.fetchWeatherByCity(city)
  } catch (error) {
    console.error('城市搜索失败:', error)
  }
}

const handleLocationSearch = async () => {
  try {
    lastSearchType.value = 'location'
    lastSearchValue.value = null
    await weatherStore.fetchWeatherByCurrentLocation()
  } catch (error) {
    console.error('位置搜索失败:', error)
  }
}

const refreshWeather = async () => {
  if (!hasWeatherData.value) return

  try {
    await weatherStore.refreshWeather()
  } catch (error) {
    console.error('刷新失败:', error)
  }
}

const handleRetry = async () => {
  if (lastSearchType.value === 'city' && lastSearchValue.value) {
    await handleCitySearch(lastSearchValue.value)
  } else if (lastSearchType.value === 'location') {
    await handleLocationSearch()
  }
}

const clearError = () => {
  weatherStore.clearError()
}

const removeFromHistory = (city: string) => {
  weatherStore.removeFromSearchHistory(city)
}

const searchPopularCity = () => {
  const randomCity = popularCities[Math.floor(Math.random() * popularCities.length)]
  handleCitySearch(randomCity)
}

// 生命周期
onMounted(async () => {
  weatherStore.initialize()

  if (preferences.value.autoLocation) {
    try {
      await handleLocationSearch()
    } catch (error) {
      console.log('自动定位失败，忽略错误')
    }
  }
})
</script>
```

**主页面设计要点**：
- **布局结构**：头部、主内容、侧边栏、页脚
- **状态管理**：集成 Pinia store
- **响应式设计**：Grid 布局适配不同屏幕
- **交互逻辑**：完整的用户交互流程

## 样式设计和响应式

### UnoCSS 样式系统

我们使用 UnoCSS 构建了完整的样式系统：

#### 1. 颜色系统
```typescript
// uno.config.ts
theme: {
  colors: {
    primary: {
      50: '#eff6ff',
      100: '#dbeafe',
      500: '#3b82f6',
      600: '#2563eb',
      // ... 完整的颜色阶梯
    },
  },
}
```

#### 2. 组件样式
```typescript
shortcuts: {
  'btn': 'px-4 py-2 rounded-lg font-medium transition-all duration-200',
  'btn-primary': 'btn bg-primary-500 text-white hover:bg-primary-600',
  'btn-secondary': 'btn bg-gray-200 text-gray-800 hover:bg-gray-300',
  'card': 'bg-white rounded-xl shadow-lg border border-gray-100',
  'input': 'px-4 py-3 rounded-lg border border-gray-300 focus:border-primary-500',
}
```

#### 3. 响应式断点
- `sm`: 640px - 小屏幕
- `md`: 768px - 中等屏幕
- `lg`: 1024px - 大屏幕
- `xl`: 1280px - 超大屏幕

### 响应式布局策略

#### 移动端优先
```vue
<!-- 基础样式为移动端，然后向上扩展 -->
<div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
  <!-- 在大屏幕上变为 3 列，小屏幕保持 1 列 -->
</div>
```

#### 弹性布局
```vue
<!-- 按钮组自适应 -->
<div class="flex flex-col sm:flex-row gap-3 justify-center">
  <!-- 小屏幕垂直排列，大屏幕水平排列 -->
</div>
```

#### 条件显示
```vue
<!-- 根据屏幕大小显示不同内容 -->
<div class="hidden lg:block">
  <!-- 仅在大屏幕显示 -->
</div>
```

## 测试和质量保证

### 单元测试实现

创建组件测试 `src/components/__tests__/WeatherCard.spec.ts`：

```typescript
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import WeatherCard from '../WeatherCard.vue'
import type { WeatherResponse } from '@/types/weather'

// 模拟数据
const mockWeatherData: WeatherResponse = {
  location: {
    name: '北京',
    country: 'CN',
    latitude: 39.9042,
    longitude: 116.4074,
    timezone: 28800,
  },
  current: {
    temperature: 25.5,
    feels_like: 27.0,
    temp_min: 20.0,
    temp_max: 30.0,
    pressure: 1013,
    humidity: 60,
    visibility: 10000,
    uv_index: 5,
    weather: [
      {
        id: 800,
        main: 'Clear',
        description: '晴天',
        icon: '01d',
      },
    ],
    wind: {
      speed: 3.5,
      direction: 180,
      gust: 5.0,
    },
    clouds: { all: 0 },
    sunrise: 1640995200,
    sunset: 1641031200,
    updated_at: '2024-01-01T12:00:00Z',
  },
  timestamp: 1640995200,
  provider: 'openweathermap',
}

describe('WeatherCard', () => {
  it('renders weather data correctly', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查城市名称
    expect(wrapper.text()).toContain('北京')
    expect(wrapper.text()).toContain('CN')

    // 检查温度
    expect(wrapper.text()).toContain('26°C')

    // 检查天气描述
    expect(wrapper.text()).toContain('晴天')

    // 检查湿度
    expect(wrapper.text()).toContain('60%')
  })

  it('displays temperature in different units', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'imperial',
      },
    })

    expect(wrapper.text()).toContain('°F')
  })

  it('applies correct background gradient', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    const cardElement = wrapper.find('.card')
    expect(cardElement.classes()).toContain('bg-gradient-to-br')
  })
})
```

### 工具函数测试

创建 `src/utils/__tests__/weather.spec.ts`：

```typescript
import { describe, it, expect } from 'vitest'
import {
  formatTemperature,
  formatWindSpeed,
  formatWindDirection,
  validateCityName,
  validateCoordinates,
} from '../weather'

describe('Weather Utils', () => {
  describe('formatTemperature', () => {
    it('formats metric temperature correctly', () => {
      expect(formatTemperature(25.7, 'metric')).toBe('26°C')
      expect(formatTemperature(-5.2, 'metric')).toBe('-5°C')
    })

    it('formats imperial temperature correctly', () => {
      expect(formatTemperature(77.5, 'imperial')).toBe('78°F')
    })

    it('defaults to metric when no unit specified', () => {
      expect(formatTemperature(20)).toBe('20°C')
    })
  })

  describe('validateCityName', () => {
    it('validates correct city names', () => {
      expect(validateCityName('北京')).toBe(true)
      expect(validateCityName('Beijing')).toBe(true)
      expect(validateCityName('New York')).toBe(true)
    })

    it('rejects invalid city names', () => {
      expect(validateCityName('')).toBe(false)
      expect(validateCityName('   ')).toBe(false)
      expect(validateCityName('a'.repeat(101))).toBe(false)
    })
  })

  describe('validateCoordinates', () => {
    it('validates correct coordinates', () => {
      expect(validateCoordinates(39.9042, 116.4074)).toBe(true)
      expect(validateCoordinates(0, 0)).toBe(true)
      expect(validateCoordinates(-90, -180)).toBe(true)
      expect(validateCoordinates(90, 180)).toBe(true)
    })

    it('rejects invalid coordinates', () => {
      expect(validateCoordinates(91, 0)).toBe(false)
      expect(validateCoordinates(-91, 0)).toBe(false)
      expect(validateCoordinates(0, 181)).toBe(false)
      expect(validateCoordinates(0, -181)).toBe(false)
    })
  })
})
```

### 代码质量工具

#### ESLint 配置
```json
{
  "extends": [
    "@vue/eslint-config-typescript",
    "@vue/eslint-config-prettier"
  ],
  "rules": {
    "@typescript-eslint/no-unused-vars": "error",
    "vue/multi-word-component-names": "off"
  }
}
```

#### TypeScript 配置
```json
{
  "compilerOptions": {
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "exactOptionalPropertyTypes": true
  }
}
```

**测试策略要点**：
- **组件测试**：验证组件渲染和交互
- **工具函数测试**：确保工具函数正确性
- **类型检查**：TypeScript 静态类型检查
- **代码规范**：ESLint + Prettier 代码质量

## 构建和部署

### 开发环境构建

```bash
# 开发模式
npm run dev

# 类型检查
npm run type-check

# 代码检查
npm run lint

# 运行测试
npm run test:unit
```

### 生产环境构建

```bash
# 构建生产版本
npm run build

# 预览构建结果
npm run preview
```

构建优化配置：

```typescript
// vite.config.ts
export default defineConfig({
  build: {
    target: 'es2015',
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: false,
    minify: 'terser',
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['vue', 'vue-router', 'pinia'],
          utils: ['axios', '@vueuse/core'],
        },
      },
    },
  },
})
```

### Docker 部署

#### Dockerfile 配置

```dockerfile
# 多阶段构建
FROM node:18-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

# 生产阶段
FROM nginx:alpine

# 复制构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制 Nginx 配置
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
```

#### Nginx 配置

```nginx
server {
    listen 3000;
    server_name localhost;
    root /usr/share/nginx/html;
    index index.html;

    # 静态资源缓存
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }

    # API 代理
    location /api/ {
        proxy_pass http://gin-weather-backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # SPA 路由支持
    location / {
        try_files $uri $uri/ /index.html;
    }

    # 健康检查
    location /health {
        return 200 "healthy\n";
        add_header Content-Type text/plain;
    }
}
```

### Docker Compose 集成

```yaml
version: '3.8'

services:
  # 前端服务
  gin-weather-frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    depends_on:
      - gin-weather-backend
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  # 后端服务
  gin-weather-backend:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - WEATHER_API_KEY=${WEATHER_API_KEY}
    restart: unless-stopped
```

## 性能优化和最佳实践

### 性能优化策略

#### 1. 代码分割
```typescript
// 路由级别的懒加载
const routes = [
  {
    path: '/',
    component: () => import('@/views/HomeView.vue')
  },
  {
    path: '/settings',
    component: () => import('@/views/SettingsView.vue')
  }
]
```

#### 2. 组件懒加载
```vue
<script setup lang="ts">
// 异步组件
const SettingsPanel = defineAsyncComponent(() => import('@/components/SettingsPanel.vue'))
</script>
```

#### 3. 图片优化
```vue
<template>
  <!-- 懒加载图片 -->
  <img
    v-lazy="imageUrl"
    :alt="description"
    loading="lazy"
  />
</template>
```

#### 4. API 缓存
```typescript
// 在 Pinia store 中实现缓存
const cache = new Map()

const fetchWeatherWithCache = async (key: string, fetcher: () => Promise<any>) => {
  if (cache.has(key)) {
    const cached = cache.get(key)
    if (Date.now() - cached.timestamp < 300000) { // 5分钟缓存
      return cached.data
    }
  }

  const data = await fetcher()
  cache.set(key, { data, timestamp: Date.now() })
  return data
}
```

### 最佳实践总结

#### 1. 组件设计原则
- **单一职责**：每个组件只负责一个功能
- **可复用性**：通过 props 和 slots 提高复用性
- **可测试性**：组件逻辑清晰，易于测试
- **性能优化**：合理使用 `v-memo` 和 `shallowRef`

#### 2. 状态管理原则
- **状态分离**：UI 状态、业务状态、用户偏好分离
- **最小化状态**：只存储必要的状态
- **不可变性**：避免直接修改状态对象
- **持久化**：重要状态持久化到本地存储

#### 3. 类型安全原则
- **完整类型定义**：为所有数据定义类型
- **严格模式**：启用 TypeScript 严格模式
- **类型推导**：充分利用 TypeScript 类型推导
- **运行时验证**：关键数据进行运行时验证

#### 4. 用户体验原则
- **加载状态**：提供清晰的加载反馈
- **错误处理**：友好的错误提示和恢复机制
- **响应式设计**：适配各种设备和屏幕
- **无障碍访问**：支持键盘导航和屏幕阅读器

#### 5. 代码质量原则
- **代码规范**：统一的代码风格和命名规范
- **测试覆盖**：重要功能的测试覆盖
- **文档完善**：清晰的代码注释和文档
- **持续集成**：自动化的构建和测试流程

## 总结

通过这个详细的前端实现教程，我们完成了一个现代化的天气查询应用前端：

### 🎯 实现的功能
- ✅ 现代化的 Vue.js 3 + TypeScript 应用
- ✅ 响应式设计，支持多设备
- ✅ 智能搜索和地理位置定位
- ✅ 完整的状态管理和数据持久化
- ✅ 优雅的加载状态和错误处理
- ✅ 个性化设置和用户偏好
- ✅ 完整的测试覆盖和代码质量保证
- ✅ Docker 容器化部署

### 🛠️ 使用的技术
- **前端框架**：Vue.js 3 (Composition API)
- **类型系统**：TypeScript
- **构建工具**：Vite
- **状态管理**：Pinia
- **样式方案**：UnoCSS
- **图标库**：Lucide Vue Next
- **HTTP 客户端**：Axios
- **测试框架**：Vitest
- **代码质量**：ESLint + Prettier

### 📚 学到的知识
- Vue.js 3 Composition API 的使用
- TypeScript 在 Vue 项目中的应用
- 现代化的前端工程化实践
- 响应式设计和移动端适配
- 状态管理和数据持久化
- 组件化开发和可复用性设计
- 测试驱动开发和代码质量保证
- Docker 容器化部署

这个教程展示了如何从零开始构建一个生产级别的现代化前端应用，涵盖了前端开发的各个方面，是学习现代前端开发的完整指南。