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
        v-for="todo of todos"
        :key="todo.id"
        :value="todo"
      />
    </template>
    <div v-else class="py-8 text-center">Add some in todo</div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  async asyncData({ $axios }) {
    const res = await $axios.$get('http://localhost:8080/todos')
    return { todos: res.data || [] }
  },
  data(): {
    newTodoLabel: string
    todos: Record<'id' | 'title' | 'status', string>[]
  } {
    return {
      todos: [],
      newTodoLabel: '',
    }
  },
  methods: {
    addTodo() {
      this.newTodoLabel = ''
    },
  },
})
</script>
