import type { App } from "vue";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Card from "primevue/card";
import Menubar from "primevue/menubar";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import ProgressSpinner from "primevue/progressspinner";
import Panel from "primevue/panel";
import ScrollPanel from "primevue/scrollpanel";

export function usePrimeVueComponents(app: App) {
	const components = [
		Button,
		InputText,
		Dropdown,
		Card,
		Panel,
		Menubar,
		DataTable,
		Column,
		ProgressSpinner,
		ScrollPanel,
	];
	for (const component of components) {
		app.component(component.name, component);
	}
}
