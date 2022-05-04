import { createRouter, createWebHistory } from 'vue-router'
import NewsFeed from '../views/NewsFeedView.vue'
import SignIn from '../views/SignInView.vue'
import SignUp from '../views/SignUpView.vue'
import PostView from '../views/PostView.vue'
import ProfileView from '../views/ProfileView.vue'
import EditProfielView from '../views/EditProfileView.vue'
import EditPrivacySettings from '../views/EditPrivacyView.vue'

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
  {
    path: '/user/:userId',
    name: 'user',
    component: ProfileView,
    props: true,
  },
  {
    path: '/post/:postId',
    name: 'post',
    component: PostView,
    props: true,
  },
  {
    path: '/profile-settings',
    name: 'profileSettings',
    component: EditProfielView,
    props: true,
  },
  {
    path: '/privacy-settings',
    name: 'privacySettings',
    component: EditPrivacySettings,
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
