<template>
  <div class="card p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">设置</h3>
      <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600">
        <X class="w-5 h-5" />
      </button>
    </div>

    <div class="space-y-6">
      <!-- 温度单位 -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-3">
          <Thermometer class="w-4 h-4 inline mr-2" />
          温度单位
        </label>
        <div class="grid grid-cols-3 gap-2">
          <button
            v-for="unit in temperatureUnits"
            :key="unit.value"
            @click="updateUnits(unit.value)"
            class="p-3 text-sm rounded-lg border transition-all"
            :class="
              preferences.units === unit.value
                ? 'border-primary-500 bg-primary-50 text-primary-700'
                : 'border-gray-200 hover:border-gray-300'
            "
          >
            <div class="font-medium">{{ unit.label }}</div>
            <div class="text-xs text-gray-500">{{ unit.description }}</div>
          </button>
        </div>
      </div>

      <!-- 语言设置 -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-3">
          <Globe class="w-4 h-4 inline mr-2" />
          语言
        </label>
        <select
          :value="preferences.language"
          @change="updateLanguage(($event.target as HTMLSelectElement).value)"
          class="input w-full"
        >
          <option v-for="lang in languages" :key="lang.value" :value="lang.value">
            {{ lang.label }}
          </option>
        </select>
      </div>

      <!-- 自动定位 -->
      <div>
        <label class="flex items-center justify-between">
          <div class="flex items-center">
            <MapPin class="w-4 h-4 mr-2 text-gray-600" />
            <div>
              <div class="text-sm font-medium text-gray-700">自动定位</div>
              <div class="text-xs text-gray-500">启动时自动获取当前位置的天气</div>
            </div>
          </div>
          <button
            @click="toggleAutoLocation"
            class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors"
            :class="preferences.autoLocation ? 'bg-primary-500' : 'bg-gray-200'"
          >
            <span
              class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
              :class="preferences.autoLocation ? 'translate-x-6' : 'translate-x-1'"
            />
          </button>
        </label>
      </div>

      <!-- 数据管理 -->
      <div class="border-t border-gray-200 pt-6">
        <h4 class="text-sm font-medium text-gray-700 mb-4">
          <Database class="w-4 h-4 inline mr-2" />
          数据管理
        </h4>

        <div class="space-y-3">
          <button
            @click="clearSearchHistory"
            class="w-full text-left p-3 rounded-lg border border-gray-200 hover:border-gray-300 transition-colors"
          >
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm font-medium text-gray-700">清除搜索历史</div>
                <div class="text-xs text-gray-500">
                  已保存 {{ searchHistory.length }} 条搜索记录
                </div>
              </div>
              <Trash2 class="w-4 h-4 text-gray-400" />
            </div>
          </button>

          <button
            @click="resetSettings"
            class="w-full text-left p-3 rounded-lg border border-gray-200 hover:border-red-200 hover:text-red-600 transition-colors"
          >
            <div class="flex items-center justify-between">
              <div>
                <div class="text-sm font-medium">重置所有设置</div>
                <div class="text-xs text-gray-500">恢复到默认设置</div>
              </div>
              <RotateCcw class="w-4 h-4 text-gray-400" />
            </div>
          </button>
        </div>
      </div>

      <!-- 关于信息 -->
      <div class="border-t border-gray-200 pt-6">
        <h4 class="text-sm font-medium text-gray-700 mb-4">
          <Info class="w-4 h-4 inline mr-2" />
          关于
        </h4>

        <div class="text-sm text-gray-600 space-y-2">
          <div>版本: 1.0.0</div>
          <div>数据来源: OpenWeatherMap</div>
          <div>
            <a
              href="https://github.com/your-username/gin-weather"
              target="_blank"
              class="text-primary-500 hover:text-primary-600"
            >
              GitHub 仓库
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { X, Thermometer, Globe, MapPin, Database, Trash2, RotateCcw, Info } from 'lucide-vue-next'
import { useWeatherStore } from '@/stores/weather'

interface Emits {
  (e: 'close'): void
}

defineEmits<Emits>()

const weatherStore = useWeatherStore()

// 计算属性
const preferences = computed(() => weatherStore.preferences)
const searchHistory = computed(() => weatherStore.searchHistory)

// 温度单位选项
const temperatureUnits = [
  {
    value: 'metric',
    label: '摄氏度',
    description: '°C',
  },
  {
    value: 'imperial',
    label: '华氏度',
    description: '°F',
  },
  {
    value: 'standard',
    label: '开尔文',
    description: 'K',
  },
] as const

// 语言选项
const languages = [
  { value: 'zh_cn', label: '简体中文' },
  { value: 'en', label: 'English' },
  { value: 'zh_tw', label: '繁體中文' },
  { value: 'ja', label: '日本語' },
  { value: 'ko', label: '한국어' },
]

// 方法
const updateUnits = (units: 'metric' | 'imperial' | 'standard') => {
  weatherStore.updatePreferences({ units })
}

const updateLanguage = (language: string) => {
  weatherStore.updatePreferences({ language })
}

const toggleAutoLocation = () => {
  weatherStore.updatePreferences({
    autoLocation: !preferences.value.autoLocation,
  })
}

const clearSearchHistory = () => {
  if (confirm('确定要清除所有搜索历史吗？')) {
    weatherStore.clearSearchHistory()
  }
}

const resetSettings = () => {
  if (confirm('确定要重置所有设置吗？这将清除所有数据并恢复默认设置。')) {
    weatherStore.updatePreferences({
      units: 'metric',
      language: 'zh_cn',
      autoLocation: false,
    })
    weatherStore.clearSearchHistory()
  }
}
</script>
