<script setup lang="ts">
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import { useRouter } from "vue-router";
import { useDeviceDetector } from "@/stores/device-detector";

const userStore = useUserStore();
const router = useRouter();

const login = ref("");
const password = ref("");

function onsubmit() {
	// TODO: REPLACE FAKE LOGIN TO NORMAL AUTHORIZATION
	userStore.userLogin = "admin";
	userStore.setTokenInfo("fake token");
	console.log("to cms~");
	router.push("/");
}

const deviceDetector = useDeviceDetector();
</script>

<template>
	<div class="container" :class="{ 'is-mobile-container': deviceDetector.isMobile()}">
		<form class="login-form" @submit.prevent="onsubmit">
			<div class="field">
				<label for="login">{{ $t("login.loginLabel") }}</label>
				<InputText id="login" required v-model="login"></InputText>
			</div>
			<div class="field">
				<label for="password">{{ $t("login.passwordLabel") }}</label>
				<InputText id="password" type="password" required v-model="password"></InputText>
			</div>

			<Button class="w-full" :label="$t('login.submit')" type="submit"></Button>
		</form>
	</div>
</template>

<style scoped>
.container {
	display: grid;
	grid-template-columns: 2rem auto 1fr;
	background-image: url("@/assets/img/login-background.jpg");
}

.is-mobile-container {
	grid-template-columns: 1fr auto 1fr;
}

.login-form {
	background-color: var(--surface-900);
	height: 100%;
	grid-column: 2/3;
	padding: 2rem;
	display: flex;
	justify-items: center;
	align-items: center;
	flex-direction: column;
	justify-content: center;
}

label {
	color: var(--primary-color-text);
}

.field {
	display: flex;
	flex-direction: column;
	gap: 5px;
}
</style>
