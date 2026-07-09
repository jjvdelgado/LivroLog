// router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import InitialPage from '@/views/InitialPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import HomePage from '@/views/HomePage.vue'
import ProfilePage from '@/views/ProfilePage.vue'
import ReviewPage from '@/views/ReviewPage.vue'
import AddNewBook from '@/views/AddNewBook.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: InitialPage
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterPage
  },
  {
    path: '/homepage',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/profile',
    name: 'Profile',
    component: ProfilePage
  },
  {
    path: '/review',
    name: 'Review',
    component: ReviewPage
  },
  {
    path: '/addnewbook',
    name: 'AddNewBook',
    component: AddNewBook
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router