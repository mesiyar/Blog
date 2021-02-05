package cache

import (
	"encoding/json"
	"fmt"
	"wechatNotify/pkg/util"
)

const Key = "article_"

func Set(id int, data interface{}) error {
	redis := util.Redis{}
	key := fmt.Sprintf("%s%d", Key, id)
	if err := redis.Set(key, data, 86400); err != nil {
		return err
	}
	return nil
}

func Exists(id int) (data interface{}, b bool) {
	redis := util.Redis{}
	key := fmt.Sprintf("%s%d", Key, id)
	b = false
	if redis.Exists(key) {
		if res, err := redis.Get(key); err == nil {
			json.Unmarshal(res, &data)
			return
		}
	}
	return
}

func Del(id int) (bool , error) {
	redis := util.Redis{}
	key := fmt.Sprintf("%s%d", Key, id)
	return redis.Delete(key)
}
