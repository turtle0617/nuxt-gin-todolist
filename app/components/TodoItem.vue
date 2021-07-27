<template>
  <div class="py-2 flex items-center">
    <input
      type="checkbox"
      :checked="value.status === TodoStatus.COMPLETED"
      @change="updateStatus"
    />
    <input class="pl-1 flex-1" :value="value.title" @blur="updateValue" />
    <button class="px-2" @click="$emit('delete', value.id)">x</button>
  </div>
</template>

<script lang="ts">
import { PropOptions } from 'vue'
import { Todo, TodoStatus } from '../types'
export default {
  props: {
    value: {
      type: Object,
      require: true,
      default: null,
    } as PropOptions<Todo | null>,
  },
  data() {
    return {
      TodoStatus,
    }
  },
  computed: {
    todoKey() {
      return `todo-${this.value.index}`
    },
  },
  methods: {
    updateValue(val) {
      this.$emit('update', { title: val.target.value })
    },
    updateStatus(val) {
      this.$emit('update', {
        status: val.target.checked ? TodoStatus.COMPLETED : TodoStatus.IDLE,
      })
    },
  },
}
</script>
