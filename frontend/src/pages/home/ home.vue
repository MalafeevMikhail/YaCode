<script setup lang="ts">
import { ref } from "vue";
import { YaModal } from "@/shared/components/modal";
import axios from "axios";

const isLoadingModal = ref(false);
const isCreateRoomModalVisible = ref(true);
const roomName = ref("");

function openCreateModal() {
  isLoadingModal.value = true;
  setTimeout(() => {
    isLoadingModal.value = false;
    isCreateRoomModalVisible.value = true;
  }, 500);
}

async function createRoom() {
  if (!roomName.value) return;

  axios
    .post("/api/create-room", {
      name: roomName.value,
    })
    .then((data) => {
      if (!("id" in data.data)) {
        alert("Ошибка!. Проверьте логи");
        console.error("Отсутствует ссылка в ответе сервера");
        return;
      }

      window.location.href = `room/${data.data.id}`;
    })
    .catch((err) => {
      alert("Ошибка! Проверьте логи.");
      console.error(err);
    })
    .finally(() => {
      isCreateRoomModalVisible.value = false;
    });
}
</script>
<template>
  <div class="create-block">
    <h3 class="mb-4">You have to create a room for live coding to begin.</h3>
    <v-btn :loading="isLoadingModal" @click="openCreateModal">
      Create room
    </v-btn>
  </div>

  <ya-modal v-model="isCreateRoomModalVisible" v-if="isCreateRoomModalVisible">
    <div class="d-flex justify-center align-center flex-column ga-3">
      <h2>Please specify the room name.</h2>
      <input type="text" v-model="roomName" />
      <v-btn :disabled="!roomName" @click="createRoom">Create</v-btn>
    </div>
  </ya-modal>
</template>

<style lang="scss" scoped>
.create-block {
  margin: 30px auto;
  max-width: 600px;
  text-align: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
$primary-color: #1976d2;
$primary-light: #42a5f5;
$primary-dark: #0d47a1;
$text-color: #333;
$light-bg: #f5f7fa;
$border-radius: 12px;
$shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
$transition: all 0.3s ease;

.introduction-container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  gap: 24px;
  min-height: 400px;
  padding: 40px;
  background: linear-gradient(135deg, $light-bg 0%, #e8eaf6 100%);
  border-radius: 20px;
  box-shadow: $shadow;
  max-width: 500px;
  margin: 40px auto;
  position: relative;
  overflow: hidden;

  &::before {
    content: "";
    position: absolute;
    top: -50px;
    right: -50px;
    width: 150px;
    height: 150px;
    background: linear-gradient(45deg, $primary-light, transparent 70%);
    border-radius: 50%;
    opacity: 0.3;
  }

  &::after {
    content: "";
    position: absolute;
    bottom: -30px;
    left: -30px;
    width: 100px;
    height: 100px;
    background: linear-gradient(45deg, transparent, $primary-light 70%);
    border-radius: 50%;
    opacity: 0.2;
  }
}

h2 {
  color: $primary-dark;
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 8px;
  text-align: center;
  position: relative;
  z-index: 1;

  &::after {
    content: "";
    position: absolute;
    bottom: -8px;
    left: 50%;
    transform: translateX(-50%);
    width: 60px;
    height: 3px;
    background: linear-gradient(90deg, $primary-color, $primary-light);
    border-radius: 2px;
  }
}

input {
  width: 100%;
  max-width: 320px;
  padding: 16px 20px;
  font-size: 16px;
  color: $text-color;
  background: white;
  border: 2px solid #e0e0e0;
  border-radius: $border-radius;
  outline: none;
  transition: $transition;
  position: relative;
  z-index: 1;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);

  &:focus {
    border-color: $primary-color;
    box-shadow: 0 0 0 3px rgba($primary-color, 0.2);
    transform: translateY(-2px);
  }

  &:hover {
    border-color: $primary-light;
  }

  &::placeholder {
    color: #9e9e9e;
  }
}

.v-btn {
  width: 100%;
  max-width: 320px;
  font-size: 16px !important;
  font-weight: 600 !important;
  text-transform: uppercase !important;
  letter-spacing: 0.5px !important;
  border-radius: $border-radius !important;
  transition: $transition !important;
  position: relative;
  z-index: 1;
  overflow: hidden;

  &:not(:disabled) {
    background: linear-gradient(
      135deg,
      $primary-color,
      $primary-light
    ) !important;
    color: white !important;
    box-shadow: 0 6px 20px rgba($primary-color, 0.3) !important;

    &:hover {
      transform: translateY(-3px);
      box-shadow: 0 10px 25px rgba($primary-color, 0.4) !important;

      &::before {
        transform: translateX(100%);
      }
    }

    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: -100%;
      width: 100%;
      height: 100%;
      background: linear-gradient(
        90deg,
        transparent,
        rgba(255, 255, 255, 0.2),
        transparent
      );
      transition: transform 0.6s ease;
    }
  }

  &:disabled {
    background: #e0e0e0 !important;
    color: #9e9e9e !important;
    cursor: not-allowed !important;
    box-shadow: none !important;
  }
}

@media (max-width: 600px) {
  .introduction-container {
    padding: 30px 20px;
    margin: 20px;
    min-height: 350px;
  }

  h2 {
    font-size: 24px;
  }

  input,
  .v-btn {
    max-width: 100%;
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.introduction-container {
  animation: fadeInUp 0.6s ease-out;
}
</style>
