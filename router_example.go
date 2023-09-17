package router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hamster1963/360-router-data-retriever/rconfig"
	"github.com/hamster1963/360-router-data-retriever/rutils"
	"time"
)

func TestRouterExample() {
	routerConfig := &rconfig.RouterConfig{
		RouterIP:       "router.example.today",
		RouterAddress:  "https://router.example.today",
		RouterPassword: "example",
	}

	myRouter := rutils.NewRouter().InitRouter(routerConfig)
	var routerMain rutils.SRouterController = myRouter
	err := routerMain.GetRandomString()
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}
	err = routerMain.GenerateAesString()
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}
	err = routerMain.Login()
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}
	_, err = routerMain.GetRouterInfo()
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}

	deviceList, err := routerMain.GetDeviceList()
	if err != nil {
		glog.Warning(context.Background(), err)
		return
	}
	g.Dump(deviceList)

	// 获取十次网速
	for i := 0; i < 10; i++ {
		info, err := routerMain.GetRouterSpeed()
		if err != nil {
			glog.Warning(context.Background(), err)
			return
		}
		g.Dump(info)
		time.Sleep(1 * time.Second)
	}

}
