package rconfig

type RouterConfig struct {
	RouterIP       string
	RouterAddress  string
	RouterPassword string
}

var (
	DefaultRouterIP      = "192.168.31.1"
	DefaultRouterAddress = "http://" + DefaultRouterIP
)
