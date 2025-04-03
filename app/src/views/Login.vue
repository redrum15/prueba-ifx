<script setup>
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';
import { ref } from 'vue';

const authStore = useAuthStore();
const router = useRouter();

const formData = ref({
    email: "",
    password: "",
});

const handleSubmit = async () => {
    try {
        const response = await fetch("http://localhost:3000/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(formData.value),
        });

        if (response.ok) {
            const data = await response.json();
            console.log(data);

            authStore.login(data);
            console.log("esta casi")
            router.push('/dashboard');
        } else {
            alert("Error al enviar el formulario");
        }
    } catch (error) {
        console.error("Error en la solicitud:", error);
        alert("No se pudo conectar con el servidor");
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