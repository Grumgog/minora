import { createApp } from "vue";
import { createPinia } from "pinia";

import PrimeVue from "primevue/config";
import { usePrimeVueComponents } from "./primevue-config";
import "primevue/resources/themes/soho-light/theme.css";
import "primeicons/primeicons.css";
import "primeflex/primeflex.css";
import i18n from "./i18n-app";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(PrimeVue);
app.use(i18n);
usePrimeVueComponents(app);

app.mount("#app");
