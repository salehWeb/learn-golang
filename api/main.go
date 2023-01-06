package main

import (
	"fmt"
)

type ToDo struct {
	Id      uint32 `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
}

type Db interface {
	InsertTodo(content string)
	GetToDos() []ToDo
}

type db struct {
	Len   uint32
	ToDos []ToDo
}

func NewDb() Db {
	return &db{
		Len:   0,
		ToDos: []ToDo{},
	}
}

func (db *db) InsertTodo(content string) {

	var newTodo ToDo = ToDo{
		Id:      db.Len + 1,
		Content: content,
		IsDone:  false,
	}

	db.ToDos = append(db.ToDos, newTodo)
	db.Len += 1
}

func (db *db) GetToDos() []ToDo {
	return db.ToDos
}

func main() {
	var DB Db = NewDb()

	DB.InsertTodo("Hello World 1")
	DB.InsertTodo("Hello World 2")
	DB.InsertTodo("Hello World 3")

	toDos := DB.GetToDos()

	for _, v := range toDos {
		fmt.Printf("Todos Id: %d Content is %s and IsDone %t\n\n", v.Id, v.Content, v.IsDone)
	}
}
