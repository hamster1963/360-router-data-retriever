package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"router_remake/configs"
	"router_remake/internal"
)

func main() {
	var routerMain internal.RouterController
	myRouter := internal.Router{
		Address:  configs.RouterAddress,
		Password: configs.RouterPassword,
	}
	routerMain = &myRouter

	err := routerMain.GetRandomString()
	if err != nil {
		g.Dump(err)
		return
	}
	err = routerMain.GenerateAesString()
	if err != nil {
		g.Dump(err)
		return
	}
	err = routerMain.Login()
	if err != nil {
		g.Dump(err)
		return
	}
	err = routerMain.GetRouterInfo()
	if err != nil {
		g.Dump(err)
		return
	}
}
