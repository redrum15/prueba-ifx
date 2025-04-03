<script setup>
import { handlerRequest } from "@/rest/index";
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';
import { onMounted, ref, computed } from 'vue';
import { formatDistanceToNow, parseISO } from "date-fns";
import { API_URLS } from "@/rest/urls";
import { Modal } from "bootstrap";
import Navbar from '@/components/Navbar.vue';

const router = useRouter();
const toastRef = ref(null);
const toastMessage = ref("");
const showToastFlag = ref(false);
const authStore = useAuthStore();
const modalRef = ref(null);

const items = ref([]);
const formData = ref({
    name: "",
    cores: "",
    ram: "",
    os: "",
    status: ""
});

const formattedItems = computed(() =>
    items.value.map(item => ({
        ...item,
        created_at: formatDistanceToNow(parseISO(item.created_at), { addSuffix: true })
    }))
);

const addNewVMToList = (newVM) => {
    items.value.unshift(newVM);
};

const handleVMCreate = async () => {
    try {
        const newVM = await handlerRequest("POST", API_URLS.VMS.DEFAULT, formData.value);
        addNewVMToList(newVM);

        const modalInstance = Modal.getInstance(modalRef.value) || new Modal(modalRef.value);
        modalInstance.hide();

        formData.value = { name: "", cores: "", ram: "", os: "", status: "" };

    } catch (error) {
        toastMessage.value = error.message;
        showToastFlag.value = true;
    }
};

onMounted(async () => {
    try {
        items.value = await handlerRequest("GET", API_URLS.VMS.DEFAULT);
    } catch (error) {
        toastMessage.value = error.message;
        showToastFlag.value = true;
    }
});

const goToDetail = (id) => {
    router.push(`/vm/${id}`);
};
</script>

<template>
    <div class="toast-container position-fixed top-0 end-0 p-3" :hidden="!showToastFlag">
        <div id="liveToast" class="toast d-flex bg-danger" role="alert" ref="toastRef" aria-live="assertive"
            aria-atomic="true">
            <div class="toast-body text-white fw-bolder">
                {{ toastMessage }}
            </div>
            <button @click="showToastFlag = false" type="button" class="btn-close me-2 m-auto" data-bs-dismiss="toast"
                aria-label="Close"></button>
        </div>
    </div>

    <div class="modal fade" id="createVMModal" tabindex="-1" aria-labelledby="createVMModalLabel" aria-hidden="true"
        ref="modalRef">
        <div class="modal-dialog">
            <div class="modal-content">
                <form @submit.prevent="handleVMCreate" class="p-3">
                    <div class="mb-3">
                        <label for="name" class="form-label">Name</label>
                        <input v-model="formData.name" type="text" class="form-control" id="name" required>
                    </div>
                    <div class="mb-3">
                        <label for="cores" class="form-label">Cores</label>
                        <input v-model="formData.cores" type="number" class="form-control" id="cores" required>
                    </div>
                    <div class="mb-3">
                        <label for="ram" class="form-label">Ram</label>
                        <input v-model="formData.ram" type="number" class="form-control" id="ram" required>
                    </div>
                    <div class="mb-3">
                        <label for="os" class="form-label">OS</label>
                        <input v-model="formData.os" type="text" class="form-control" id="os" required>
                    </div>
                    <div class="mb-3">
                        <label for="status" class="form-label">Status</label>
                        <input v-model="formData.status" type="text" class="form-control" id="status" required>
                    </div>

                    <button type="submit" class="btn btn-primary">Submit</button>
                </form>
            </div>
        </div>
    </div>

    <Navbar />

    <p class="p-2">User type: {{ authStore.user.user_type }}</p>
    <h1 class="mt-4">Virtual Machines:</h1>
    <div class="list-group mt-1">
        <a v-for="item in formattedItems" :key="item.id" href="#" class="list-group-item list-group-item-action"
            aria-current="true" @click="goToDetail(item.id)">
            <div class="d-flex w-100 justify-content-between">
                <div>
                    <small>{{ item.id }}</small>
                    <h5 class="mb-1">{{ item.name }}</h5>
                </div>
                <small>{{ item.created_at }}</small>
            </div>
            <p class="mb-1">Cores: {{ item.cores }} - Ram: {{ item.ram }}</p>
            <p>OS: {{ item.os }} - Status: {{ item.status }}</p>
        </a>
    </div>
</template>
