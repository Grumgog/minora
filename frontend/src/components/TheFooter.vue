<script setup lang="ts">
import LocaleSelector from "./LocaleSelector.vue";
import TheVersion from "./TheVersion.vue";
</script>

<template>
	<div class="footer-container">
		<div>{{ $t("credits.author") }}</div>
		<LocaleSelector />
		<TheVersion />
	</div>
</template>

<style scoped>
.footer-container {
	display: flex;
	gap: 1rem;
	width: 100%;
	background-color: var(--primary-900);
	color: var(--primary-color-text);
}
</style>
