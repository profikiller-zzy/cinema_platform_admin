import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePresetEnv } from 'vite-preset-env';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    VitePresetEnv({
      include: ['jsx', 'typescript'], // 启用 jsx 和 typescript 支持
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
