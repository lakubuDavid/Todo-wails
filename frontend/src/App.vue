<script lang="ts" setup>
import TodoList from "./components/TodoList.vue";
import NewTodo from "./components/NewTodo.vue";
import TodoController from "./lib/TodoController";
import * as M from "@wailsjs/models";
import { onMounted, ref, watch } from "vue";

const controller = TodoController();

const taskss = ref<M.db.Todo[]>([]);

onMounted(async () => {
  watch(
    controller.asChanged,
    async (val, prev) => {
      console.log("changes");
      taskss.value = await controller.list();
      console.log(taskss);
      controller.validate();
    },
    {
      immediate: true,
      deep: true,
    },
  );
});
</script>

<template>
  <Suspense>
    <div id="main-view">
      <h1>Todo App</h1>
      <NewTodo :controller="controller" />
      <hr />
      <TodoList :controller="controller" :tasks="taskss" />
    </div>
  </Suspense>
</template>

<style>
* {
  color: #000;
  font-family: "Advent Pro";
}

body {
  background: #fff;
}

input,
button {
  border: 2px solid #0005;
  background: #fff;
  //height: 24px;
  border-radius: 5px;
  padding: 10px 15px;
}
input {
  width: 100%;
}

button {
  margin: 0px 10px;
}
button:hover {
  background: #eee;
}

#main-view {
  display: flex;
  text-alight: left;
  flex-direction: column;
  justify-content: center;
  align-items: start;
  padding: 50px;
}
</style>
