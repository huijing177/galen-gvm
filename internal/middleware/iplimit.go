package middleware

import (
	"sync"
	"time"

	"galen-gvm/global"

	"github.com/gin-gonic/gin"
	cmp "github.com/orcaman/concurrent-map/v2"
	"go.uber.org/ratelimit"
)

var (
	IPLimitMap = cmp.New[ratelimit.Limiter]()
	IPOnline   = cmp.New[int64]()
	once       sync.Once
)

func IPLimit(ctx *gin.Context) {
	once.Do(func() {
		go func() {
			// 每隔600秒，清理一次两个map，防止长时间运行后，内存占用太多
			ticker := time.NewTicker(time.Second * 600)
			for range ticker.C {
				global.GVA_LOG.Info("开始清理IPLimitMap和IPOnline")
				now := time.Now().Unix()
				for k, v := range IPOnline.Items() {
					// 清理掉6个小时没上线的
					if now-v > 60*60*6 {
						IPOnline.Remove(k)
						IPLimitMap.Remove(k)
					}
				}
			}
		}()
	})

	cliIP := ctx.ClientIP()
	if cliIP == "" {
		ctx.Abort()
		return
	}
	IPOnline.Set(cliIP, time.Now().Unix())
	v, ok := IPLimitMap.Get(cliIP)
	if !ok {
		// 限制每秒只能60次调用，多了绝对是恶意的，直接给他阻塞
		IPLimitMap.Set(cliIP, ratelimit.New(60))
		v, _ = IPLimitMap.Get(cliIP)
	}
	v.Take()
	ctx.Next()
}
