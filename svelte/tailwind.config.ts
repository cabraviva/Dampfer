import type { Config } from "tailwindcss";
import flowbitePlugin from 'flowbite/plugin'

export default {
  content: ["./src/**/*.{html,js,svelte,ts}", , './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
  darkMode: 'selector',

  theme: {
    extend: {
      colors: {
        // flowbite-svelte
        primary: {
          50:  '#C7F7FF',
          100: '#B0EFF8',
          200: '#9AE7F1',
          300: '#84DFEB',
          400: '#6ED7E4',
          500: '#58D0DD',
          600: '#42C8D7',
          700: '#00D0E9',
          800: '#16B8C9',
          900: '#00B1C3'
        },
      }
    },
  },
  plugins: [flowbitePlugin]
} as Config
// #00D0E9 is main