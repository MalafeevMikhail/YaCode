<script setup lang="ts">
import axios from "axios";
import { useRoute } from "vue-router";
import { useToString, useClipboard } from "@vueuse/core";
import { computed, onMounted, ref, watch, nextTick } from "vue";
import hljs from "highlight.js";
import "highlight.js/styles/github-dark.css";

const route = useRoute();
const { copy } = useClipboard();

const roomId = useToString(() => route.params.roomId as string);
const url = computed(() => {
  return window.location.href;
});

const codeEditor = ref<HTMLTextAreaElement>();
const codeDisplay = ref<HTMLElement>();
const currentLanguage = ref("javascript");
const codeContent = ref();
const loading = ref(true);
const roomObject = ref();
const loadedContent = ref();
const loadedLanguage = ref();

const highlightCode = () => {
  if (!codeDisplay.value) return;

  try {
    const highlighted = hljs.highlight(codeContent.value, {
      language: currentLanguage.value,
    }).value;
    codeDisplay.value.innerHTML = highlighted;
  } catch (e) {
    codeDisplay.value.textContent = codeContent.value;
  }
};

const handleInput = () => {
  if (codeEditor.value) {
    codeContent.value = codeEditor.value.value;
    highlightCode();
  }
};

const handleScroll = () => {
  if (codeEditor.value && codeDisplay.value) {
    codeDisplay.value.scrollTop = codeEditor.value.scrollTop;
    codeDisplay.value.scrollLeft = codeEditor.value.scrollLeft;
  }
};

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === "Tab") {
    e.preventDefault();

    if (codeEditor.value) {
      const start = codeEditor.value.selectionStart;
      const end = codeEditor.value.selectionEnd;

      const newText =
        codeEditor.value.value.substring(0, start) +
        "  " +
        codeEditor.value.value.substring(end);

      codeEditor.value.value = newText;
      codeEditor.value.selectionStart = codeEditor.value.selectionEnd =
        start + 2;

      codeContent.value = newText;
      highlightCode();
    }
  }
};

const checkRoomId = () => {
  if (!roomId) window.location.replace("/");

  axios
    .get(`/api/room/${roomId.value}`)
    .then((data) => {
      roomObject.value = data.data;
      if ("Language" in data.data) {
        loadedLanguage.value = data.data.Language;
      }
      if ("Code" in data.data) {
        loadedContent.value = data.data.Code;
      }
    })
    .catch(() => {
      window.location.replace("/");
    })
    .finally(() => {
      loading.value = false;
    });
};

onMounted(() => {
  checkRoomId();

  nextTick(() => {
    highlightCode();

    if (codeEditor.value) {
      codeEditor.value.value = codeContent.value;
    }
  });
});

watch(currentLanguage, () => {
  highlightCode();
});

watch(codeContent, () => {
  highlightCode();
});

watch(loading, (value) => {
  if (!value) {
    nextTick(() => {
      currentLanguage.value = loadedLanguage.value ?? "";
      codeContent.value = loadedContent.value ?? "";
    });
  }
});
</script>

<template>
  <div v-if="loading" class="simple-loader">
    <div class="spinner"></div>
  </div>
  <div v-else class="container">
    <div class="header">
      <div class="title">
        <span class="title-icon">◉</span>
        <span>YaCode</span>
      </div>
      <div class="controls">
        <select class="language-select" v-model="currentLanguage">
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
        <div class="editor-title">
          {{
            currentLanguage === "javascript"
              ? "main.js"
              : currentLanguage === "python"
              ? "main.py"
              : currentLanguage === "java"
              ? "Main.java"
              : currentLanguage === "cpp"
              ? "main.cpp"
              : currentLanguage === "html"
              ? "index.html"
              : currentLanguage === "css"
              ? "style.css"
              : `main.${currentLanguage}`
          }}
        </div>
      </div>
      <div class="editor-container">
        <textarea
          ref="codeEditor"
          class="code-input"
          :value="codeContent"
          @input="handleInput"
          @scroll="handleScroll"
          @keydown="handleKeydown"
          spellcheck="false"
          autocapitalize="off"
          autocomplete="off"
          autocorrect="off"
        ></textarea>
        <pre ref="codeDisplay" class="code-display hljs"></pre>
      </div>
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

.editor-container {
  position: relative;
  width: 100%;
  height: calc(100% - 50px);
}

.code-input {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: transparent;
  color: transparent;
  font-family: "Monaco", "Menlo", "Ubuntu Mono", "Consolas", monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 20px;
  border: none;
  resize: none;
  outline: none;
  overflow: auto;
  caret-color: #fff;
  z-index: 2;
  tab-size: 2;
  white-space: pre;

  &::selection {
    background: rgba(101, 115, 255, 0.5);
  }
}

.code-display {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 20px;
  background: transparent;
  font-family: "Monaco", "Menlo", "Ubuntu Mono", "Consolas", monospace;
  font-size: 14px;
  line-height: 1.5;
  overflow: auto;
  pointer-events: none;
  z-index: 1;
  white-space: pre;
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
    height: 500px;
    margin: 10px;
  }

  .code-input,
  .code-display {
    font-size: 13px;
    padding: 15px;
  }
}

.simple-loader {
  position: fixed;
  inset: 0;
  background: rgba(255, 255, 255, 0.7);
  display: grid;
  place-items: center;
  z-index: 9999;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #e0e0e0;
  border-top-color: #3b82f6;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
