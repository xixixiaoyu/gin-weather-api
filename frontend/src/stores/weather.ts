import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { WeatherResponse, LoadingState, GeolocationPosition } from '@/types/weather'
import WeatherApiService from '@/services/weatherApi'
import GeolocationService from '@/services/geolocation'

export const useWeatherStore = defineStore('weather', () => {
  // 状态
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
  
  // 当前位置
  const currentLocation = ref<GeolocationPosition | null>(null)
  
  // 计算属性
  const isLoading = computed(() => loadingState.value.isLoading)
  const error = computed(() => loadingState.value.error)
  const hasWeatherData = computed(() => currentWeather.value !== null)
  
  // 设置加载状态
  const setLoading = (loading: boolean) => {
    loadingState.value.isLoading = loading
    if (loading) {
      loadingState.value.error = null
    }
  }
  
  // 设置错误状态
  const setError = (error: string | null) => {
    loadingState.value.error = error
    loadingState.value.isLoading = false
  }
  
  // 清除错误
  const clearError = () => {
    loadingState.value.error = null
  }
  
  // 根据城市名称获取天气
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
  
  // 获取当前位置的天气
  const fetchWeatherByCurrentLocation = async () => {
    try {
      setLoading(true)
      
      // 获取当前位置
      const position = await GeolocationService.getCurrentPosition()
      currentLocation.value = position
      
      // 获取天气数据
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
  
  // 刷新当前天气数据
  const refreshWeather = async () => {
    if (!currentWeather.value) {
      return
    }
    
    const { location } = currentWeather.value
    await fetchWeatherByCoordinates(location.latitude, location.longitude)
  }
  
  // 添加到搜索历史
  const addToSearchHistory = (city: string) => {
    const trimmedCity = city.trim()
    if (!trimmedCity) return
    
    // 移除重复项
    const filtered = searchHistory.value.filter(item => 
      item.toLowerCase() !== trimmedCity.toLowerCase()
    )
    
    // 添加到开头
    searchHistory.value = [trimmedCity, ...filtered].slice(0, 10) // 最多保存10个
    
    // 保存到本地存储
    saveSearchHistory()
  }
  
  // 清除搜索历史
  const clearSearchHistory = () => {
    searchHistory.value = []
    localStorage.removeItem('weather-search-history')
  }
  
  // 移除搜索历史项
  const removeFromSearchHistory = (city: string) => {
    searchHistory.value = searchHistory.value.filter(item => 
      item.toLowerCase() !== city.toLowerCase()
    )
    saveSearchHistory()
  }
  
  // 保存搜索历史到本地存储
  const saveSearchHistory = () => {
    localStorage.setItem('weather-search-history', JSON.stringify(searchHistory.value))
  }
  
  // 从本地存储加载搜索历史
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
  
  // 更新用户偏好设置
  const updatePreferences = (newPreferences: Partial<typeof preferences.value>) => {
    preferences.value = { ...preferences.value, ...newPreferences }
    savePreferences()
  }
  
  // 保存用户偏好设置
  const savePreferences = () => {
    localStorage.setItem('weather-preferences', JSON.stringify(preferences.value))
  }
  
  // 加载用户偏好设置
  const loadPreferences = () => {
    try {
      const saved = localStorage.getItem('weather-preferences')
      if (saved) {
        preferences.value = { ...preferences.value, ...JSON.parse(saved) }
      }
    } catch (error) {
      console.error('加载用户偏好设置失败:', error)
    }
  }
  
  // 初始化
  const initialize = () => {
    loadPreferences()
    loadSearchHistory()
  }
  
  return {
    // 状态
    currentWeather,
    loadingState,
    preferences,
    searchHistory,
    currentLocation,
    
    // 计算属性
    isLoading,
    error,
    hasWeatherData,
    
    // 方法
    setLoading,
    setError,
    clearError,
    fetchWeatherByCity,
    fetchWeatherByCoordinates,
    fetchWeatherByCurrentLocation,
    refreshWeather,
    addToSearchHistory,
    clearSearchHistory,
    removeFromSearchHistory,
    updatePreferences,
    initialize,
  }
})
