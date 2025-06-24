import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import WeatherCard from '../WeatherCard.vue'
import type { WeatherResponse } from '@/types/weather'

// 模拟天气数据
const mockWeatherData: WeatherResponse = {
  location: {
    name: '北京',
    country: 'CN',
    latitude: 39.9042,
    longitude: 116.4074,
    timezone: 28800,
  },
  current: {
    temperature: 25.5,
    feels_like: 27.0,
    temp_min: 20.0,
    temp_max: 30.0,
    pressure: 1013,
    humidity: 60,
    visibility: 10000,
    uv_index: 5,
    weather: [
      {
        id: 800,
        main: 'Clear',
        description: '晴天',
        icon: '01d',
      },
    ],
    wind: {
      speed: 3.5,
      direction: 180,
      gust: 5.0,
    },
    clouds: {
      all: 0,
    },
    sunrise: 1640995200,
    sunset: 1641031200,
    updated_at: '2024-01-01T12:00:00Z',
  },
  timestamp: 1640995200,
  provider: 'openweathermap',
}

describe('WeatherCard', () => {
  it('renders weather data correctly', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查城市名称
    expect(wrapper.text()).toContain('北京')
    expect(wrapper.text()).toContain('CN')

    // 检查温度
    expect(wrapper.text()).toContain('26°C') // 四舍五入后的温度

    // 检查天气描述
    expect(wrapper.text()).toContain('晴天')

    // 检查湿度
    expect(wrapper.text()).toContain('60%')
  })

  it('displays temperature in different units', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'imperial',
      },
    })

    // 应该显示华氏度
    expect(wrapper.text()).toContain('°F')
  })

  it('shows wind information', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查风速
    expect(wrapper.text()).toContain('3.5 m/s')
    
    // 检查风向
    expect(wrapper.text()).toContain('S') // 180度对应南风
  })

  it('displays pressure and visibility', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查气压
    expect(wrapper.text()).toContain('1013 hPa')

    // 检查能见度
    expect(wrapper.text()).toContain('10.0 km')
  })

  it('shows sunrise and sunset times', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查日出日落标签
    expect(wrapper.text()).toContain('日出')
    expect(wrapper.text()).toContain('日落')
  })

  it('applies correct background gradient based on weather', () => {
    const wrapper = mount(WeatherCard, {
      props: {
        weather: mockWeatherData,
        units: 'metric',
      },
    })

    // 检查是否应用了背景渐变类
    const cardElement = wrapper.find('.card')
    expect(cardElement.classes()).toContain('bg-gradient-to-br')
  })

  it('handles missing optional data gracefully', () => {
    const incompleteWeatherData = {
      ...mockWeatherData,
      current: {
        ...mockWeatherData.current,
        uv_index: 0, // 测试 UV 指数为 0 的情况
        rain: undefined,
        snow: undefined,
      },
    }

    const wrapper = mount(WeatherCard, {
      props: {
        weather: incompleteWeatherData,
        units: 'metric',
      },
    })

    // 应该正常渲染，不会崩溃
    expect(wrapper.text()).toContain('北京')
  })
})
