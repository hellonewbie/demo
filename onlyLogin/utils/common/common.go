package common

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

//查到某个值是否在数组中
func InArrayString(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

//sha1加密
func Sha1En(data string) string {
	t := sha1.New()
	_, _ = io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//加密
func EBase64(data string) string {
	enc_str := base64.StdEncoding.EncodeToString([]byte(data))
	return enc_str
}

//解密
func DBase64(data string) string {
	dec_str, _ := base64.StdEncoding.DecodeString(data)
	return fmt.Sprintf("%s", dec_str)
}
