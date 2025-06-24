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
  if (errorMsg.includes('网络') || errorMsg.includes('连接')) {
    return 'network'
  }
  if (errorMsg.includes('位置') || errorMsg.includes('定位')) {
    return 'location'
  }
  if (errorMsg.includes('城市') || errorMsg.includes('搜索')) {
    return 'search'
  }
  if (errorMsg.includes('服务器') || errorMsg.includes('api')) {
    return 'server'
  }
  return 'general'
})

// 热门城市
const popularCities = [
  '北京',
  '上海',
  '广州',
  '深圳',
  '杭州',
  '南京',
  '武汉',
  '成都',
  '西安',
  '重庆',
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
  // 初始化 store
  weatherStore.initialize()

  // 如果启用了自动定位，尝试获取当前位置的天气
  if (preferences.value.autoLocation) {
    try {
      await handleLocationSearch()
    } catch (error) {
      console.log('自动定位失败，忽略错误')
    }
  }
})
</script>

<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
    <!-- 头部 -->
    <header class="bg-white/80 backdrop-blur-sm border-b border-white/20 sticky top-0 z-40">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <div
              class="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center"
            >
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
          <WeatherSearch @search="handleCitySearch" @location-search="handleLocationSearch" />
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
              <div
                class="w-24 h-24 bg-gradient-to-br from-blue-400 to-purple-600 rounded-full mx-auto mb-6 flex items-center justify-center"
              >
                <Cloud class="w-12 h-12 text-white" />
              </div>
              <h2 class="text-2xl font-bold text-gray-900 mb-4">欢迎使用天气查询</h2>
              <p class="text-gray-600 mb-6">输入城市名称或使用当前位置来查询实时天气信息</p>
              <div class="flex flex-col sm:flex-row gap-3 justify-center">
                <button @click="handleLocationSearch" class="btn-primary flex items-center gap-2">
                  <MapPin class="w-4 h-4" />
                  使用当前位置
                </button>
                <button @click="searchPopularCity" class="btn-secondary flex items-center gap-2">
                  <Search class="w-4 h-4" />
                  查看热门城市
                </button>
              </div>
            </div>
          </div>

          <!-- 侧边栏 -->
          <div class="space-y-6">
            <!-- 设置面板 -->
            <SettingsPanel v-if="showSettings" @close="showSettings = false" />

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
          <p>
            数据来源: OpenWeatherMap |
            <a
              href="https://github.com/your-username/gin-weather"
              class="text-primary-500 hover:text-primary-600"
            >
              GitHub
            </a>
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>
