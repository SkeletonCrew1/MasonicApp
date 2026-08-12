import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";


const AUTH_SERVICE_URL = process.env.VITE_AUTH_SERVICE_URL;
const POSTING_SERVICE_URL = process.env.VITE_POSTING_SERVICE_URL;
const DJANGO_API_BASE = process.env.VITE_BACKEND_DJANGO_URL;
const VOTING_API_BASE = process.env.VITE_VOTING_API_BASE;
export default defineConfig({
  
  plugins: [vue()],
  server: {
    allowedHosts: true,
    proxy: {
      '/auth': {
        target: AUTH_SERVICE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api\/auth/, '')
      },
      '/posting': {
        target: POSTING_SERVICE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api\/posting/, '')
      },
      '/backend': {
        target: DJANGO_API_BASE,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/backend/, '')
      },
      '/voting': {
        target: VOTING_API_BASE,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/voting/, '')
      },
    }

  },
});

