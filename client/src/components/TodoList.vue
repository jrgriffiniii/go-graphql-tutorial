<template>
  <div>
    <h4 v-if="todos.length == 0">Loading...</h4>
    <todo
      v-for="todo in todos"
      :key="todo.id"
      :todo="todo">
    </todo>
  </div>
</template>

<script>
import { TODOS_QUERY } from '@/constants/graphql'
import Todo from './Todo'

export default {
  name: 'TodoList',
  data () {
    return {
      todos: []
    }
  },
  components: {
    Todo
  },
  apollo: {
    todos: {
      query: TODOS_QUERY
    }
  }
}
</script>
