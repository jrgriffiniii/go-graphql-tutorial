<template>
  <div class="todo-list">
    <svg viewBox="0 0 0 0" style="position: absolute; z-index: -1; opacity: 0">
      <defs>
        <linearGradient
          id="boxGradient"
          gradientUnits="userSpaceOnUse"
          x1="0"
          y1="0"
          x2="25"
          y2="25"
        >
          <stop offset="0%" stop-color="#27FDC7" />
          <stop offset="100%" stop-color="#0FC0F5" />
        </linearGradient>

        <linearGradient id="lineGradient">
          <stop offset="0%" stop-color="#0FC0F5" />
          <stop offset="100%" stop-color="#27FDC7" />
        </linearGradient>

        <path
          id="todo__line"
          stroke="url(#lineGradient)"
          d="M21 12.3h168v0.1z"
        ></path>
        <path
          id="todo__box"
          stroke="url(#boxGradient)"
          d="M21 12.7v5c0 1.3-1 2.3-2.3 2.3H8.3C7 20 6 19 6 17.7V7.3C6 6 7 5 8.3 5h10.4C20 5 21 6 21 7.3v5.4"
        ></path>
        <path
          id="todo__check"
          stroke="url(#boxGradient)"
          d="M10 13l2 2 5-5"
        ></path>
        <circle id="todo__circle" cx="13.5" cy="12.5" r="10"></circle>
      </defs>
    </svg>

    <todo
      v-for="todo in todos"
      :key="todo.id"
      :todo="todo"
      @deleteTodo="updateStoreAfterTodoDeleted($event)"
    >
    </todo>

    <form class="todo-form">
      <input type="text" name="todo-text" id="todo-text" v-model="text" />
      <button type="button" @click="createTodo()">Add</button>
    </form>
  </div>
</template>

<script>
import { TODOS_QUERY, CREATE_TODO_MUTATION } from '@/constants/graphql'
import Todo from './Todo'

export default {
  name: 'TodoList',
  data () {
    return {
      todos: [],
      text: ''
    }
  },
  methods: {
    updateStoreAfterTodoDeleted ({ store, deleteTodo }) {
      const data = store.readQuery({
        query: TODOS_QUERY
      })

      const updated = data.todos.filter((e) => parseInt(e.id) === deleteTodo)
      data.todos = updated

      store.writeQuery({ query: TODOS_QUERY, data })
    },
    updateStoreAfterTodoCreated (store, createdTodo) {
      const data = store.readQuery({
        query: TODOS_QUERY
      })

      data.todos.push(createdTodo)
      store.writeQuery({ query: TODOS_QUERY, data })
    },
    createTodo () {
      const { text } = this.$data

      this.$apollo.mutate({
        mutation: CREATE_TODO_MUTATION,
        variables: {
          input: {
            text: text
          }
        },
        update: (store, { data: { createTodo } }) => {
          this.updateStoreAfterTodoCreated(store, createTodo)
        }
      })
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

<style scoped>
.todo-list {
  background: #fff;
  font-size: 20px;
  max-width: 22em;
  margin: auto;
  padding: 0.5em 1em;
  box-shadow: 0 5px 30px rgba(0, 0, 0, 0.2);
}

.todo-form {
  display: block;
  position: relative;
  padding: 1em 1em 1em 16%;
  margin: 0 auto;
  cursor: pointer;
  border-bottom: none;
}

.todo-form .todo-input {
  color: saturate(#1b4a4e, 15%);
  transition: all 0.8s/2 linear 0.8s/2;
  font-size: 20px;
}

form {
  margin-top: 3rem;
  display: flex;
  flex-wrap: wrap;
}
form label {
  min-width: 100%;
  margin-bottom: 0.5rem;
  font-size: 1.3rem;
}
form input {
  flex-grow: 1;
  border: none;
  background: #f7f1f1;
  padding: 0 1.5em;
  font-size: initial;
}
form button {
  padding: 0 1.3rem;
  border: none;
  background: #1b4a4e;
  color: white;
  text-transform: uppercase;
  font-weight: bold;
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-left: 5px;
  cursor: pointer;
  transition: background 0.2s ease-out;
}

form button:hover {
  background: #27fdc7;
}

form input,
form button {
  font-family: "Quicksand", sans-serif;
  height: 3rem;
}
</style>
