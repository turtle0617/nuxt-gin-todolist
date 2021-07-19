<template>
  <div class="py-2 flex items-center">
    <input
      :id="todoKey"
      type="checkbox"
      :checked="value.status === 'completed'"
    />
    <label :for="todoKey" class="pl-1">{{ value.title }}</label>
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
