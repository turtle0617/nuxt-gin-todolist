<template>
  <div class="py-2 flex items-center">
    <input
      :id="todoKey"
      type="checkbox"
      :checked="value.status === 'completed'"
    />
    <label :for="todoKey" class="pl-1 flex-1">{{ value.title }}</label>
    <button class="px-2" @click="$emit('delete', value.id)">x</button>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: Object,
      require: true,
      default: () => {},
    },
  },
  computed: {
    todoKey() {
      return `todo-${this.value.index}`
    },
  },
}
</script>
