package middleware

import (
	"encoding/json"
	"galen-gvm/global"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type LogLayout struct {
	Time      time.Time
	Path      string        // 访问路径
	Query     string        // 携带query
	Body      string        // 携带body数据
	IP        string        // ip地址
	UserAgent string        // 代理
	Error     string        // 错误
	Cost      time.Duration // 花费时间
}

func Record(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	body, _ := ctx.GetRawData()

	ctx.Next()
	cost := time.Since(start)
	layout := LogLayout{
		Time:      time.Now(),
		Path:      path,
		Query:     query,
		Body:      string(body),
		IP:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Error:     strings.TrimRight(ctx.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
		Cost:      cost,
	}
	res, _ := json.Marshal(layout)
	global.GVA_LOG.Info(string(res))
}
