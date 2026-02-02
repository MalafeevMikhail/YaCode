<script setup lang="ts">
import { ref } from "vue";

const theme = ref("light");
const switchedTheme = ref(false);

function updateTheme() {
  theme.value = switchedTheme.value ? "dark" : "light";
}
</script>
<template>
  <v-app :theme="theme">
    <v-responsive class="border rounded">
      <v-app-bar class="px-3 d-flex justify-end" floating>
        <v-switch
          :theme="theme"
          prepend-icon="mdi-weather-sunny"
          append-icon="mdi-weather-night"
          v-model="switchedTheme"
          @change="updateTheme"
        ></v-switch>
      </v-app-bar>
      <v-main> <slot></slot> </v-main>
    </v-responsive>
  </v-app>
</template>

<style lang="scss" scoped></style>
