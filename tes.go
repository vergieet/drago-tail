package main

import (
	"fmt"
	//"io/ioutil"
	//"encoding/json"
	"net/http"
	"github.com/vergieet/drago-tail/core"
	"go/importer"

)

func handler(w http.ResponseWriter, r *http.Request) {
	//var data map[string]map[string]string
	//plan, _ := ioutil.ReadFile(	"./config/route.json")
	//json.Unmarshal(plan, &data)
	var ra string
	ra = core.GetRoute(r.URL.RequestURI());
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,"%s",ra)
	//fmt.Fprintf(w, "Hi the %s %T!", data,data)
}
//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe(":3016", nil)
//	fmt.Print("Starting server at port :3016")
//}

type tes struct {

}
func main() {

	//	//x := struct{Foo string; Bar int }{"foo", 2}
	//

	//	reflectValue := reflect.ValueOf(&controller.UserController{})
	//	controllerName := reflect.Indirect(reflectValue).Type().NumMethod()
	//
	//	//
	//	//values := make([]interface{}, v.NumField())
	//	//
	//	//for i := 0; i < v.NumField(); i++ {
	//	//	//values[i] = v.Field(i).Interface()
	//	//}
	//
	//	fmt.Println(reflectType,"\n",reflectValue,"\n",controllerName)
	//}
	pkg, err := importer.Default().Import("github.com/vergieet/drago-tail/controller")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	for _, declName := range pkg.Scope().Names() {
		//reflect.ValueOf(reflect.)
		fmt.Println(declName)
	}
}
