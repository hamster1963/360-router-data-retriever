package internal

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	"hamster1963/360-router-data-retriever/configs"
)

type RouterMethod interface {
	GetRouterInfo() (g.Map, error)
	GetRouterSpeed() (g.Map, error)
}

func (r *Router) GetRouterSpeed() (speedData g.Map, err error) {
	if r.state == false {
		err = errors.New("please login first")
		return nil, err
	}
	apiUrl := r.Address + configs.RouterSpeedUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			g.Dump(err)
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
	if r.state == false {
		err = errors.New("please login first")
		return
	}
	apiUrl := r.Address + configs.RouterInfoUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			g.Dump(err)
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
