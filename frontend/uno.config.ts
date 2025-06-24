import { defineConfig, presetUno, presetAttributify, presetIcons } from 'unocss'

export default defineConfig({
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons({
      collections: {
        lucide: () => import('@iconify-json/lucide/icons.json').then(i => i.default),
      },
    }),
  ],
  theme: {
    colors: {
      primary: {
        50: '#eff6ff',
        100: '#dbeafe',
        200: '#bfdbfe',
        300: '#93c5fd',
        400: '#60a5fa',
        500: '#3b82f6',
        600: '#2563eb',
        700: '#1d4ed8',
        800: '#1e40af',
        900: '#1e3a8a',
      },
      gray: {
        50: '#f9fafb',
        100: '#f3f4f6',
        200: '#e5e7eb',
        300: '#d1d5db',
        400: '#9ca3af',
        500: '#6b7280',
        600: '#4b5563',
        700: '#374151',
        800: '#1f2937',
        900: '#111827',
      },
    },
  },
  shortcuts: {
    'btn': 'px-4 py-2 rounded-lg font-medium transition-all duration-200',
    'btn-primary': 'btn bg-primary-500 text-white hover:bg-primary-600 active:bg-primary-700',
    'btn-secondary': 'btn bg-gray-200 text-gray-800 hover:bg-gray-300 active:bg-gray-400',
    'card': 'bg-white rounded-xl shadow-lg border border-gray-100',
    'input': 'px-4 py-3 rounded-lg border border-gray-300 focus:border-primary-500 focus:ring-2 focus:ring-primary-200 outline-none transition-all duration-200',
  },
})
