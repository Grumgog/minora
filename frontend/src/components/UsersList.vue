<script setup lang="ts">
import type { UserInfo } from "@/fixtures";

const props = defineProps({
	data: {
		type: Array<UserInfo>,
		required: false,
	},
});
</script>

<template>
	<div v-for="item in props.data" v-bind:key="item.login">
		{{ item.login }}
	</div>
</template>

<style scoped></style>
