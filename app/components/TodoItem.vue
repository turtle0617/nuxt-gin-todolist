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
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { Todo, TodoStatus } from '../types'

@Component
export default class TodoItem extends Vue {
  @Prop({ type: Object, required: true, default: null })
  public value!: Todo | null

  public TodoStatus = TodoStatus

  get todoKey() {
    return `todo-${this.value?.id}`
  }

  updateValue(val: any) {
    this.$emit('update', { title: val.target.value })
  }

  updateStatus(val: any) {
    this.$emit('update', {
      status: val.target.checked ? TodoStatus.COMPLETED : TodoStatus.IDLE,
    })
  }
}
</script>
