package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"myfw/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var data map[string]map[string]string
	plan, _ := ioutil.ReadFile(	"./config/route.json")
	json.Unmarshal(plan, &data)


	fmt.Fprintf(w,"%s",controller.GetUsername())
	//fmt.Fprintf(w, "Hi the %s %T!", data,data)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3016", nil)
	fmt.Print("Starting server at port :3016")
}
