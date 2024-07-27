<script lang="ts" setup>
import type { Todo } from "../model/todo";
import { ITodoController } from "../lib/TodoController";
const props = defineProps<{
  controller: ITodoController;
  done: boolean;
  content: string;
  id: number;
}>();
</script>
<template>
  <div class="card">
    <span>{{ content }}</span>
    <div>
      <span @click="controller.tickTask(id, !done)"
        ><span v-if="done"> ✅ </span><span v-else>⬜️</span></span
      >
      <span @click="controller.removeTask(id)"> 🗑️ </span>
    </div>
  </div>
</template>
<style scoped>
.card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;

  color: #000;
  border: 2px solid #0003;
  border-radius: 5px;
  max-width: 100%;
  height: 50px;
  background: #fff;
}
</style>
