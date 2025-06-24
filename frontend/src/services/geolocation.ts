import type { GeolocationPosition, GeolocationError } from '@/types/weather'

export class GeolocationService {
  /**
   * 获取当前位置
   */
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

  /**
   * 监听位置变化
   */
  static watchPosition(
    onSuccess: (position: GeolocationPosition) => void,
    onError: (error: GeolocationError) => void
  ): number | null {
    if (!navigator.geolocation) {
      onError({
        code: 0,
        message: '浏览器不支持地理位置功能',
      })
      return null
    }

    const options: PositionOptions = {
      enableHighAccuracy: true,
      timeout: 10000,
      maximumAge: 60000, // 1分钟缓存
    }

    return navigator.geolocation.watchPosition(
      (position) => {
        onSuccess({
          coords: {
            latitude: position.coords.latitude,
            longitude: position.coords.longitude,
            accuracy: position.coords.accuracy,
          },
        })
      },
      (error) => {
        onError(this.handleGeolocationError(error))
      },
      options
    )
  }

  /**
   * 停止监听位置变化
   */
  static clearWatch(watchId: number): void {
    if (navigator.geolocation) {
      navigator.geolocation.clearWatch(watchId)
    }
  }

  /**
   * 处理地理位置错误
   */
  private static handleGeolocationError(error: GeolocationPositionError): GeolocationError {
    switch (error.code) {
      case error.PERMISSION_DENIED:
        return {
          code: error.code,
          message: '用户拒绝了地理位置请求',
        }
      case error.POSITION_UNAVAILABLE:
        return {
          code: error.code,
          message: '位置信息不可用',
        }
      case error.TIMEOUT:
        return {
          code: error.code,
          message: '获取位置信息超时',
        }
      default:
        return {
          code: error.code,
          message: '获取位置信息时发生未知错误',
        }
    }
  }

  /**
   * 检查地理位置权限
   */
  static async checkPermission(): Promise<PermissionState> {
    if (!navigator.permissions) {
      throw new Error('浏览器不支持权限 API')
    }

    try {
      const permission = await navigator.permissions.query({ name: 'geolocation' })
      return permission.state
    } catch (error) {
      throw new Error('检查地理位置权限失败')
    }
  }

  /**
   * 格式化坐标显示
   */
  static formatCoordinates(lat: number, lon: number, precision: number = 4): string {
    const latDir = lat >= 0 ? 'N' : 'S'
    const lonDir = lon >= 0 ? 'E' : 'W'
    
    return `${Math.abs(lat).toFixed(precision)}°${latDir}, ${Math.abs(lon).toFixed(precision)}°${lonDir}`
  }

  /**
   * 计算两点间距离（公里）
   */
  static calculateDistance(lat1: number, lon1: number, lat2: number, lon2: number): number {
    const R = 6371 // 地球半径（公里）
    const dLat = this.toRadians(lat2 - lat1)
    const dLon = this.toRadians(lon2 - lon1)
    
    const a = 
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(this.toRadians(lat1)) * Math.cos(this.toRadians(lat2)) *
      Math.sin(dLon / 2) * Math.sin(dLon / 2)
    
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
    return R * c
  }

  /**
   * 角度转弧度
   */
  private static toRadians(degrees: number): number {
    return degrees * (Math.PI / 180)
  }
}

export default GeolocationService
