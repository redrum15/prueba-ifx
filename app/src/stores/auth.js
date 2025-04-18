import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')) || null,
    }),
    getters: {
        isAuthenticated: (state) => !!state.user,
    },
    actions: {
        login(userData) {
            this.user = userData;
            localStorage.setItem('user', JSON.stringify(userData));
        },
        logout() {
            localStorage.removeItem('user');
            this.user = null;
        },
    },
});
