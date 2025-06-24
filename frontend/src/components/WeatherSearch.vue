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
          :class="{ 'border-red-300 focus:border-red-500 focus:ring-red-200': hasError }"
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
          class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-primary-500 disabled:opacity-50"
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
        class="absolute top-full left-0 right-0 mt-1 bg-white rounded-lg shadow-lg border border-gray-200 z-50 max-h-60 overflow-y-auto"
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
              class="opacity-0 group-hover:opacity-100 text-gray-400 hover:text-red-500 transition-opacity"
            >
              <X class="w-3 h-3" />
            </button>
          </button>
          
          <!-- 清除历史按钮 -->
          <button
            v-if="searchHistory.length > 0"
            @click="clearHistory"
            class="w-full text-left px-3 py-2 text-sm text-gray-500 hover:text-red-500 border-t border-gray-100 mt-1"
          >
            清除搜索历史
          </button>
        </div>
        
        <!-- 当前搜索项 -->
        <div v-if="searchQuery && !filteredHistory.includes(searchQuery)" class="p-2 border-t border-gray-100">
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
    return searchHistory.value.slice(0, 5) // 显示最近5个
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

const clearHistory = () => {
  weatherStore.clearSearchHistory()
  showSuggestions.value = false
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
