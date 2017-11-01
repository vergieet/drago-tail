package config

import (
	"fmt"
	"myfw/controller"
)

const POST  = "post"
const GET  = "get"
var r []Routes
var rm map[string]Routes
func main() {
	//Routes{"/halo",controller.GetAllUser};
}
func Mains() {
	rm = make(map[string]Routes)
	add(Routes{GET,"/halo/",controller.GetAllUser})
	add(Routes{GET,"/halo",controller.GetAllUser})
}
func GetRoute(url string) string {
	Mains()
	return rm[url].run()
}

func add(routes Routes)  {
	rm[routes.url] = routes
	r = append(r,routes)
}

type Routes struct {
	method string
	url string
	run
}

type run func() string;

func (r *Routes) POST() {
	fmt.Println("Hi, my name is")
}
func (r *Routes) GET() {
	fmt.Print(r.run())
	fmt.Println("GET")
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
