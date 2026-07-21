import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
<<<<<<< HEAD
    plugins: [vue()],
    server: {
        host: true,
        port: 5173,
    },
=======
  plugins: [vue()],
  server: {
    host: true,
    port: 5173,
  },
>>>>>>> fix/SKEL2-20-implement-broadcast-feature
});
