import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/', // Главная страница с созданием комнаты
    name: 'RoomCreate',
    component: null,
  },
  {
    path: '/room/:rid/:uid',
    name: 'Room',
    component: null,
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
