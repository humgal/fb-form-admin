import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import FormList from '@/views/FormList.vue'
import FormHisList from '@/views/FormHisList.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: FormList
  },
  {
    path: '/formlist',
    name: 'formlist',
    component: FormList
  },
  {
    path: '/formhislist',
    name: 'formhislist',
    component: FormHisList
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
