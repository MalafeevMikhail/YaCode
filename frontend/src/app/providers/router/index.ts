import { createRouter, createWebHistory } from "vue-router";

import { routes } from "@/pages";
import { uploadLayouts } from "./lib";

export const router = createRouter({
  history: createWebHistory(),
  strict: true,
  routes,
});

[uploadLayouts].forEach(router.beforeEach);

export { RouteLayout } from "./ui/route-layout";
