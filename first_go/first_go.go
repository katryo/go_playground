// first_go.go
package main

import (
	"fmt"
)

type Task struct {
	ID     int
	Detail string
	done   bool
	*User
}

type Stringer interface {
	String() string
}

type Any interface {
}

func Do(e Any) {
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func Finish(task *Task) {
	task.done = true
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d %s", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func Print(value interface{}) {
	s, ok := value.(string)
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}

func main() {
	Print("abc")
	Print(10)
}
