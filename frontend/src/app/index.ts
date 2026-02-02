import { createApp } from "vue";
import { router, vuetify } from "./providers";
import App from "./app.vue";

import "@/shared/styles/index.scss";

export const app = createApp(App).use(router).use(vuetify);
