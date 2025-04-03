<script setup>
import { useWebSocketStore } from "@/stores/socket";
import { computed, onMounted } from "vue";

const websocketStore = useWebSocketStore();
const lastMessage = computed(() => websocketStore.messages.at(-1));

const toastClass = computed(() => {
    return {
        "bg-success": lastMessage.value?.event_type === "created",
        "bg-danger": lastMessage.value?.event_type === "deleted",
        "bg-warning": lastMessage.value?.event_type === "updated"
    };
});

onMounted(() => {
    websocketStore.connect();
});
</script>

<template>
    <div v-if="lastMessage" class="toast-container position-fixed top-0 end-0 p-3">
        <div id="liveToast" class="toast d-flex" :class="toastClass" role="alert" ref="toastRef" aria-live="assertive"
            aria-atomic="true">
            <div class="toast-body text-white fw-bolder">
                {{ lastMessage.message }}
            </div>
            <button @type="button" class="btn-close me-2 m-auto" @click="websocketStore.messages.pop()"
                aria-label="Close"></button>
        </div>
    </div>
</template>

<style scoped>
.toast-container {
    z-index: 1050;
}
</style>
