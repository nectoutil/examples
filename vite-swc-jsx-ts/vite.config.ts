import path from "@necto/path";
import { defineConfig } from 'vite';
import react from "@necto/vite/plugins/react"

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
