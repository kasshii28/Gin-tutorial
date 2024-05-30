package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int64
	Content string
	ContentTag string
	Done bool
	UserID int64
	CreatedAt time.Time
	UpdatedAt time.Time
	ReservedAt time.Time
}

// Todoの作成
func (u *User) CreateTodo(content string, content_tag string, done bool, reserved_at time.Time) (err error) {
	cmd := `insert into todos(
		content,
		content_tag,
		done,
		user_id,
		created_at,
		updated_at,
		reserved_at) values (? ? ? ? ? ? ?)`
	
	_, err = Db.Exec(
		cmd, 
		content,
		u.ID, // userインスタンスのI
		content_tag,
		done, 
		time.Now(),
		time.Now(),
		reserved_at,
	)

	if err != nil {
		log.Fatalln("failed to create Todo",err)
	}

	return err
}

// Todoの作成
func GetTodo(id int) (todo Todo, err error) {
	cmd := `
	select id, 
	content, 
	content_tag, 
	done,
	user_id, 
	created_at, 
	updated_at, 
	reserved_at 
	from todos where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.ContentTag,
		&todo.Done,
		&todo.UserID,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.ReservedAt)

	return todo, err
}

// Todoの一覧取得
func GetTodos() (todos []Todo, err error) {
	cmd := `select
	id,
	content,
	content_tag,
	done,
	user_id,
	created_at,
	updated_at,
	reserved_at
	from todos`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln("query error",err)
	}
	
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.ContentTag,
			&todo.Done,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.ReservedAt)
		if err != nil {
			log.Fatalln("failed to GetTodos",err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// 特定のユーザーのtodo取得
func (u *User) GetTodosByUser() (todos []Todo, err error){
	cmd := `
	select
	id,
	content,
	content_tag,
	done,
	user_id,
	created_at,
	updated_at,
	reserved_at
	from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln("query error",err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.ContentTag,
			&todo.Done,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.ReservedAt)
		if err != nil {
			log.Fatalln("failed to GetTodosByUser",err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// Todoの更新
func (t *Todo) UpdateTodo() error {
	cmd := `
	update todos set 
	content = ?,
	content_tag = ?,
	user_id = ? where id = ?`

	_, err := Db.Exec(cmd, t.Content, t.ContentTag, t.UserID, t.ID)
	if err != nil {
		log.Fatalln("failed to update Todo",err)
	}
	return err
}

// Todoの削除
func (t *Todo) DeleteTodo() error {
	cmd :=`delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln("failed to delete Todo", err)
	}
	return err
}