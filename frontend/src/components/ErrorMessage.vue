<template>
  <div class="card p-6 text-center">
    <!-- 错误图标 -->
    <div class="flex justify-center mb-4">
      <div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center">
        <component :is="errorIcon" class="w-8 h-8 text-red-500" />
      </div>
    </div>
    
    <!-- 错误标题 -->
    <h3 class="text-lg font-semibold text-gray-900 mb-2">
      {{ title }}
    </h3>
    
    <!-- 错误消息 -->
    <p class="text-gray-600 mb-6">
      {{ message }}
    </p>
    
    <!-- 操作按钮 -->
    <div class="flex flex-col sm:flex-row gap-3 justify-center">
      <button
        v-if="showRetry"
        @click="$emit('retry')"
        class="btn-primary flex items-center gap-2"
      >
        <RefreshCw class="w-4 h-4" />
        重试
      </button>
      
      <button
        v-if="showHome"
        @click="$emit('home')"
        class="btn-secondary flex items-center gap-2"
      >
        <Home class="w-4 h-4" />
        返回首页
      </button>
      
      <button
        v-if="showDismiss"
        @click="$emit('dismiss')"
        class="btn-secondary flex items-center gap-2"
      >
        <X class="w-4 h-4" />
        关闭
      </button>
    </div>
    
    <!-- 额外信息 -->
    <div v-if="details" class="mt-6 p-4 bg-gray-50 rounded-lg text-left">
      <details>
        <summary class="cursor-pointer text-sm text-gray-500 hover:text-gray-700">
          查看详细信息
        </summary>
        <pre class="mt-2 text-xs text-gray-600 whitespace-pre-wrap">{{ details }}</pre>
      </details>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { 
  AlertTriangle, 
  Wifi, 
  MapPin, 
  Search, 
  Server,
  RefreshCw, 
  Home, 
  X 
} from 'lucide-vue-next'

interface Props {
  type?: 'network' | 'location' | 'search' | 'server' | 'general'
  title?: string
  message: string
  details?: string
  showRetry?: boolean
  showHome?: boolean
  showDismiss?: boolean
}

interface Emits {
  (e: 'retry'): void
  (e: 'home'): void
  (e: 'dismiss'): void
}

const props = withDefaults(defineProps<Props>(), {
  type: 'general',
  showRetry: true,
  showHome: false,
  showDismiss: true
})

defineEmits<Emits>()

// 错误图标映射
const errorIcons = {
  network: Wifi,
  location: MapPin,
  search: Search,
  server: Server,
  general: AlertTriangle,
}

const errorIcon = computed(() => errorIcons[props.type])

// 默认标题映射
const defaultTitles = {
  network: '网络连接错误',
  location: '位置获取失败',
  search: '搜索失败',
  server: '服务器错误',
  general: '出现错误',
}

const title = computed(() => props.title || defaultTitles[props.type])
</script>
