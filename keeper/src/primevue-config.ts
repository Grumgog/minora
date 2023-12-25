import type { App } from "vue";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Card from "primevue/card";
import Menubar from "primevue/menubar";
import DataTable from "primevue/datatable";
import Column from "primevue/column";

export function usePrimeVueComponents(app: App) {
	const components = [Button, InputText, Dropdown, Card, Menubar, DataTable, Column];
	for (const component of components) {
		app.component(component.name, component);
	}
}
