package util

import (
	"crypto/md5"
	"encoding/hex"
)

type Encoder struct {

}

func (encoder *Encoder) Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
