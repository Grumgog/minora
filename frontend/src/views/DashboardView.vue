<script setup lang="ts">
import LastUserInfo from "@/components/LastUserInfo.vue";
import TheGeneralInfo from "@/components/TheGeneralInfo.vue";
import TheUpdateNews from "@/components/TheUpdateNews.vue";
import TheWelcomeSplash from "@/components/TheWelcomeSplash.vue";
</script>

<template>
	<div class="flex flex-column align-items-center p-2">
		<TheWelcomeSplash />
		<div class="container">
			<TheUpdateNews />
			<LastUserInfo />
			<TheGeneralInfo />
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
