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

const data = ref("")
const modalRef = ref(null);
let modalInstance = null;


const websocketStore = useWebSocketStore();
const lastMessage = computed(() => websocketStore.messages.at(-1));



const handleVMDelete = async () => {
    await handlerRequest("DELETE", API_URLS.VMS.DELETE(route.params.id));
    goBack();
}


onMounted(async () => {
    modalInstance = new Modal(modalRef.value);

    data.value = await handlerRequest("GET", API_URLS.VMS.DETAIL(route.params.id));
});

const goBack = () => {
    router.push("/dashboard");
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
    <div class="container d-flex justify-content-between">
        <div>
            <p>Name: {{ data.name }}</p>
            <p>Cores: {{ data.cores }}</p>
            <p>Ram: {{ data.ram }}</p>
            <p>OS: {{ data.os }}</p>
            <p>Status: {{ data.status }}</p>
        </div>
        <div v-if="authStore.user.user_type === 'admin'">
            <button @click="goBack()" class="btn btn-link">EDIT</button>
            <button class="btn btn-link" data-bs-toggle="modal" data-bs-target="#deleteVMModal">DELETE</button>
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
