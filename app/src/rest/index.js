
import { API_URL } from "@/config";
import { useAuthStore } from "@/stores/auth";

export async function handlerRequest(method, route, body = null, includeAuth = true) {
    const authStore = useAuthStore();
    try {
        const response = await fetch(`${API_URL}/api/${route}`, {
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
