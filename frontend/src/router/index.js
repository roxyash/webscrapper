import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from './components/LoginForm.vue'

const routes = [
    {
        path: '/auth/sign_in',
        name: 'SignIn',
        component: LoginForm
    },
    // другие маршруты...
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router