import { createRouter, createWebHistory } from 'vue-router'
import NewsFeed from '../views/NewsFeedView.vue'
import SignIn from '../views/SignInView.vue'
import SignUp from '../views/SignUpView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: NewsFeed
  },
  {
    path: '/sign-up',
    name: 'sign-up',
    component: SignUp
  },
  {
    path: '/sign-in',
    name: 'sign-in',
    component: SignIn
  },  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
