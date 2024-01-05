<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { ref, watch, type Ref } from "vue";
import type { MenuItem } from "primevue/menuitem";
import { useUserStore } from "@/stores/user";
import { useRouter } from "vue-router";

const router = useRouter();
const { t, locale } = useI18n();
const user = useUserStore();
const items: Ref<any> = ref(initMenu({ userLogin: user.userLogin }));

watch(locale, () => {
	items.value = initMenu({ userLogin: user.userLogin });
});

function initMenu(context: { userLogin: string }): Array<MenuItem> {
	return [
		{
			label: t("dashboard.menu.site.title"),
			icon: "pi pi-home",
			items: [
				{
					label: t("dashboard.menu.site.main"),
					icon: "pi pi-th-large",
				},
				{
					label: t("dashboard.menu.site.view"),
					icon: "pi pi-eye",
				},
				{
					label: t("dashboard.menu.site.clearCache"),
					icon: "pi pi-eraser",
				},
				{
					label: t("dashboard.menu.site.find"),
					icon: "pi pi-search",
				},
				{
					label: t("dashboard.menu.site.newResource"),
					icon: "pi pi-file",
				},
				{
					label: t("dashboard.menu.site.newLink"),
					icon: "pi pi-link",
				},
			],
		},
		{
			label: t("dashboard.menu.elements"),
			icon: "pi pi-th-large",
		},
		{
			label: t("dashboard.menu.modules"),
			icon: "pi pi-box",
		},
		{
			label: t("dashboard.menu.users"),
			icon: "pi pi-users",
			url: "/users",
		},
		{
			label: t("dashboard.menu.tools"),
			icon: "pi pi-wrench",
		},
		{
			label: t("dashboard.menu.reports"),
			icon: "pi pi-file",
		},
		{
			label: context.userLogin,
			icon: "pi pi-user",
			items: [
				{
					label: t("dashboard.menu.user.changePassword"),
				},
				{
					label: t("dashboard.menu.user.logOut"),
				},
			],
		},
	];
}
</script>

<template>
	<Menubar class="last-item-to-right" :model="items"> </Menubar>
</template>

<style scoped lang="scss">
@media screen and (min-width: 960px) {
	.last-item-to-right :deep(.p-menubar-root-list > li.p-menuitem:last-child) {
		position: absolute;
		right: 1rem;
		.p-submenu-list {
			transform: translateX(-65px);
		}
	}
}
</style>
