package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit := rate.NewLimiter(rate.Every(1*time.Second), 2)
	ctx := context.Background()
	cnt := 0
	for cnt < 10 {
		limit.Wait(ctx)
		fmt.Println("cnt", cnt)
		cnt++
	}
}
