package controller

import "fmt"

type User interface {
	GetAllUser() string
}
type Userx struct {
	name string
}

func GetFunc() func(){
	return func(){
		fmt.Print("ha");
	}
}

func  GetAllUser() string {
	return "[{\"username\":\"HALLLOO\"},{\"username\":\"HAhai\"}]"
}

func (t Userx) get() string {
	return t.name
}