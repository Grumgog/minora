<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import TheFooter from "@/components/TheFooter.vue";
</script>

<template>
	<div class="container">
		<RouterView />
		<TheFooter class="footer" />
	</div>
</template>

<style scoped>
.container {
	height: 100%;
	display: flex;
	flex-direction: column;
}
</style>
