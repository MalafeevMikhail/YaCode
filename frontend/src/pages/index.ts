import { type RouteRecordRaw } from "vue-router";
import { ROUTES } from "@/shared/constants/routes";

export const routes: RouteRecordRaw[] = [
  {
    path: ROUTES.HOME.path,
    meta: {
      auth: false,
      layout: "base",
    },
    component: () => import("./home"),
  },
];
