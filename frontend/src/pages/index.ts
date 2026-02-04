import { type RouteRecordRaw } from "vue-router";
import { ROUTES } from "@/shared/constants/routes";

export const routes: RouteRecordRaw[] = [
  {
    path: ROUTES.HOME.path,
    meta: {
      layout: "base",
    },
    component: () => import("./home"),
  },
  {
    path: ROUTES.ROOM.path,
    meta: {
      layout: "base",
    },
    component: () => import("./room"),
  },
];
