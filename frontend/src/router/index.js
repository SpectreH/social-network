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
import GroupsFeedView from '../views/GroupsFeedView.vue'
import GroupView from '../views/GroupView.vue'
import LogoutView from '../views/LogoutView.vue'
import store from '@/store'
import auth from '@/middleware/auth'
import guest from '@/middleware/guest'
import middlewarePipeline from './middlewarePipeline'


const routes = [
  {
    path: '',
    name: 'home',
    component: NewsFeed,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/logout',
    name: 'logout',
    component: LogoutView,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/sign-up',
    name: 'sign-up',
    component: SignUp,
    meta: {
      middleware: [
        guest
      ]
    } 
  },
  {
    path: '/sign-in',
    name: 'sign-in',
    component: SignIn,
    meta: {
      middleware: [
        guest
      ]
    }
  },
  {
    path: '/user/:userId',
    name: 'user',
    component: ProfileView,
    props: true,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/post/:postId',
    name: 'post',
    component: PostView,
    props: true,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/profile-settings',
    name: 'profileSettings',
    component: EditProfielView,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/privacy-settings',
    name: 'privacySettings',
    component: EditPrivacyView,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/follow-request',
    name: 'followRequest',
    component: FollowRequestView,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/chat',
    name: 'chat',
    component: ChatView,
    meta: {
      middleware: [
        auth
      ]
    }
  },  
  {
    path: '/chat/:chatId',
    name: 'exactChat',
    component: ChatView,
    props: true,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/group/',
    name: 'groupsFeed',
    component: GroupsFeedView,
    meta: {
      middleware: [
        auth
      ]
    }
  },
  {
    path: '/group/:groupId',
    name: 'exactGroup',
    component: GroupView,
    props: true,
    meta: {
      middleware: [
        auth
      ]
    }
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (!to.meta.middleware) {
    return next()
  }
  const middleware = to.meta.middleware

  const context = {
    to,
    from,
    next,
    store
  }


  return middleware[0]({
    ...context,
    next: middlewarePipeline(context, middleware, 1)
  })
})

export default router
