<script setup lang="ts">
const modelValue = defineModel();

function closeModal() {
  modelValue.value = false;
}
</script>
<template>
  <div :class="['y-modal-overlay', { show: modelValue }]">
    <div class="modal-content" v-click-outside="closeModal">
      <slot />
    </div>
  </div>
</template>
<style lang="scss" scoped>
.y-modal-overlay {
  display: none;
  position: fixed;
  z-index: 1000;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);

  &.show {
    display: flex;
  }
}

.modal-content {
  background-color: #fefefe;
  padding: 20px;
  border: 1px solid #888;
  width: auto;
  height: auto;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
