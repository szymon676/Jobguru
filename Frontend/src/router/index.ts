import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../pages/HelloWorld/Hello.vue'
import NotFound from '@/pages/Error404Page.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HelloWorld
    },
    {
      path: '/:catchAll(.*)',
      name: 'NotFound',
      component: NotFound
    }
  ]
})

export default router
