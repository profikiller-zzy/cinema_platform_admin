import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/

export default ({mode}) => {

  // process.cwd() 获取当前工作目录的路径并将其作为 root 参数传递给 loadEnv 函数。
  // 这样做可以确保 Vite 能够正确地查找和加载环境变量文件。
  const env = loadEnv(mode, process.cwd())
  const baseUrl = env.VITE_API

  return defineConfig({
    // 以“VITE_”为前缀的代理才走这个配置
    envPrefix:["VITE_"],
    plugins: [vue(),],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
    },
    server: {
      proxy: {
        "/uploads" : {
          target: baseUrl
        },
        "/api": {
          target: baseUrl
        }
      }
    }
  })
}
