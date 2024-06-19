package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	println("❯ context.hotel start")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	// ctx := context.WithValue(ctx, "key", "value")
	// ctx.Value("key")
	defer cancel()

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	bookHotel(ctx)
	println("❯ context.hotel end")
}

func bookHotel(ctx context.Context) {
	// ctx := context.TODO()
	select {
	case <-ctx.Done():
		fmt.Println("🗙 time out to book hotel")
	case <-time.After(time.Second * 5):
		fmt.Println("✓ book hotel done with success")
	}
}
