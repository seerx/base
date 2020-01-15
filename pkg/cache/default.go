/**
默认缓存
*/
package cache

import (
	"fmt"
	"github.com/seerx/base/pkg/cache/memory"
	"time"
)

var cacheInstance Cache

type Provider int

const (
	Local Provider = iota
	Redis
)

// InitCache 初始化缓存
// 使用 Local 时，默认设置 30 分钟检查一次超时缓存，默认缓存不超时
// Redis 未实现
func InitCache(provider Provider) {
	if provider == Local {
		cacheInstance = memory.New(10, 30*time.Minute, 0)
	}

	panic(fmt.Errorf("Cann't find provider [%d]", provider))
}

// GetInstance 获取缓存实例
func GetInstance() Cache {
	return cacheInstance
}

// Exists 判断缓存是否存在
func Exists(key string) bool {
	return cacheInstance.Exists(key)
}

// Set 设置缓存
func Set(key string, value interface{}) error {
	return cacheInstance.Set(key, value)
}

// SetX 设置缓存，附带超时
func SetX(key string, value interface{}, expire time.Duration) error {
	return cacheInstance.SetX(key, value, expire)
}

// Get 获取缓存
// error 不是空时，不存在此 key
func Get(key string) (interface{}, error) {
	return cacheInstance.Get(key)
}

// Remove 移除缓存
func Remove(key string) error {
	return cacheInstance.Remove(key)
}

// Info 获取缓存整体状态信息
func Info() string {
	return cacheInstance.Info()
}
