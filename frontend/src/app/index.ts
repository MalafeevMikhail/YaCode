import { createApp } from "vue";
import { router, vuetify } from "./providers";
import App from "./app.vue";

export const app = createApp(App).use(router).use(vuetify);
