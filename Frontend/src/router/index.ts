import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../pages/HelloWorld/Hello.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HelloWorld
    },
  ]
})

export default router
