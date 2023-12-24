import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import type { App } from "vue";

export function usePrimeVueComponents(app: App) {
	const components = [Button, InputText, Dropdown];
	for (const component of components) {
		app.component(component.name, component);
	}
}
