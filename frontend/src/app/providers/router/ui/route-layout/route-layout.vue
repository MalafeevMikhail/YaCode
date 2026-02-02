<script lang="ts" setup>
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<template>
  <component :is="route.meta.layout || 'div'">
    <template v-for="(_, name) in $slots" #[name]="slotData">
      <slot :name="name" v-bind="slotData || {}"></slot>
    </template>
  </component>
</template>
