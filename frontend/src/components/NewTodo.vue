<template>
  <div class="new-todo">
    <h2>Add new todo</h2>
    <hr />
    <span class="input-group">
      <input v-model="input" />
      <button @click="submit()">Add</button>
    </span>
  </div>
</template>
<script lang="ts" setup>
import { ITodoController } from "../lib/TodoController";
import { defineModel } from "vue";

const props = defineProps<{
  controller: ITodoController;
}>();

const input = defineModel<string>();

const submit = () => {
  if (input.value == "") {
    return;
  }
  props.controller.addTask(input.value ?? "");
  input.value = "";
};
</script>
<style scoped>
.new-todo {
  width: 100%;
}

.input-group {
  display: flex;
  flex-direction: row;
  width: 100%;
}
</style>
