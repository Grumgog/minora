<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref, watch, type Ref } from "vue";
import type { MenuItem } from "primevue/menuitem";
const { t, locale } = useI18n();
const items: Ref<any> = ref(initMenu());
watch(locale, () => {
	items.value = initMenu();
});

function initMenu(): Array<MenuItem> {
	return [
		{
			label: t("dashboard.menu.site.title"),
			icon: "pi pi-home",
			items: [
				{
					label: t("dashboard.menu.site.main"),
					icon: "pi pi-th-large"
				},
				{
					label: t("dashboard.menu.site.view"),
					icon: "pi pi-eye"
				},
				{
					label: t("dashboard.menu.site.clearCache"),
					icon: "pi pi-eraser"
				},
				{
					label: t("dashboard.menu.site.find"),
					icon: "pi pi-search",
				},
				{
					label: t("dashboard.menu.site.newResource"),
					icon: "pi pi-file"
				},
				{
					label: t("dashboard.menu.site.newLink"),
					icon: "pi pi-link"
				}
			]
		},
		{
			label: t("dashboard.menu.elements"),
			icon: "pi pi-th-large"
		},
		{
			label: t("dashboard.menu.modules"),
			icon: "pi pi-box"
		},
		{
			label: t("dashboard.menu.users"),
			icon: "pi pi-users"
		},
		{
			label: t("dashboard.menu.tools"),
			icon: "pi pi-wrench"
		},
		{
			label: t("dashboard.menu.reports"),
			icon: "pi pi-file"
		}
	]
}
</script>

<template>
	<div class="main-content">
		<Menubar :model="items"></Menubar>
		<header>HEADER</header>
		<main>
			<RouterView />
		</main>
		<footer>footer</footer>
	</div>
</template>

<style scoped>
.main-content {
	flex-grow: 1;
}
</style>
