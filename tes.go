package main

import (
	"fmt"
	//"io/ioutil"
	//"encoding/json"
	"net/http"
	"github.com/vergieet/drago-tail/config"

)

func handler(w http.ResponseWriter, r *http.Request) {
	//var data map[string]map[string]string
	//plan, _ := ioutil.ReadFile(	"./config/route.json")
	//json.Unmarshal(plan, &data)
	var ra string
	ra = config.GetRoute(r.URL.RequestURI());
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,"%s",ra)
	//fmt.Fprintf(w, "Hi the %s %T!", data,data)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3016", nil)
	fmt.Print("Starting server at port :3016")
}
