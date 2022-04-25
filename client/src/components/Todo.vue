<template>
  <div class="todo">
      <input class="todo__state" type="checkbox" />

      <div class="todo__text">{{todo.text}}</div>

    <button @click="deleteTodo()" class="todo__delete" type="button" aria-label="Delete" title="Delete">
      <i aria-hidden="true" class="material-icons">delete</i>
    </button>
  </div>
</template>

<script>
import { DELETE_TODO_MUTATION } from '@/constants/graphql'

export default {
  name: 'Todo',
  props: ['todo'],
  methods: {
    deleteTodo () {
      this.$apollo.mutate({
        mutation: DELETE_TODO_MUTATION,
        variables: {
          id: this.todo.id
        },
        update: (store, { data: { deleteTodo } }) => {
          this.$emit('deleteTodo', { store, deleteTodo })
        }
      })
    }
  }
}
</script>

<style scoped>

.todo {
  display: flex;
  position: relative;
  padding: 1em 1em 1em 16%;
  margin: 0 auto;
  border-bottom: solid 1px #ddd;

  align-items: center;
  justify-content: space-between;
}

.todo:last-child {
  border-bottom: none;
}

.todo__state {

}

.todo__text {
  color: saturate(#1B4A4E,15%);
  transition: all 0.8s/2 linear 0.8s/2;
  text-align: left;

  width: 100%;
  padding-right: 1.2rem;
  padding-left: 1.2rem;
}

.todo__icon {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: auto;
  margin: auto;

  fill: none;
  stroke: #27FDC7;
  stroke-width: 2;
  stroke-linejoin: round;
  stroke-linecap: round;
}

.todo__line,
.todo__box,
.todo__check {
  transition: stroke-dashoffset 0.8s cubic-bezier(.9,.0,.5,1);
}

.todo__circle {
  stroke: #27FDC7;
  stroke-dasharray: 1 6;
  stroke-width: 0;

  transform-origin: 13.5px 12.5px;
  transform: scale(0.4) rotate(0deg);
  animation: none 0.8s linear;

  @keyframes explode {
    30% {
      stroke-width: 3;
      stroke-opacity: 1;
      transform: scale(0.8) rotate(40deg);
    }
    100% {
      stroke-width: 0;
      stroke-opacity: 0;
      transform: scale(1.1) rotate(60deg);
    }
  }
}

.todo__box {
  stroke-dasharray: 56.1053, 56.1053; stroke-dashoffset: 0;
  transition-delay: 0.8s * 0.2;
}

.todo__check {
  stroke: #27FDC7;
  stroke-dasharray: 9.8995, 9.8995; stroke-dashoffset: 9.8995;
  transition-duration: 0.8s * 0.4;
}

.todo__line {
  stroke-dasharray: 168, 1684;
  stroke-dashoffset: 168;
}

.todo__circle {
  animation-delay: 0.8s * 0.7;
  animation-duration: 0.8s * 0.7;
}

.todo__state:checked ~ .todo__text {
  color: #5EBEC1;
  opacity: 0.6;
  text-decoration: line-through;
  transition-delay: 0.8s * 0.2;
}

.todo__state:checked ~ .todo__icon .todo__box { stroke-dashoffset: 57.1053; transition-delay: 0s; }
.todo__state:checked ~ .todo__icon .todo__line { stroke-dashoffset: -8; }
.todo__state:checked ~ .todo__icon .todo__check { stroke-dashoffset: 0; transition-delay: 0.8s * 0.6; }
.todo__state:checked ~ .todo__icon .todo__circle { animation-name: explode; }

.todo__delete {
  display: inline-block;

	border:none;
	background:none;
	-webkit-appearance:none;
	cursor:pointer;
  color: #1B4A4E;
}
</style>
