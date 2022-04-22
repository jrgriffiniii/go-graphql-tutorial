import Vue from 'vue'
import Todo from '@/components/Todo'

describe('Todo.vue', () => {
  it('should render the text description for the todo item', () => {
    const Constructor = Vue.extend(Todo)
    const constructed = new Constructor({

      propsData: {
        todo: {
          text: 'todo item 1'
        }
      }
    })

    const vm = constructed.$mount()

    expect(vm.$el.textContent).toEqual(expect.stringContaining('todo item 1'))
  })
})
