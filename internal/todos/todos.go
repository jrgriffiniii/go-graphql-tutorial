package todos

import (
	"log"

	database "github.com/jrgriffiniii/go-graphql-tutorial/internal/pkg/db/mysql"
	"github.com/jrgriffiniii/go-graphql-tutorial/internal/users"
)

type Todo struct {
	ID   string
	Text string
	Done bool
	User *users.User
}

func (todo Todo) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Todos(Text,Done) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(todo.Text, todo.Done)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("Row inserted!")
	return id
}

func GetAll() []Todo {
	stmt, err := database.Db.Prepare("SELECT id, text, done FROM Todos")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Text, &todo.Done)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todos
}
