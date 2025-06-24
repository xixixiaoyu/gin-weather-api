<template>
  <div class="flex flex-col items-center justify-center p-8">
    <!-- 加载动画 -->
    <div class="relative">
      <!-- 外圈 -->
      <div class="w-16 h-16 border-4 border-primary-200 rounded-full animate-spin border-t-primary-500"></div>
      
      <!-- 内圈 -->
      <div class="absolute top-2 left-2 w-12 h-12 border-4 border-primary-100 rounded-full animate-spin animate-reverse border-t-primary-400"></div>
      
      <!-- 中心图标 -->
      <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
        <Cloud class="w-6 h-6 text-primary-500 animate-pulse" />
      </div>
    </div>
    
    <!-- 加载文本 -->
    <div class="mt-4 text-center">
      <p class="text-gray-600 font-medium">{{ message }}</p>
      <p v-if="subMessage" class="text-gray-400 text-sm mt-1">{{ subMessage }}</p>
    </div>
    
    <!-- 加载进度点 -->
    <div class="flex space-x-1 mt-4">
      <div 
        v-for="i in 3" 
        :key="i"
        class="w-2 h-2 bg-primary-400 rounded-full animate-bounce"
        :style="{ animationDelay: `${(i - 1) * 0.2}s` }"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Cloud } from 'lucide-vue-next'

interface Props {
  message?: string
  subMessage?: string
}

withDefaults(defineProps<Props>(), {
  message: '正在加载天气数据...',
  subMessage: '请稍候'
})
</script>

<style scoped>
@keyframes reverse {
  from {
    transform: rotate(360deg);
  }
  to {
    transform: rotate(0deg);
  }
}

.animate-reverse {
  animation: reverse 1s linear infinite;
}
</style>
