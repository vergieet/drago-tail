package config

import (
	"github.com/vergieet/drago-lib/response"
	"myfw/controller"
)

const POST  = "post"
const GET  = "get"

var r map[string]Routes

func Mains() {
	r= make(map[string]Routes)
	add(Routes{GET,"/halo",controller.GetAllUser})
}
func GetRoute(url string) string {
	Mains()
	if val, ok := r[url]; ok {
		return val.run()
	}else{
		last  := url[len(url)-1:]
		if(last == "/"){
			url = url[0:len(url)-1]
		}else {
			url = url + "/"
		}
		if val, ok := r[url]; ok {
			return val.run()
		}else {
			return response.Fail("404 Not Found")
		}
	}

}

func add(routes Routes)  {
	r[routes.url] = routes
}

type Routes struct {
	method string
	url string
	run
}

type run func() string;

