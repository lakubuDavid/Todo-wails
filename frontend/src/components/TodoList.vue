<template>
  <div class="todo-list">
    <Todo
      v-for="todo in tasks"
      :content="todo.content"
      :done="todo.done"
      :id="todo.id"
      :key="todo.id"
      :controller="controller"
    />
  </div>
</template>
<script lang="ts" setup>
import Todo from "./Todo.vue";
import { ITodoController } from "../lib/TodoController";
import * as M from "@wailsjs/models";
const props = defineProps<{
  tasks: M.db.Todo[];
  controller: ITodoController;
}>();
</script>
<style scoped>
.todo-list {
  width: 100%;
  overflow-y: scroll;
  max-height: 50vh;
  gap: 10px;
  display: flex;
  flex-direction: column;
}
</style>
