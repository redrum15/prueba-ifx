// stores/websocket.js
import { defineStore } from "pinia";
import { ref } from "vue";

export const useWebSocketStore = defineStore("websocket", () => {

    const socket = ref(null);
    const messages = ref([]);

    const connect = () => {
        if (socket.value) return;

        socket.value = new WebSocket("ws://localhost:3000/ws");

        socket.value.onmessage = (event) => {

            console.log(event);

            const data = JSON.parse(event.data);
            messages.value.push(data);
        };

        socket.value.onclose = () => {
            setTimeout(connect, 3000);
        };
    };

    return { messages, connect };
});
