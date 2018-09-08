package controller

import (
	"fmt"
	"github.com/vergieet/drago-lib/response"
)

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
func (u Userx)getUser(){

}

func  GetAllUser() string {
	//var Users []model.User
	//Users = append(Users,user.Find(1))
	//urlsJson, _ := json.Marshal(Users)
	//return response.Success(string(urlsJson))
	return response.Success("{}")
	}

func (t Userx) get() string {
	return t.name
}
func main(){
	var user Userx;
	user.getUser();

}