<template>
  <div class="todo">{{todo.text}}</div>
</template>

<script>
export default {
  name: 'Todo',
  props: ['todo']
}
</script>
