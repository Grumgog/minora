<script setup lang="ts">
import UsersList from "@/components/UsersList.vue";
import type { UserInfo } from "@/fixtures";
import { useUsersStore } from "@/stores/users-store";
import { ref } from "vue";
const users = useUsersStore();
const data = ref<Array<UserInfo>>([]);
const isLoading = ref(false);
async function loadData() {
	try {
		isLoading.value = true;
		data.value = await users.getUsers()
	} finally {
		isLoading.value = false;
	}
}

loadData().then();
</script>

<template>
	<ProgressSpinner v-if="isLoading" />
	<UsersList v-else :data="data" />
</template>

<style scoped></style>
