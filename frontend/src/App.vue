<script lang="ts" setup>
import TodoList from "./components/TodoList.vue";
import NewTodo from "./components/NewTodo.vue";
import TodoController from "./lib/TodoController";
import * as M from "@wailsjs/models";
import { Quit } from "../wailsjs/runtime/runtime";
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
    <div id="app-view">
      <div id="titlebar">
        <span @click="Quit()" class="btn">✕</span>
      </div>
      <div id="main-view">
        <h1><em>Super Todo App</em></h1>
        <NewTodo :controller="controller" />
        <hr />
        <TodoList :controller="controller" :tasks="taskss" />
      </div>
    </div>
  </Suspense>
</template>

<style>
* {
  /*color: #000;*/
  font-family: "Advent Pro";
}

input,
button {
  border: 2px solid #0005;
  background: #fff;
  /*height: 24px;*/
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

#app-view {
  display: flex;
  flex-direction: column;
}
#titlebar {
  width: 100%;
  text-align: right;
  translate: -10px 10px;
}
</style>
