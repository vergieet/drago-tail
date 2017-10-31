package controller

import "fmt"

func main() {
}

func GetFunc() func(){
	return func(){
		fmt.Print("ha");
	}
}

func GetAllUser() string {
	return "a"
}
