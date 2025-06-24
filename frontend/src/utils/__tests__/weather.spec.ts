import { describe, it, expect } from 'vitest'
import {
  formatTemperature,
  formatWindSpeed,
  formatWindDirection,
  formatPressure,
  formatHumidity,
  formatVisibility,
  formatTime,
  getWeatherIcon,
  getWeatherBackground,
  validateCityName,
  validateCoordinates,
  getUVIndexDescription,
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

    it('formats standard temperature correctly', () => {
      expect(formatTemperature(298.15, 'standard')).toBe('298K')
    })

    it('defaults to metric when no unit specified', () => {
      expect(formatTemperature(20)).toBe('20°C')
    })
  })

  describe('formatWindSpeed', () => {
    it('formats metric wind speed correctly', () => {
      expect(formatWindSpeed(3.5, 'metric')).toBe('3.5 m/s')
    })

    it('formats imperial wind speed correctly', () => {
      expect(formatWindSpeed(7.8, 'imperial')).toBe('7.8 mph')
    })

    it('defaults to metric when no unit specified', () => {
      expect(formatWindSpeed(5.2)).toBe('5.2 m/s')
    })
  })

  describe('formatWindDirection', () => {
    it('converts degrees to cardinal directions correctly', () => {
      expect(formatWindDirection(0)).toBe('N')
      expect(formatWindDirection(90)).toBe('E')
      expect(formatWindDirection(180)).toBe('S')
      expect(formatWindDirection(270)).toBe('W')
      expect(formatWindDirection(45)).toBe('NE')
      expect(formatWindDirection(225)).toBe('SW')
    })

    it('handles edge cases', () => {
      expect(formatWindDirection(360)).toBe('N')
      expect(formatWindDirection(361)).toBe('N')
    })
  })

  describe('formatPressure', () => {
    it('formats pressure correctly', () => {
      expect(formatPressure(1013)).toBe('1013 hPa')
      expect(formatPressure(1020)).toBe('1020 hPa')
    })
  })

  describe('formatHumidity', () => {
    it('formats humidity correctly', () => {
      expect(formatHumidity(65)).toBe('65%')
      expect(formatHumidity(100)).toBe('100%')
    })
  })

  describe('formatVisibility', () => {
    it('formats visibility in kilometers for large values', () => {
      expect(formatVisibility(10000)).toBe('10.0 km')
      expect(formatVisibility(5500)).toBe('5.5 km')
    })

    it('formats visibility in meters for small values', () => {
      expect(formatVisibility(500)).toBe('500 m')
      expect(formatVisibility(100)).toBe('100 m')
    })
  })

  describe('formatTime', () => {
    it('formats timestamp to time string', () => {
      const timestamp = 1640995200 // 2022-01-01 12:00:00 UTC
      const result = formatTime(timestamp)
      
      // 结果应该是 HH:MM 格式
      expect(result).toMatch(/^\d{2}:\d{2}$/)
    })
  })

  describe('getWeatherIcon', () => {
    it('returns correct icon for known weather codes', () => {
      expect(getWeatherIcon('01d')).toBe('sun')
      expect(getWeatherIcon('01n')).toBe('moon')
      expect(getWeatherIcon('02d')).toBe('cloud-sun')
      expect(getWeatherIcon('09d')).toBe('cloud-rain')
    })

    it('returns default icon for unknown codes', () => {
      expect(getWeatherIcon('unknown')).toBe('cloud')
      expect(getWeatherIcon('')).toBe('cloud')
    })
  })

  describe('getWeatherBackground', () => {
    it('returns correct background for known weather types', () => {
      expect(getWeatherBackground('clear')).toBe('from-blue-400 to-blue-600')
      expect(getWeatherBackground('rain')).toBe('from-gray-500 to-blue-700')
      expect(getWeatherBackground('snow')).toBe('from-gray-300 to-blue-400')
    })

    it('returns default background for unknown weather types', () => {
      expect(getWeatherBackground('unknown')).toBe('from-blue-400 to-purple-600')
      expect(getWeatherBackground('')).toBe('from-blue-400 to-purple-600')
    })

    it('handles case insensitive input', () => {
      expect(getWeatherBackground('CLEAR')).toBe('from-blue-400 to-blue-600')
      expect(getWeatherBackground('Rain')).toBe('from-gray-500 to-blue-700')
    })
  })

  describe('validateCityName', () => {
    it('validates correct city names', () => {
      expect(validateCityName('北京')).toBe(true)
      expect(validateCityName('Beijing')).toBe(true)
      expect(validateCityName('New York')).toBe(true)
      expect(validateCityName('São Paulo')).toBe(true)
      expect(validateCityName('Los Angeles')).toBe(true)
    })

    it('rejects invalid city names', () => {
      expect(validateCityName('')).toBe(false)
      expect(validateCityName('   ')).toBe(false)
      expect(validateCityName('a'.repeat(101))).toBe(false) // 太长
    })

    it('handles special characters correctly', () => {
      expect(validateCityName('Saint-Denis')).toBe(true) // 连字符
      expect(validateCityName('St. Louis')).toBe(true) // 点号
      expect(validateCityName('City@Name')).toBe(false) // 特殊字符
    })
  })

  describe('validateCoordinates', () => {
    it('validates correct coordinates', () => {
      expect(validateCoordinates(39.9042, 116.4074)).toBe(true) // 北京
      expect(validateCoordinates(0, 0)).toBe(true) // 赤道和本初子午线交点
      expect(validateCoordinates(-90, -180)).toBe(true) // 边界值
      expect(validateCoordinates(90, 180)).toBe(true) // 边界值
    })

    it('rejects invalid coordinates', () => {
      expect(validateCoordinates(91, 0)).toBe(false) // 纬度超出范围
      expect(validateCoordinates(-91, 0)).toBe(false) // 纬度超出范围
      expect(validateCoordinates(0, 181)).toBe(false) // 经度超出范围
      expect(validateCoordinates(0, -181)).toBe(false) // 经度超出范围
    })
  })

  describe('getUVIndexDescription', () => {
    it('returns correct description for different UV index levels', () => {
      expect(getUVIndexDescription(1)).toEqual({ text: '低', color: 'text-green-500' })
      expect(getUVIndexDescription(3)).toEqual({ text: '中等', color: 'text-yellow-500' })
      expect(getUVIndexDescription(6)).toEqual({ text: '高', color: 'text-orange-500' })
      expect(getUVIndexDescription(8)).toEqual({ text: '很高', color: 'text-red-500' })
      expect(getUVIndexDescription(12)).toEqual({ text: '极高', color: 'text-purple-500' })
    })

    it('handles edge cases', () => {
      expect(getUVIndexDescription(0)).toEqual({ text: '低', color: 'text-green-500' })
      expect(getUVIndexDescription(2)).toEqual({ text: '低', color: 'text-green-500' })
      expect(getUVIndexDescription(5)).toEqual({ text: '中等', color: 'text-yellow-500' })
    })
  })
})
