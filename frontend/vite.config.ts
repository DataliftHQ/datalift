import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: "build",
  },
  server: {
    host: "localhost",
    port: 8888,
    proxy: {
      "^/api/.*": {
        target: "http://0.0.0.0:8080",
        changeOrigin: true,
        xfwd: true,
      },
      "^/auth/.*": {
        target: "http://0.0.0.0:8080",
        changeOrigin: true,
        xfwd: true,
      },
    },
  },
  plugins: [react(), tsconfigPaths()],
});
