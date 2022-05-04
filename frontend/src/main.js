import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import EmojiPicker from 'vue3-emoji-picker'
import { loadFonts } from './plugins/webfontloader'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import "./assets/styles/style.css"
import "./assets/styles/emoji.css"
import 'remixicon/fonts/remixicon.css'
import 'line-awesome/dist/line-awesome/css/line-awesome.css'

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .use(EmojiPicker)
  .mount('#app')
