package configs

var (
	DefaultAesIv = []byte{0x33, 0x36, 0x30, 0x6c, 0x75, 0x79, 0x6f, 0x75, 0x40, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c}

	DefaultHeaders = map[string]string{
		"Accept":       "application/json, text/plain, */*",
		"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"Host":         RouterIP,
		"Origin":       RouterAddress,
		"Referer":      RouterAddress + "/",
		"Cookie":       "",
		"Token-ID":     "",
		"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
	}
)
