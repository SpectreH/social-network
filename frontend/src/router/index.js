import { createRouter, createWebHistory } from 'vue-router'
import PostsView from "../views/PostsView.vue";

const routes = [
  {
    path: "/",
    name: "Posts",
    component: PostsView
  },
  {
    path: "/about",
    name: "About",
    component: () => import("../views/AboutView.vue")
  }
]; 

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
