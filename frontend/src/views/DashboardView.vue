<script setup lang="ts">
import LastUserInfo from "@/components/LastUserInfo.vue";
import TheWelcomeSplash from "@/components/TheWelcomeSplash.vue";
</script>

<template>
	<div class="flex flex-column align-items-center">
		<TheWelcomeSplash />
		<div class="container">
			<div class="debug-under-construction">PLACEHOLDER</div>
			<LastUserInfo />
			<div class="debug-under-construction">PLACEHOLDER</div>
		</div>
	</div>
</template>

<style scoped>
.container {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 1rem;
}
</style>
