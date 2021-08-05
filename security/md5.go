package security

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
