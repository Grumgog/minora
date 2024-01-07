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
		text: "welcomeSplash.links.editMain"
	},
	{
		url: "/resources",
		icon: "pi pi-plus",
		text: "welcomeSplash.links.addNewResource"
	},
	{
		url: "/parameters",
		icon: "pi pi-cog",
		text: "welcomeSplash.links.addNewParameter"
	},
	{
		url: "/resource/new",
		icon: "pi pi-plus",
		text: "welcomeSplash.links.addNewRecord"
	},
	{
		url: "http://www.example.com",
		icon: "pi pi-search",
		text: "welcomeSplash.links.learn"
	},
	{
		url: "http://www.discord.com/server/join",
		icon: "pi pi-discord",
		text: "welcomeSplash.links.communityJoin"
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
				<div class="font-bold header">{{$t("welcomeSplash.header")}}</div>
				<div>{{ $t("welcomeSplash.secondaryText") }}</div>
			</div>
		</template>
		<template #icons>
			<Button icon="pi pi-times" text @click="onCloseClick()"></Button>
		</template>
		<div class="section-links">
			<div v-for="link in links" :key="link.url">
				<span class="p-1" :class="link.icon"></span>
				<a :href="link.url">{{ $t(link.text) }}</a>
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