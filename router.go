package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hamster1963/360-router-data-retriever/rconfig"
	"github.com/hamster1963/360-router-data-retriever/rutils"
	"time"
)

func RouterSimple() {
	routerConfig := &rconfig.RouterConfig{
		RouterIP:       "router.xinyu.today:580",
		RouterAddress:  "http://router.xinyu.today:580",
		RouterPassword: "deny1963",
	}

	myRouter := rutils.NewRouter().InitRouter(routerConfig)
	var routerMain rutils.SRouterController = myRouter
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

	// 获取十次网速
	for i := 0; i < 10; i++ {
		info, err := routerMain.GetRouterSpeed()
		if err != nil {
			g.Dump(err)
			return
		}
		g.Dump(info)
		time.Sleep(1 * time.Second)
	}

}
