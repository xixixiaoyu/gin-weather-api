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
  /**
   * 健康检查
   */
  static async healthCheck(): Promise<APIResponse> {
    try {
      const response = await api.get('/health')
      return response.data
    } catch (error) {
      throw this.handleError(error)
    }
  }

  /**
   * 通用天气查询
   */
  static async getWeather(params: WeatherRequest): Promise<WeatherResponse> {
    try {
      const response = await api.get<APIResponse<WeatherResponse>>('/weather', { params })
      
      if (response.data.success && response.data.data) {
        return response.data.data
      } else {
        throw new Error(response.data.error?.message || '获取天气数据失败')
      }
    } catch (error) {
      throw this.handleError(error)
    }
  }

  /**
   * 根据城市名称查询天气
   */
  static async getWeatherByCity(
    city: string,
    units: 'metric' | 'imperial' | 'standard' = 'metric',
    lang: string = 'zh_cn'
  ): Promise<WeatherResponse> {
    try {
      const response = await api.get<APIResponse<WeatherResponse>>(`/weather/city/${encodeURIComponent(city)}`, {
        params: { units, lang }
      })
      
      if (response.data.success && response.data.data) {
        return response.data.data
      } else {
        throw new Error(response.data.error?.message || '获取天气数据失败')
      }
    } catch (error) {
      throw this.handleError(error)
    }
  }

  /**
   * 根据坐标查询天气
   */
  static async getWeatherByCoordinates(
    lat: number,
    lon: number,
    units: 'metric' | 'imperial' | 'standard' = 'metric',
    lang: string = 'zh_cn'
  ): Promise<WeatherResponse> {
    try {
      const response = await api.get<APIResponse<WeatherResponse>>(`/weather/coordinates/${lat}/${lon}`, {
        params: { units, lang }
      })
      
      if (response.data.success && response.data.data) {
        return response.data.data
      } else {
        throw new Error(response.data.error?.message || '获取天气数据失败')
      }
    } catch (error) {
      throw this.handleError(error)
    }
  }

  /**
   * 错误处理
   */
  private static handleError(error: any): Error {
    if (axios.isAxiosError(error)) {
      if (error.response) {
        // 服务器响应错误
        const errorData = error.response.data
        if (errorData?.error?.message) {
          return new Error(errorData.error.message)
        }
        return new Error(`请求失败: ${error.response.status} ${error.response.statusText}`)
      } else if (error.request) {
        // 网络错误
        return new Error('网络连接失败，请检查网络设置')
      }
    }
    
    // 其他错误
    return new Error(error.message || '未知错误')
  }
}

export default WeatherApiService
