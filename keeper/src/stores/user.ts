import { ref, computed } from "vue";
import { defineStore } from "pinia";

export const useUserStore = defineStore("user", () => {
	const userLogin = ref("");
	const storagedToken = localStorage.getItem("auth");
	const APIToken = ref(storagedToken ?? "");
	const isLogined = computed(() => userLogin.value !== "" && APIToken.value !== "");

	function setTokenInfo(token: string): void {
		APIToken.value = token;
		localStorage.setItem("auth", token);
	}

	return { userLogin, APIToken, isLogined, setTokenInfo };
});
