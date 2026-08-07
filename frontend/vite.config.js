import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";


const AUTH_SERVICE_URL = process.env.VITE_AUTH_SERVICE_URL;
const POSTING_SERVICE_URL = process.env.VITE_POSTING_SERVICE_URL;
const API_BASE = process.env.VITE_BACKEND_DJANGO_URL;
export default defineConfig({
  
  plugins: [vue()],
  server: {
    proxy: {
      '/auth': {
        target: AUTH_SERVICE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/auth/, '')
      },
      '/posting': {
        target: POSTING_SERVICE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/posting/, '')
      },
      '/backend': {
        target: API_BASE,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/backend/, '')
      },
    }

  },
});
// const AUTH_SERVICE = env.VITE_AUTH_SERVICE_URL;

// export default defineConfig({

//   server: {
//     proxy: {
//       '/register': AUTH_SERVICE + '/register',
//     },
//   },
//   plugins: [vue({
//     template: {
//       compilerOptions: {
//         // Enable JSX
//         isCustomElement: (tag) => tag.includes('-')
//       }
//     }
//   })],
//   esbuild: {
//     jsxFactory: 'h',
//     jsxFragment: 'Fragment'
//   },
//   host: true,
//   port: 5173,
//   base: '/',
// });
