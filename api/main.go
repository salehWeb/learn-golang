package main

import (
	"fmt"
)

type ToDo struct {
	Id      uint8  `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"isDone"`
}

type Db interface {
	InsertTodo(content string)
	GetToDos() map[uint8]ToDo
	DeleteItemByID(id uint8) bool
	UpdateItemByID(id uint8, content string, isDone bool) bool
}

type db struct {
	Len   uint8          `json:"length"`
	ToDos map[uint8]ToDo `json:"toDos"`
}

func NewDb() Db {

	return &db{
		Len:   0,
		ToDos: make(map[uint8]ToDo),
	}
}

func (db *db) GetToDos() map[uint8]ToDo {
	return db.ToDos
}

func (db *db) InsertTodo(content string) {

	if db.Len == 254 {
		return
	}

	var id uint8 = db.Len + 1
	var newTodo ToDo = ToDo{
		Id:      id,
		Content: content,
		IsDone:  false,
	}

	db.ToDos[id] = newTodo
	db.Len += 1
}

func (db *db) DeleteItemByID(id uint8) bool {
	_, ok := db.ToDos[id]

	if ok {
		delete(db.ToDos, id)
		return true
	} else {
		return false
	}

}

func (db *db) UpdateItemByID(id uint8, content string, isDone bool) bool {

	if item, ok := db.ToDos[id]; ok {

		item.Content = content
		item.IsDone = isDone

		db.ToDos[id] = item
		return true
	} else {
		return false
	}

}

func main() {
	var DB Db = NewDb()

	DB.InsertTodo("Hello World 1")
	DB.InsertTodo("Hello World 2")
	DB.InsertTodo("Hello World 3")

	var ToDos map[uint8]ToDo
	ToDos = DB.GetToDos()
	ReadTodos(&ToDos)

	var isSuccess1 bool = DB.DeleteItemByID(1)

	if isSuccess1 {
		fmt.Println("Success Deleting Todo")
		fmt.Println("Todos After Deleting")

		ToDos = DB.GetToDos()

		ReadTodos(&ToDos)

	} else {
		fmt.Println("Failed to Delete Todo")
	}

	var isSuccess2 bool = DB.UpdateItemByID(3, "World Hello Updated", true)

	if isSuccess2 {
		fmt.Println("Success Updating Todo")
		fmt.Println("Todos After Updating")

		ToDos = DB.GetToDos()

		ReadTodos(&ToDos)

	} else {
		fmt.Println("Failed to Update Todo")
	}

}

func ReadTodos(toDos *map[uint8]ToDo) {
	for _, v := range *toDos {
		fmt.Printf("Todos Id: %d Content is %s and IsDone %t\n\n", v.Id, v.Content, v.IsDone)
	}
}
