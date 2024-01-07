<script setup lang="ts">
import { ref } from 'vue';

const visible = ref(true);

interface Link {
	url: string;
	text: string;
	icon: string;
}

const links = ref<Array<Link>>([
	{
		url: "/resource/edit/main",
		icon: "pi pi-file-edit",
		text: "Отредактируйте главную страницу"
	},
	{
		url: "/resources",
		icon: "pi pi-plus",
		text: "Добавте другие страницы"
	},
	{
		url: "/parameters",
		icon: "pi pi-cog",
		text: "Добавте новые параметры"
	},
	{
		url: "/resource/new",
		icon: "pi pi-plus",
		text: "Добавить новую запись"
	},
	{
		url: "http://www.example.com",
		icon: "pi pi-search",
		text: "Изучите работу с CMS Минора"
	},
	{
		url: "http://www.discord.com/server/join",
		icon: "pi pi-discord",
		text: "Присоединиться к серверу сообщества"
	},
]);

function onCloseClick() {
	visible.value = false;
}
</script>

<template>
	<Panel v-if="visible" class="p-2">
		<template #header>
			<div class="header-container">
				<div class="font-bold header">Добро пожаловать в Минору!</div>
				<div>Здесь вы можете найти несколько полезных ссылок</div>
			</div>
		</template>
		<template #icons>
			<Button icon="pi pi-times" text @click="onCloseClick()"></Button>
		</template>
		<div class="section-links">
			<div v-for="link in links" :key="link.url">
				<span class="p-1" :class="link.icon"></span>
				<a :href="link.url">{{ link.text }}</a>
			</div>
		</div>
	</Panel>
</template>

<style scoped>
.header {
	color: var(--text-color);
}
.header-container {
	display: flex;
	flex-direction: column;
}

.section-links {
	display: grid;
	grid-template-columns: auto auto;
	gap: 5px;
}
</style>