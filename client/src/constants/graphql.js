import gql from 'graphql-tag'

export const TODOS_QUERY = gql`
  query {
    todos {
      id
      text,
      done
    }
  }
`

export const CREATE_TODO_MUTATION = gql`
  mutation CreateTodo($input: NewTodo!) {
    createTodo(
      input: $input,
    ) {
      id
      text
    }
  }
`
