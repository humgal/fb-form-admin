import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import FormList from '@/views/FormList.vue'
import FormHisList from '@/views/FormHisList.vue'
import FormDetail from '@/components/FormDetail.vue'
import FormHis from '@/components/FormHis.vue'
import UserLogin from '@/views/UserLogin.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: UserLogin
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
    path: '/formDetail',
    name: 'formDetail',
    component: FormDetail
  },
  {
    path: '/formHis',
    name: 'formHis',
    component: FormHis
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

router.beforeEach((to, from, next) => {
  if (to.path === '/login') {
    next()
  } else {
    const token = sessionStorage.getItem('token')
    if (token == null || token === '') {
      next('/login')
    } else {
      next()
    }
  }
})

export default router
