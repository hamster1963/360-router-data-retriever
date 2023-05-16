package rglobal

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/v2/os/glog"
)

type AesEncoder struct {
	Key       []byte
	Iv        []byte
	PKCS7ESrc []byte
}

func (a *AesEncoder) AesStringEncoder() (aesString string, err error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}
	encryptor := cipher.NewCBCEncrypter(block, a.Iv)
	p7 := PKCS7Encoder{BlockSize: 16}
	padded := p7.Encode(a.PKCS7ESrc)
	cipherText := make([]byte, len(padded))
	encryptor.CryptBlocks(cipherText, padded)
	aesString = hex.EncodeToString(cipherText)
	if aesString == "" {
		glog.Warning(context.TODO(), "aesStr is empty")
		return "", errors.New("aesStr is empty")
	}
	return aesString, nil
}
