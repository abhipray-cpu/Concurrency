import { createRouter, createWebHistory } from 'vue-router'
import HomePage from './views/home-page.vue'
import LoginPage from './views/login-page.vue'
import NotFound from './views/not-found.vue'
import ProfilePage from './views/profile-page.vue'
import SearchPage from './views/search-page.vue'
import SignupPage from './views/signup-page.vue'
import UpdatePage from './views/update-page.vue'
import WebPage from './views/web-page.vue'
const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    name: 'home',
    path: '/home',
    component: HomePage
  },
  {
    name: 'login',
    path: '/login',
    component: LoginPage
  },
  {
    name: 'profile',
    path: '/profile',
    component: ProfilePage
  },
  {
    name: 'search',
    path: '/search',
    component: SearchPage
  },
  {
    name: 'signup',
    path: '/signup',
    component: SignupPage
  },
  {
    name: 'update',
    path: '/update',
    component: UpdatePage
  },
  {
    name: 'page',
    path: '/page/:id',
    component: WebPage,
    props: true
  },
  {
    path: '/:catchAll(.*)',
    name: 'NotFound',
    component: NotFound
  }
] // Define your routes here

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
