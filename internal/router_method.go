package internal

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"router_remake/configs"
)

type RouterMethod interface {
	GetRouterInfo() error
}

func (r *Router) GetRouterInfo() (err error) {
	if r.state == false {
		err = errors.New("please login first")
		return
	}
	apiUrl := r.Address + configs.RouterInfoUrl
	httpClient := g.Client().SetHeaderMap(r.Headers)
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			g.Dump(err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		return err
	}
	model := gjson.New(res.ReadAllString()).Get("data.model").String()
	g.Dump("获取到机器型号" + model)
	return
}
