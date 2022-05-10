import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { loadFonts } from './plugins/webfontloader'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import "./assets/styles/style.css"
import 'remixicon/fonts/remixicon.css'
import 'line-awesome/dist/line-awesome/css/line-awesome.css'
import "emoji-mart-vue-fast/css/emoji-mart.css";
import axios from 'axios'

axios.defaults.baseURL = 'http://127.0.0.1:4000/'

loadFonts()

createApp(App)
  .use(router)
  .use(store)
  .mount('#app')
