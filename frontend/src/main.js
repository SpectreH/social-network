import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueNativeSock from "vue-native-websocket-vue3";
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import "./assets/styles/style.css"
import 'remixicon/fonts/remixicon.css'
import 'line-awesome/dist/line-awesome/css/line-awesome.css'
import "emoji-mart-vue-fast/css/emoji-mart.css";
import axios from 'axios'
import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-sugar.css';

export default function getTime(givenTime) {
  const cutTime = givenTime.substring(0, 16)
  const [date, time] = cutTime.split("T");
  const [year, month, day] = date.split("-");

  return `${day}.${month}.${year} ${time}`
}

require('@/store/auth/subscriber')
axios.defaults.baseURL = 'http://127.0.0.1:4000/'

const app = createApp(App)

store.dispatch('auth/authMe', localStorage.getItem('sn_token')).then(function() {
  app.use(store)
    .use(router)
    .use(VueNativeSock, "ws://127.0.0.1:4000/api/socket", { store: store, "connectManually": true }).use(VueToast)
    .mount("#app")
}.bind(app))


export { app }