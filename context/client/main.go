package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	println("❯ context.client start")

	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		fmt.Println("🗙 prepare client error:", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("🗙 do client error: %v", err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

	println("❯ context.client end")
}
