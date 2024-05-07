package middleware

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestRateLimit(t *testing.T) {
	rl := ratelimit.New(100) // per second
	rl.Take()
	time.Sleep(time.Second * 5)
	start := time.Now()
	prev := time.Now()
	for i := 0; i < 30; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
		// if i == 2 {
		// 	time.Sleep(time.Second)
		// }
	}
	fmt.Println(time.Since(start))
}
