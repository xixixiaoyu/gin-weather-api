// 天气数据类型定义

export interface WeatherRequest {
  city?: string
  lat?: number
  lon?: number
  units?: 'metric' | 'imperial' | 'standard'
  lang?: string
}

export interface Location {
  name: string
  country: string
  latitude: number
  longitude: number
  timezone: number
}

export interface Weather {
  id: number
  main: string
  description: string
  icon: string
}

export interface Wind {
  speed: number
  direction: number
  gust: number
}

export interface Clouds {
  all: number
}

export interface Rain {
  '1h'?: number
  '3h'?: number
}

export interface Snow {
  '1h'?: number
  '3h'?: number
}

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
  wind: Wind
  clouds: Clouds
  rain?: Rain
  snow?: Snow
  sunrise: number
  sunset: number
  updated_at: string
}

export interface WeatherResponse {
  location: Location
  current: Current
  timestamp: number
  provider: string
}

export interface APIResponse<T = any> {
  success: boolean
  data?: T
  error?: {
    error: string
    code: number
    message?: string
  }
}

// 地理位置相关类型
export interface GeolocationPosition {
  coords: {
    latitude: number
    longitude: number
    accuracy: number
  }
}

export interface GeolocationError {
  code: number
  message: string
}

// UI 状态类型
export interface LoadingState {
  isLoading: boolean
  error: string | null
}

// 天气图标映射
export const WEATHER_ICONS: Record<string, string> = {
  '01d': 'sun',
  '01n': 'moon',
  '02d': 'cloud-sun',
  '02n': 'cloud-moon',
  '03d': 'cloud',
  '03n': 'cloud',
  '04d': 'clouds',
  '04n': 'clouds',
  '09d': 'cloud-rain',
  '09n': 'cloud-rain',
  '10d': 'cloud-sun-rain',
  '10n': 'cloud-moon-rain',
  '11d': 'cloud-lightning',
  '11n': 'cloud-lightning',
  '13d': 'cloud-snow',
  '13n': 'cloud-snow',
  '50d': 'cloud-fog',
  '50n': 'cloud-fog',
}

// 天气背景颜色映射
export const WEATHER_BACKGROUNDS: Record<string, string> = {
  'clear': 'from-blue-400 to-blue-600',
  'clouds': 'from-gray-400 to-gray-600',
  'rain': 'from-gray-500 to-blue-700',
  'drizzle': 'from-gray-400 to-blue-500',
  'thunderstorm': 'from-gray-700 to-purple-900',
  'snow': 'from-gray-300 to-blue-400',
  'mist': 'from-gray-400 to-gray-500',
  'fog': 'from-gray-400 to-gray-500',
  'haze': 'from-yellow-300 to-orange-400',
  'default': 'from-blue-400 to-purple-600',
}
