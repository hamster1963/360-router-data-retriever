package rglobal

import "bytes"

type PKCS7Encoder struct {
	BlockSize int
}

func (p *PKCS7Encoder) Encode(src []byte) []byte {
	padding := p.BlockSize - len(src)%p.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}
