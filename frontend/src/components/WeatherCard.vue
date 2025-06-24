<template>
  <div class="card p-6 min-h-96 relative overflow-hidden" :class="backgroundClass">
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
            <component :is="weatherIcon" class="w-16 h-16 text-white/90" />
          </div>
          <div class="text-white/90 text-sm font-medium">
            {{ getWeatherDescription(weather.current.weather[0]?.description || '') }}
          </div>
        </div>
      </div>

      <!-- 详细信息网格 -->
      <div class="grid grid-cols-2 gap-4 text-sm">
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

        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Droplets class="w-4 h-4 mr-2" />
            <span class="text-white/80">湿度</span>
          </div>
          <div class="font-medium">{{ formatHumidity(weather.current.humidity) }}</div>
        </div>

        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Wind class="w-4 h-4 mr-2" />
            <span class="text-white/80">风速</span>
          </div>
          <div class="font-medium">
            {{ formatWindSpeed(weather.current.wind.speed, units) }}
            {{ formatWindDirection(weather.current.wind.direction) }}
          </div>
        </div>

        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Gauge class="w-4 h-4 mr-2" />
            <span class="text-white/80">气压</span>
          </div>
          <div class="font-medium">{{ formatPressure(weather.current.pressure) }}</div>
        </div>

        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Eye class="w-4 h-4 mr-2" />
            <span class="text-white/80">能见度</span>
          </div>
          <div class="font-medium">{{ formatVisibility(weather.current.visibility) }}</div>
        </div>

        <div class="bg-white/10 rounded-lg p-3 backdrop-blur-sm">
          <div class="flex items-center mb-1">
            <Sun class="w-4 h-4 mr-2" />
            <span class="text-white/80">紫外线</span>
          </div>
          <div class="font-medium">
            {{ weather.current.uv_index || 'N/A' }}
            <span
              class="ml-1 text-xs"
              :class="getUVIndexDescription(weather.current.uv_index || 0).color"
            >
              {{ getUVIndexDescription(weather.current.uv_index || 0).text }}
            </span>
          </div>
        </div>
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
  Wind,
  Gauge,
  Eye,
  Sun,
  Sunrise,
  Sunset,
  Cloud,
  CloudRain,
  CloudSnow,
  CloudLightning,
  Moon,
} from 'lucide-vue-next'
import type { WeatherResponse } from '@/types/weather'
import {
  formatTemperature,
  formatHumidity,
  formatWindSpeed,
  formatWindDirection,
  formatPressure,
  formatVisibility,
  formatTime,
  getWeatherDescription,
  getWeatherBackground,
  getWeatherIcon,
  getUVIndexDescription,
} from '@/utils/weather'

interface Props {
  weather: WeatherResponse
  units?: 'metric' | 'imperial' | 'standard'
}

const props = withDefaults(defineProps<Props>(), {
  units: 'metric',
})

// 天气图标组件映射
const iconComponents = {
  sun: Sun,
  moon: Moon,
  cloud: Cloud,
  'cloud-rain': CloudRain,
  'cloud-snow': CloudSnow,
  'cloud-lightning': CloudLightning,
  clouds: Cloud,
  'cloud-sun': Cloud,
  'cloud-moon': Cloud,
  'cloud-sun-rain': CloudRain,
  'cloud-moon-rain': CloudRain,
  'cloud-fog': Cloud,
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
