package rutils

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcache"
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
	// 进行路由器实例的构建
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
	apiUrl := r.RouterAddress + rconfig.GetRandStringUrl
	// g.Dump(apiUrl)
	httpClient := gclient.New()
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
		err = errors.New("GetRandomString status code error")
		return err
	}
	r.randStr = gconv.String(gconv.Map(res.ReadAllString())["rand_key"])
	// g.Dump(gtime.Now().String() + " Get RandomKey " + r.randStr)
	return nil
}

// GenerateAesString 生成加密字符串
func (r *Router) GenerateAesString() (err error) {
	// 判断随机字符串是否为空
	if r.randStr == "" {
		g.Dump("randStr is empty")
		err := r.GetRandomString()
		if err != nil {
			g.Dump(err)
			return err
		}
	}
	// randKey := "fbf8a1ca3b31ace17adece7f6941a278017ff28b58200c5a153e07f5dc840b3f"
	decodeString, err := hex.DecodeString(r.randStr[32:])
	if err != nil {
		g.Dump(err)
		return
	}
	block, err := aes.NewCipher(decodeString)
	if err != nil {
		panic(err)
	}
	encryptor := cipher.NewCBCEncrypter(block, rconfig.DefaultAesIv)
	p7 := rglobal.PKCS7Encoder{BlockSize: 16}
	padded := p7.Encode([]byte(r.RouterPassword))
	cipherText := make([]byte, len(padded))
	encryptor.CryptBlocks(cipherText, padded)
	r.aesStr = hex.EncodeToString(cipherText)
	if r.aesStr == "" {
		g.Dump("aesStr is empty")
		return errors.New("aesStr is empty")
	}
	// g.Dump(gtime.Now().String() + " Generate AESKey " + r.aesStr)
	return
}

func (r *Router) Login() (err error) {
	loginUrl := r.RouterAddress + rconfig.LoginUrl
	payload := "user=admin&pass=" + r.randStr[:32] + r.aesStr + "&form=1"
	httpClient := gclient.New()
	httpClient.SetHeaderMap(rconfig.DefaultHeaders)
	res, err := httpClient.Post(context.Background(), loginUrl, payload)
	if err != nil {
		g.Dump(err)
	}
	defer func(res *gclient.Response) {
		err := res.Close()
		if err != nil {
			g.Dump(err)
		}
	}(res)
	if res.StatusCode != 200 {
		g.Dump(res.ReadAllString())
		err = errors.New("status code error")
		return err
	}
	// set cookie and token
	resData := gconv.Map(res.ReadAllString())
	r.cookie = res.Header.Get("Set-Cookie")
	r.token = resData["Token-ID"].(string)
	// g.Dump(gtime.Now().String() + " Login Success ")
	r.Headers["Cookie"] = r.cookie
	r.Headers["Token-ID"] = r.token
	err = gcache.Set(context.Background(), "loginState", 1, 30*time.Second)
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
