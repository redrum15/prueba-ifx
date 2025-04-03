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
    <form @submit.prevent="handleSubmit">
        <div data-mdb-input-init class="form-outline mb-4">
            <input v-model="formData.email" type="email" id="email" class="form-control" required />
            <label class="form-label" for="email">Email address</label>
        </div>
        <div data-mdb-input-init class="form-outline mb-4">
            <input v-model="formData.password" type="password" id="password" class="form-control" required />
            <label class="form-label" for="password">Password</label>
        </div>
        <button type="submit" data-mdb-button-init data-mdb-ripple-init class="btn btn-primary btn-block mb-4">Sign
            in</button>
    </form>
</template>