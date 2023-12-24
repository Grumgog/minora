<script setup lang="ts">
import i18n from "@/i18n-app";
import Dropdown from "primevue/dropdown";
import { ref } from "vue";

const options = ref([
	{
		locale: "ru",
		title: "Русский",
	},
	{
		locale: "en",
		title: "English",
	},
]);
</script>

<template>
	<Dropdown :options="options" option-label="title" option-value="locale" v-model="$i18n.locale"></Dropdown>
</template>

<style scoped></style>
