
import { useAuthStore } from "@/stores/auth";

export async function handlerRequest(method, url, body = null, includeAuth = true) {
    const authStore = useAuthStore();
    try {
        const response = await fetch(`https://api-broken-thunder-9310.fly.dev/api/${url}`, {
            method,
            headers: {
                ...(includeAuth ? { "Authorization": `Bearer ${authStore.user.token}` } : {}),
                "Content-Type": "application/json",
            },
            body: body ? JSON.stringify(body) : null,
        });

        if (response.status == 204) {
            return;
        }

        const responseData = await response.json();

        if (!response.ok) {
            throw new Error(responseData.error);
        }

        return responseData;
    } catch (error) {
        throw error;
    }
}
