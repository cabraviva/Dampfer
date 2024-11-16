import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

export default {
  preprocess: vitePreprocess({
    style: {
      scss: {
        prependData: '' // Optional: add global styles or variables here
      }
    }
  }),
}
