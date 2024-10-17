import path from "@necto/path";
import { defineConfig } from "vite";
import pug from "@necto/vite/plugins/pug";
import glsl from "@necto/vite/plugins/glsl";
import wasm from "@necto/vite/plugins/wasm";
import vue from "@necto/vite/plugins/vue"

export default defineConfig({
  plugins: [
    vue(),
    pug(),
    wasm(),
    glsl()
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});