<script setup lang="ts">
import TheMenubar from "@/components/TheMenubar.vue";
</script>

<template>
	<div class="main-content">
		<header class="sticky">
			<TheMenubar />
		</header>
		<main>
			<RouterView />
		</main>
	</div>
</template>

<style scoped>
.sticky {
	position: sticky;
	top: 0;
}

.main-content {
	flex-grow: 1;
}
</style>
