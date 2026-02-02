import type { RouteLocationNormalized } from "vue-router";

export const uploadLayouts = async (route: RouteLocationNormalized) => {
  const layout = route.meta.layout;

  if (!layout) return;

  const component = await import(`@/widgets/layouts/${layout}/index.ts`);

  route.meta.layout = component.default;
};
