import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NewUserView from '../views/NewUserView.vue'
import UserInfoView from '../views/UserInfoView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/new',
      name: 'new',
      component: NewUserView,
    },
    {
      path: '/user/:id',
      name: 'userInfo',
      component: UserInfoView,
      props: true,
    },
  ],
})

export default router
