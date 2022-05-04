import { createRouter, createWebHistory } from 'vue-router'
import NewsFeed from '../views/NewsFeedView.vue'
import SignIn from '../views/SignInView.vue'
import SignUp from '../views/SignUpView.vue'
import PostView from '../views/PostView.vue'
import ProfileView from '../views/ProfileView.vue'
import EditProfielView from '../views/EditProfileView.vue'
import EditPrivacyView from '../views/EditPrivacyView.vue'
import FollowRequestView from '../views/FollowRequestView.vue'
import ChatView from '../views/ChatView.vue'

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
  },
  {
    path: '/privacy-settings',
    name: 'privacySettings',
    component: EditPrivacyView,
  },
  {
    path: '/follow-request',
    name: 'followRequest',
    component: FollowRequestView,
  },
  {
    path: '/chat',
    name: 'chat',
    component: ChatView,
  },  
  {
    path: '/chat/:chatId',
    name: 'exactChat',
    component: ChatView,
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
