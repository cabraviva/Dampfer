import { mount } from 'svelte'
import './app.scss'
import App from './App.svelte'
import { defineKnorryOptions } from 'knorry'
defineKnorryOptions({
    easyMode: false
})

const app = mount(App, {
  target: document.getElementById('app')!,
})

export default app
