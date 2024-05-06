package redis

import (
	"context"
	"galen-gvm/global"
	"time"

	"go.uber.org/zap"
)

const (
	_RedisCaptchaPrefix = "Captcha_"
)

type CaptchaRedis struct {
	Prefix      string
	Context     context.Context
	ExpiresTime time.Duration
}

func NewCaptchaRedis() *CaptchaRedis {
	return &CaptchaRedis{
		Prefix:      _RedisCaptchaPrefix,
		Context:     context.TODO(),
		ExpiresTime: time.Second * 60,
	}
}

func (c *CaptchaRedis) Set(id string, value string) error {
	err := global.GVA_REDIS.Set(c.Context, c.Prefix+id, value, c.ExpiresTime).Err()
	if err != nil {
		global.GVA_LOG.Error("CaptchaRedis set err:", zap.Error(err))
	}
	return err
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (c *CaptchaRedis) Get(id string, clear bool) string {
	v, err := global.GVA_REDIS.Get(c.Context, c.Prefix+id).Result()
	if err != nil {
		global.GVA_LOG.Error("CaptchaRedis get err:", zap.Error(err))
		return ""
	}
	if clear {
		err = global.GVA_REDIS.Del(c.Context, c.Prefix+id).Err()
		if err != nil {
			global.GVA_LOG.Error("CaptchaRedis del err:", zap.Error(err))
			return v
		}
	}
	return v
}

// Verify captcha's answer directly
func (c *CaptchaRedis) Verify(id, answer string, clear bool) bool {
	v := c.Get(id, clear)
	return v == answer
}
