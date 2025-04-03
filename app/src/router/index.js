import { createRouter, createWebHistory } from 'vue-router';
import Login from '@/views/Login.vue';
import Dashboard from '@/views/Dashboard.vue';
import { useAuthStore } from '@/stores/auth';
import DetailVm from '@/views/DetailVm.vue';

const routes = [
    { path: '/login', component: Login },
    {
        path: '/',
        component: Dashboard,
        meta: { requiresAuth: true }
    },
    {
        path: '/vm/:id',
        component: DetailVm,
        meta: { requiresAuth: true }
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();

    if (to.meta.requiresAuth && !authStore.isAuthenticated) {
        next('/login');
    } else {
        next();
    }
});

export default router;
