<template>
  <div class="flex flex-col h-full">
    <div class="flex">
      <input
        v-model="newTodoLabel"
        class="flex-auto border-solid rounded-md border-2 border-light-blue-500"
        @keyup.enter="addTodo"
      />
      <button class="px-2" @click="addTodo">+</button>
    </div>
    <template v-if="todos.length">
      <TodoItem
        v-for="(todo, index) of todos"
        :key="index"
        :value="{ ...todo, index }"
      />
    </template>
    <div v-else class="py-8 text-center	" >Add some in todo</div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data() {
    return {
      newTodoLabel: '',
      todos: [],
    }
  },
  methods: {
    addTodo() {
      this.todos.push({
        label: this.newTodoLabel,
        finished: false,
      })
      this.newTodoLabel = ''
    },
  },
})
</script>
