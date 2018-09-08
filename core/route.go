package core

import (
	"github.com/vergieet/drago-lib/response"
)

const (
	POST  = "post"
	GET  = "get"
	)

var r map[string]Routes

func GetRoute(url string) string {

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


type run func() string;

func (r *Routes) Get(){

}