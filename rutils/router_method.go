package rutils

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hamster1963/360-router-data-retriever/rconfig"
)

type RouterMethod interface {
	GetRouterInfo() (g.Map, error)
	GetRouterSpeed() (g.Map, error)
	GetDeviceList() (g.Map, error)
}

func (r *Router) GetRouterSpeed() (speedData g.Map, err error) {
	if loginState, err := r.CheckLogin(); err != nil || loginState == false {
		err := r.Login()
		if err != nil {
			return nil, err
		}
	}
	apiUrl := r.RouterAddress + rconfig.RouterSpeedUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		return nil, err
	}
	resData := res.ReadAllString()
	speedData = gconv.Map(resData)
	return speedData, nil
}

func (r *Router) GetRouterInfo() (infoData g.Map, err error) {
	if loginState, err := r.CheckLogin(); err != nil || loginState == false {
		err := r.Login()
		if err != nil {
			return nil, err
		}
	}
	apiUrl := r.RouterAddress + rconfig.RouterInfoUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		return nil, err
	}
	resData := res.ReadAllString()
	infoData = gconv.Map(resData)
	return infoData, nil
}

func (r *Router) GetDeviceList() (DeviceListData g.Map, err error) {
	if loginState, err := r.CheckLogin(); err != nil || loginState == false {
		err := r.Login()
		if err != nil {
			return nil, err
		}
	}
	apiUrl := r.RouterAddress + rconfig.RouterDeviceInfoUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		return nil, err
	}
	resData := res.ReadAllString()
	DeviceListData = gjson.New(resData).Get("data.0").Map()
	return DeviceListData, nil
}
