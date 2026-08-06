import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// const env = loadEnv(
//     'all',
//     process.cwd()
// );

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/auth': {
        target: 'http://auth-service:8081',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/auth/, '')
      }
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
