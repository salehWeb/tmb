import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import Home from '../pages/Home.vue'
import Chat from '../pages/Chat.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/login', name: 'Login', component: Login },
    { path: '/chat', name: 'Chat', component: Chat }
]

const router = createRouter({ history: createWebHistory(), routes })
export default router
