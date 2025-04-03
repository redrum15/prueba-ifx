<script setup>
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';
import { ref } from 'vue';
import { API_URLS } from "@/rest/urls";
import { handlerRequest } from "@/rest/index";


const authStore = useAuthStore();
const router = useRouter();

const formData = ref({
    email: "",
    password: "",
});

const handleSubmit = async () => {
    try {
        const response = await handlerRequest("POST", API_URLS.AUTH.LOGIN, formData.value, false);
        authStore.login(response);
        router.push('/');
    } catch (error) {
        console.log(error);
        alert("Invalid credentials");
    }
};
</script>

<template>
    <div class="d-flex justify-content-center align-items-center vh-100">
        <form @submit.prevent="handleSubmit" class="p-4 border rounded shadow bg-white" style="width: 350px;">
            <h3 class="text-center mb-4">Sign In</h3>
            <div class="form-outline mb-4">
                <label class="form-label" for="email">Email</label>
                <input v-model="formData.email" type="email" id="email" class="form-control" required />
            </div>
            <div class="form-outline mb-4">
                <label class="form-label" for="password">Password</label>
                <input v-model="formData.password" type="password" id="password" class="form-control" required />
            </div>
            <button type="submit" class="btn btn-primary btn-block w-100">Sign in</button>
        </form>
    </div>
</template>