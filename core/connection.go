package core

import (
	"github.com/vergieet/drago-tail/controller"
)


func Tes() {
	r= make(map[string]Routes)
	add(Routes{GET,"/halo",controller.GetAllUser})
}

