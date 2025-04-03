import { defineStore } from "pinia";
import { ref } from "vue";
import { WS_URL } from "@/config";

export const useWebSocketStore = defineStore("websocket", () => {

    const socket = ref(null);
    const messages = ref([]);

    const connect = () => {
        if (socket.value) return;

        socket.value = new WebSocket(`${WS_URL}/ws`);
        socket.value.onmessage = (event) => {
            const data = JSON.parse(event.data);
            messages.value.push(data);
        };

        socket.value.onclose = () => {
            setTimeout(connect, 3000);
        };
    };

    return { messages, connect };
});
