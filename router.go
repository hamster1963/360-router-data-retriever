package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hamster1963/360-router-data-retriever/configs"
	"github.com/hamster1963/360-router-data-retriever/router"
)

func RouterSimple() {
	var routerMain router.SRouterController
	routerConfig := &configs.RouterConfig{
		RouterIP:       "router.xinyu.today:580",
		RouterAddress:  "http://router.xinyu.today:580",
		RouterPassword: "deny1963",
	}
	myRouter := router.Router{
		RouterConfig: routerConfig,
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
