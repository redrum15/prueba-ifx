<script setup>
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const logout = () => {
    authStore.logout();
    router.push('/login');
};

</script>

<template>
    <nav class="navbar bg-body-tertiary">
        <form class="container-fluid d-flex justify-content-between align-items-center  ">
            <p class="m-0">Prueba IFX - Daniel Alba</p>
            <div>
                <button v-if="authStore.user.user_type === 'admin'" class="btn btn-outline-success me-2" type="button"
                    data-bs-toggle="modal" data-bs-target="#createVMModal">
                    Create VM</button>
                <button @click="logout" class="btn btn-outline-secondary" type="button">Logout</button>
            </div>
        </form>
    </nav>
</template>
