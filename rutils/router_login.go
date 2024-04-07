package rutils

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hamster1963/360-router-data-retriever/rconfig"
	"github.com/hamster1963/360-router-data-retriever/rglobal"
	"time"
)

type LoginMethod interface {
	InitRouter(config *rconfig.RouterConfig) *Router
	Login() error
	CheckLogin() (bool, error)
}

type AesMethod interface {
	GetRandomString() error
	GenerateAesString() error
}

type Router struct {
	*rconfig.RouterConfig
	aesIv     []byte
	inHeaders map[string]string
	randStr   string
	aesStr    string
	token     string
	cookie    string
	Headers   map[string]string
}

func NewRouter() *Router {
	return new(Router)
}

func (r *Router) InitRouter(config *rconfig.RouterConfig) (newRouter *Router) {
	// init router
	newRouter = &Router{
		RouterConfig: config,
		Headers:      rconfig.DefaultHeaders,
	}
	newRouter.Headers["Host"] = config.RouterIP
	newRouter.Headers["Origin"] = config.RouterAddress
	newRouter.Headers["Referer"] = config.RouterAddress + "/"
	return
}

// GetRandomString 获取随机字符串
func (r *Router) GetRandomString() (err error) {
	var (
		apiUrl     = r.RouterAddress + rconfig.GetRandStringUrl
		httpClient = gclient.New()
	)
	// Get RandomKey by http
	res, err := httpClient.Get(context.Background(), apiUrl)
	if err != nil {
		return err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			glog.Warning(context.Background(), "GetRandomString close error", err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		glog.Warning(context.Background(), "GetRandomString status code error", res.StatusCode)
		return err
	}
	r.randStr = gconv.String(gconv.Map(res.ReadAllString())["rand_key"])
	return nil
}

// GenerateAesString 生成加密字符串
func (r *Router) GenerateAesString() (err error) {
	var (
		ctx = context.TODO()
	)

	// Validate randStr
	if r.randStr == "" {
		glog.Warning(ctx, "randStr is empty")
		err := r.GetRandomString()
		if err != nil {
			return err
		}
	}
	// Example: randKey := "fbf8a1ca3b31ace17adece7f6941a278017ff28b58200c5a153e07f5dc840b3f"
	decodeString, err := hex.DecodeString(r.randStr[32:])
	if err != nil {
		return err
	}
	aesEncoder := rglobal.AesEncoder{
		Key:       decodeString,
		Iv:        rconfig.DefaultAesIv,
		PKCS7ESrc: []byte(r.RouterPassword),
	}
	r.aesStr, err = aesEncoder.AesStringEncoder()
	if err != nil {
		return err
	}
	return
}

func (r *Router) Login() (err error) {
	var (
		loginUrl   = r.RouterAddress + rconfig.LoginUrl
		payload    = "user=admin&pass=" + r.randStr[:32] + r.aesStr + "&form=1"
		httpClient = gclient.New()
	)

	httpClient.SetHeaderMap(rconfig.DefaultHeaders)
	res, err := httpClient.Post(context.Background(), loginUrl, payload)
	if err != nil {
		return err
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			glog.Warning(context.Background(), err)
		}
	}(res)
	if res.StatusCode != 200 {
		err = errors.New("status code error")
		return err
	}

	// set cookie and token
	resData := gconv.Map(res.ReadAllString())
	r.cookie = res.Header.Get("Set-Cookie")
	r.token = gconv.String(resData["Token-ID"])
	r.Headers["Cookie"] = r.cookie
	r.Headers["Token-ID"] = r.token
	// set login state
	err = gcache.Set(context.Background(), "loginState", 1, 60*time.Second)
	if err != nil {
		return err
	}
	return
}

func (r *Router) CheckLogin() (bool, error) {
	state, err := gcache.Get(context.Background(), "loginState")
	if err != nil {
		return false, err
	}
	if state.Int() == 1 {
		return true, nil
	}
	return false, nil
}
