import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Card from "primevue/card";
import type { App } from "vue";
import Menubar from "primevue/menubar";

export function usePrimeVueComponents(app: App) {
	const components = [Button, InputText, Dropdown, Card, Menubar];
	for (const component of components) {
		app.component(component.name, component);
	}
}
