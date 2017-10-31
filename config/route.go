package config

import (
	"fmt"
	"myfw/controller"
)
func main() {
	Routes{"/halo",controller.UserController{}.GetUsername()}.POST();
}

type Routes struct {
	url string
	data string
}

func (r *Routes) POST() {
	fmt.Println("Hi, my name is")
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
