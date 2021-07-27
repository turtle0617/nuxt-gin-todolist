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
        @update="(payload) => updateTodo(todo.id, payload)"
        @delete="deleteTodo"
      />
    </template>
    <div v-else class="py-8 text-center">Add some in todo</div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Todo } from '../types'

export default Vue.extend({
  data(): {
    newTodoLabel: string
    todos: Todo[]
  } {
    return {
      todos: [
        {
          id: 'a',
          title: 'b',
          status: 'idle',
        },
      ],
      newTodoLabel: '',
    }
  },
  async fetch() {
    const res = await this.$axios.$get('http://localhost:8080/todos')
    this.todos = res.data || []
  },
  methods: {
    async addTodo() {
      try {
        if (!this.newTodoLabel) return
        await this.$axios.post('http://localhost:8080/todos', {
          title: this.newTodoLabel,
        })
        this.newTodoLabel = ''
        this.$fetch()
      } catch (error) {
        console.error(error)
      }
    },
    async updateTodo(id: string, payload: Partial<Todo>) {
      await this.$axios.$patch(`http://localhost:8080/todos/${id}`, payload)
      this.$fetch()
    },
    async deleteTodo(id: string) {
      await this.$axios.$delete(`http://localhost:8080/todos/${id}`)
      this.$fetch()
    },
  },
})
</script>
