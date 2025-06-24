import { WEATHER_ICONS, WEATHER_BACKGROUNDS } from '@/types/weather'

/**
 * 获取天气图标名称
 */
export function getWeatherIcon(iconCode: string): string {
  return WEATHER_ICONS[iconCode] || 'cloud'
}

/**
 * 获取天气背景渐变类名
 */
export function getWeatherBackground(weatherMain: string): string {
  const main = weatherMain.toLowerCase()
  return WEATHER_BACKGROUNDS[main] || WEATHER_BACKGROUNDS.default
}

/**
 * 格式化温度显示
 */
export function formatTemperature(temp: number, units: string = 'metric'): string {
  const rounded = Math.round(temp)
  switch (units) {
    case 'imperial':
      return `${rounded}°F`
    case 'standard':
      return `${rounded}K`
    default:
      return `${rounded}°C`
  }
}

/**
 * 格式化风速显示
 */
export function formatWindSpeed(speed: number, units: string = 'metric'): string {
  switch (units) {
    case 'imperial':
      return `${speed.toFixed(1)} mph`
    default:
      return `${speed.toFixed(1)} m/s`
  }
}

/**
 * 格式化风向
 */
export function formatWindDirection(degrees: number): string {
  const directions = [
    'N',
    'NNE',
    'NE',
    'ENE',
    'E',
    'ESE',
    'SE',
    'SSE',
    'S',
    'SSW',
    'SW',
    'WSW',
    'W',
    'WNW',
    'NW',
    'NNW',
  ]

  const index = Math.round(degrees / 22.5) % 16
  return directions[index]
}

/**
 * 格式化压力显示
 */
export function formatPressure(pressure: number): string {
  return `${pressure} hPa`
}

/**
 * 格式化湿度显示
 */
export function formatHumidity(humidity: number): string {
  return `${humidity}%`
}

/**
 * 格式化能见度显示
 */
export function formatVisibility(visibility: number): string {
  if (visibility >= 1000) {
    return `${(visibility / 1000).toFixed(1)} km`
  }
  return `${visibility} m`
}

/**
 * 格式化时间显示
 */
export function formatTime(timestamp: number): string {
  return new Date(timestamp * 1000).toLocaleTimeString('zh-CN', {
    hour: '2-digit',
    minute: '2-digit',
  })
}

/**
 * 格式化日期时间显示
 */
export function formatDateTime(timestamp: number): string {
  return new Date(timestamp * 1000).toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

/**
 * 获取相对时间显示
 */
export function getRelativeTime(timestamp: number): string {
  const now = Date.now()
  const time = timestamp * 1000
  const diff = now - time

  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)

  if (minutes < 1) {
    return '刚刚'
  } else if (minutes < 60) {
    return `${minutes}分钟前`
  } else if (hours < 24) {
    return `${hours}小时前`
  } else {
    return `${days}天前`
  }
}

/**
 * 获取天气状况的中文描述
 */
export function getWeatherDescription(description: string): string {
  const descriptions: Record<string, string> = {
    'clear sky': '晴空',
    'few clouds': '少云',
    'scattered clouds': '多云',
    'broken clouds': '阴天',
    'overcast clouds': '阴天',
    'light rain': '小雨',
    'moderate rain': '中雨',
    'heavy rain': '大雨',
    'very heavy rain': '暴雨',
    'extreme rain': '大暴雨',
    'freezing rain': '冻雨',
    'light snow': '小雪',
    snow: '雪',
    'heavy snow': '大雪',
    sleet: '雨夹雪',
    mist: '薄雾',
    fog: '雾',
    haze: '霾',
    thunderstorm: '雷暴',
    drizzle: '毛毛雨',
  }

  return descriptions[description.toLowerCase()] || description
}

/**
 * 获取空气质量描述
 */
export function getAirQualityDescription(aqi: number): { text: string; color: string } {
  if (aqi <= 50) {
    return { text: '优', color: 'text-green-500' }
  } else if (aqi <= 100) {
    return { text: '良', color: 'text-yellow-500' }
  } else if (aqi <= 150) {
    return { text: '轻度污染', color: 'text-orange-500' }
  } else if (aqi <= 200) {
    return { text: '中度污染', color: 'text-red-500' }
  } else if (aqi <= 300) {
    return { text: '重度污染', color: 'text-purple-500' }
  } else {
    return { text: '严重污染', color: 'text-red-800' }
  }
}

/**
 * 获取紫外线指数描述
 */
export function getUVIndexDescription(uvIndex: number): { text: string; color: string } {
  if (uvIndex <= 2) {
    return { text: '低', color: 'text-green-500' }
  } else if (uvIndex <= 5) {
    return { text: '中等', color: 'text-yellow-500' }
  } else if (uvIndex <= 7) {
    return { text: '高', color: 'text-orange-500' }
  } else if (uvIndex <= 10) {
    return { text: '很高', color: 'text-red-500' }
  } else {
    return { text: '极高', color: 'text-purple-500' }
  }
}

/**
 * 验证城市名称
 */
export function validateCityName(city: string): boolean {
  if (!city || city.trim().length === 0) {
    return false
  }

  // 基本长度检查
  if (city.trim().length > 100) {
    return false
  }

  // 检查是否包含特殊字符（允许中文、英文、数字、空格、连字符、点号、重音符号等）
  const validPattern = /^[\u4e00-\u9fa5a-zA-Z0-9\s\-\.\u00C0-\u017F\u0100-\u024F]+$/
  return validPattern.test(city.trim())
}

/**
 * 验证坐标
 */
export function validateCoordinates(lat: number, lon: number): boolean {
  return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}
