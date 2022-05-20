package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func JsonDecode(str string) (data map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(str), &data)
	return data, err
}

func JsonEncode(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	return string(bytes), err
}
