<script setup lang="ts">
import axios from "axios";
import { useRoute } from "vue-router";
import { useToString, useClipboard } from "@vueuse/core";

import { computed, onMounted } from "vue";

const route = useRoute();
const { text, copy, copied, isSupported } = useClipboard();

const roomId = useToString(() => route.params.roomId as string);
const url = computed(() => {
  return window.location.href;
});
const checkRoomId = () => {
  if (!roomId) window.location.replace("/");

  axios
    .post("/api/check-room", {
      roomId: String(roomId.value),
    })
    .then((data) => {
      if (!("isValid" in data.data)) {
        window.location.replace("/");
        return;
      }

      if (!data.data.isValid) window.location.replace("/");
    })
    .catch(() => {
      window.location.replace("/");
    });
};

onMounted(() => {
  checkRoomId();
});
</script>
<template>
  <div class="container">
    <div class="header">
      <div class="title">
        <span class="title-icon">◉</span>
        <span>YaCode</span>
      </div>
      <div class="controls">
        <select class="language-select">
          <option value="javascript">JavaScript</option>
          <option value="python">Python</option>
          <option value="java">Java</option>
          <option value="cpp">C++</option>
          <option value="go">Go</option>
          <option value="rust">Rust</option>
          <option value="php">PHP</option>
          <option value="csharp">C#</option>
          <option value="ruby">Ruby</option>
          <option value="swift">Swift</option>
          <option value="kotlin">Kotlin</option>
          <option value="typescript">TypeScript</option>
          <option value="html">HTML</option>
          <option value="css">CSS</option>
          <option value="sql">SQL</option>
          <option value="bash">Bash/Shell</option>
        </select>

        <div class="link-container" @click="copy(url)">
          <span class="link">
            {{ url }}
          </span>
        </div>

        <button class="copy-btn" @click="copy(url)">
          <i class="fas fa-copy"></i>
          Копировать
        </button>
      </div>
    </div>

    <div class="code-editor">
      <div class="editor-header">
        <div class="editor-dot red"></div>
        <div class="editor-dot yellow"></div>
        <div class="editor-dot green"></div>
        <div class="editor-title">main.js</div>
      </div>
      <div class="code-area" id="codeEditor" placeholder=""></div>
    </div>

    <div class="footer">
      YaCode &copy; 202Шесть — Совместный редактор кода в реальном времени
    </div>
  </div>
</template>
<style lang="scss" scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}

body {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  padding: 20px;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  overflow: hidden;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 30px;
  background: #1a1a2e;
  color: white;
}

.title {
  font-size: 28px;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 10px;
}

.title-icon {
  color: #764ba2;
  font-size: 32px;
}

.controls {
  display: flex;
  align-items: center;
  gap: 15px;
  flex-wrap: wrap;
}

.language-select {
  padding: 10px 15px;
  border-radius: 8px;
  border: 2px solid #667eea;
  background: white;
  color: #333;
  font-size: 14px;
  font-weight: 600;
  min-width: 180px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.language-select:hover {
  border-color: #764ba2;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
}

.language-select option {
  padding: 10px;
  font-weight: 600;
}

.link-container {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #f8f9fa;
  padding: 8px 15px;
  border-radius: 8px;
  border: 2px solid #e9ecef;
  cursor: pointer;
}

.link {
  color: #667eea;
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 300px;
}

.copy-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 8px;
}

.copy-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(102, 126, 234, 0.4);
}

.copy-btn:active {
  transform: translateY(0);
}

.copy-btn i {
  font-size: 16px;
}

.code-editor {
  height: 600px;
  background: #1e1e1e;
  margin: 20px;
  border-radius: 10px;
  overflow: hidden;
  position: relative;
  border: 3px solid #2d2d2d;
}

.editor-header {
  background: #252526;
  padding: 15px 20px;
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 2px solid #333;
}

.editor-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.red {
  background: #ff5f56;
}
.yellow {
  background: #ffbd2e;
}
.green {
  background: #27ca3f;
}

.editor-title {
  color: #ccc;
  font-size: 14px;
  font-weight: 600;
  margin-left: 10px;
}

.code-area {
  width: 100%;
  height: calc(100% - 50px);
  background: transparent;
  color: #d4d4d4;
  font-family: "Courier New", monospace;
  font-size: 16px;
  line-height: 1.6;
  padding: 25px;
  border: none;
  resize: none;
  outline: none;
  white-space: pre;
  overflow: auto;
}

.code-area::placeholder {
  color: #666;
  font-style: italic;
}

.footer {
  text-align: center;
  padding: 15px;
  color: #666;
  font-size: 14px;
  background: #f8f9fa;
  border-top: 2px solid #e9ecef;
}

@media (max-width: 768px) {
  .header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }

  .controls {
    justify-content: center;
  }

  .link {
    max-width: 200px;
  }

  .code-editor {
    height: 400px;
    margin: 10px;
  }
}
</style>
