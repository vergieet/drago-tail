package main

import (
	"fmt"
	//"io/ioutil"
	//"encoding/json"
	"net/http"
	"myfw/config"
	//"encoding/json"
	"myfw/model"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//var data map[string]map[string]string
	//plan, _ := ioutil.ReadFile(	"./config/route.json")
	//json.Unmarshal(plan, &data)
	var ra string
	ra = config.GetRoute(r.URL.RequestURI());
	w.Header().Set("Content-Type", "application/json")
	var sa model.User
	sa.Name = "ha"
	sa.Email = "hi"
	fmt.Fprintf(w,"%s %s",ra,sa.All())
	//fmt.Fprintf(w, "Hi the %s %T!", data,data)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3016", nil)
	fmt.Print("Starting server at port :3016")
}
