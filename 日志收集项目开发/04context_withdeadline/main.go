package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Microsecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(time.Second * 1):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}
