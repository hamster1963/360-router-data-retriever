package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hamster1963/360-router-data-retriever/configs"
	"github.com/hamster1963/360-router-data-retriever/router"
)

func RouterSimple() {
	var routerMain router.SRouterController
	configs.RouterIP = "router.xinyu.today:580"
	myRouter := router.Router{
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
	_, err = routerMain.GetRouterInfo()
	if err != nil {
		g.Dump(err)
		return
	}
	_, err = routerMain.GetRouterSpeed()
	if err != nil {
		g.Dump(err)
		return
	}
}
