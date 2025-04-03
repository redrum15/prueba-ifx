<script setup>
import { useRoute, useRouter, onBeforeRouteLeave } from "vue-router";
import { useAuthStore } from '@/stores/auth';
import { handlerRequest } from "@/rest/index";
import { useWebSocketStore } from "@/stores/socket";
import { onMounted, ref, computed } from "vue";
import { API_URLS } from "@/rest/urls";
import { Modal } from 'bootstrap';
import Navbar from "@/components/Navbar.vue";

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();

const data = ref(null);
const isEditing = ref(false);
const modalRef = ref(null);
let modalInstance = null;

const websocketStore = useWebSocketStore();
const lastMessage = computed(() => websocketStore.messages.at(-1));

const handleVMDelete = async () => {
    await handlerRequest("DELETE", API_URLS.VMS.DELETE(route.params.id));
    goBack();
};

const startEditing = () => {
    isEditing.value = true;
};

const saveChanges = async () => {
    await handlerRequest("PUT", API_URLS.VMS.EDIT(route.params.id), data.value);
    isEditing.value = false;
};

onMounted(async () => {
    modalInstance = new Modal(modalRef.value);
    data.value = await handlerRequest("GET", API_URLS.VMS.DETAIL(route.params.id));
});

const goBack = () => {
    router.push("/");
};

onBeforeRouteLeave((to, from, next) => {
    if (modalInstance) {
        modalInstance.hide();
    }
    next();
});
</script>

<template>
    <Navbar />

    <button @click="goBack()" class="btn btn-link">&lt; back</button>
    <div class="container d-flex justify-content-between mt-4">
        <div>
            <template v-if="!isEditing">
                <p>Name: {{ data?.name }}</p>
                <p>Cores: {{ data?.cores }}</p>
                <p>Ram: {{ data?.ram }}</p>
                <p>OS: {{ data?.os }}</p>
                <p>Status: {{ data?.status }}</p>
            </template>
            <template v-else>
                <div class="mb-2">
                    <label class="form-label">Name</label>
                    <input v-model="data.name" type="text" class="form-control">
                </div>
                <div class="mb-2">
                    <label class="form-label">Cores</label>
                    <input v-model="data.cores" type="number" class="form-control">
                </div>
                <div class="mb-2">
                    <label class="form-label">Ram</label>
                    <input v-model="data.ram" type="number" class="form-control">
                </div>
                <div class="mb-2">
                    <label class="form-label">OS</label>
                    <input v-model="data.os" type="text" class="form-control">
                </div>
                <div class="mb-2">
                    <label class="form-label">Status</label>
                    <input v-model="data.status" type="text" class="form-control">
                </div>
            </template>
        </div>

        <div v-if="authStore.user.user_type === 'admin'">
            <button v-if="!isEditing" @click="startEditing" class="btn btn-warning me-2">EDIT</button>
            <button v-if="isEditing" @click="saveChanges" class="btn btn-success me-2">SAVE</button>
            <button class="btn btn-danger" data-bs-toggle="modal" data-bs-target="#deleteVMModal">DELETE</button>
        </div>
    </div>

    <div class="modal" tabindex="-1" id="deleteVMModal" aria-labelledby="deleteVMModal" aria-hidden="true"
        ref="modalRef">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Delete VM</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <p>Are you sure you want to delete this virtual machine?</p>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                    <button @click="handleVMDelete" type="button" class="btn btn-primary btn-danger">Delete</button>
                </div>
            </div>
        </div>
    </div>
</template>
