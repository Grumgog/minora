import { ref, computed } from "vue";
import { defineStore } from "pinia";

export const useDeviceDetector = defineStore("device-detector", () => {
	// TODO: replace fake method on computed props
	function isMobile(): boolean {
		return true;
	}

	function isDesktop(): boolean {
		return false;
	}
	return { isMobile, isDesktop };
});
