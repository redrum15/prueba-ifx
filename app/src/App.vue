<script setup>
import { onMounted } from "vue";
import { useWebSocketStore } from "@/stores/socket";
import NotificationToast from "@/components/NotificationToast.vue";



</script>

<template>
  <router-view />
  <NotificationToast />

</template>
