<script setup lang="ts">
import LastUserInfo from "@/components/LastUserInfo.vue";
</script>

<template>
	<div class="container">
		<div class="debug-under-construction">PLACEHOLDER</div>
		<LastUserInfo />
		<DIV class="debug-under-construction">PLACEHOLDER</DIV>
	</div>
</template>

<style scoped>
.container {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 1rem;
}
</style>
