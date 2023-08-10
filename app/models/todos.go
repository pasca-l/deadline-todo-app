package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserID int
	Deadline time.Time
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string, deadline time.Time) (err error) {
	cmd := `INSERT INTO todos (content, user_id, created_at, deadline)
			VALUES (?, ?, ?, ?)`

	_, err = Db.Exec(
		cmd, content, u.ID, time.Now(), deadline,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT id, content, user_id, deadline, created_at
			FROM todos WHERE id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.Deadline,
		&todo.CreatedAt,
	)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, deadline, created_at
			FROM todos`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.Deadline,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, deadline, created_at
			FROM todos WHERE user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.Deadline,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = ?, deadline = ?, user_id = ?
			WHERE id = ?`

	_, err = Db.Exec(cmd, t.Content, t.Deadline, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `DELETE FROM todos WHERE id = ?`

	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
