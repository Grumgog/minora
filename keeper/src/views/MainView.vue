<script setup lang="ts">
import TheMenubar from "@/components/TheMenubar.vue";
</script>

<template>
	<div class="main-content">
		<TheMenubar />
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
