import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";

const env = loadEnv(
    'all',
    process.cwd()
);

const AUTH_SERVICE = env.VITE_AUTH_SERVICE_URL;

export default defineConfig({

  server: {
    proxy: {
      '/register': AUTH_SERVICE + '/register',
    },
  },
  plugins: [vue({
    template: {
      compilerOptions: {
        // Enable JSX
        isCustomElement: (tag) => tag.includes('-')
      }
    }
  })],
  esbuild: {
    jsxFactory: 'h',
    jsxFragment: 'Fragment'
  }
});
