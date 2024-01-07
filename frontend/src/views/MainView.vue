<script setup lang="ts">
import TheMenubar from "@/components/TheMenubar.vue";
</script>

<template>
	<div class="main-content">
		<header class="header">
			<TheMenubar />
		</header>
		<main>
			<RouterView />
		</main>
	</div>
</template>

<style scoped>
.header {
	top: 0;
	z-index: 9999;
	position: sticky;
}

.main-content {
	flex-grow: 1;
}
</style>
