<script setup>
import { useWebSocketStore } from "@/stores/socket";
import { computed, onMounted } from "vue";

const websocketStore = useWebSocketStore();
const lastMessage = computed(() => websocketStore.messages.at(-1));


onMounted(() => {
    websocketStore.connect();
});
</script>

<template>
    <div v-if="lastMessage" class="toast-container position-fixed top-0 end-0 p-3">
        <div class="toast show bg-info text-white" role="alert">
            <div class="toast-header">
                <strong class="me-auto">Notificación</strong>
                <button type="button" class="btn-close" @click="websocketStore.messages.pop()"
                    aria-label="Close"></button>
            </div>
            <div class="toast-body">
                {{ lastMessage.text }}
            </div>
        </div>
    </div>
</template>

<style scoped>
.toast-container {
    z-index: 1050;
}
</style>
