package main

import (
	"context"
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func pushToChanFive(ctx context.Context, channel chan int) {
	for i := 0; i < 1000000; i++ {
		channel <- i
	}
	close(channel)
}
func Five() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	channel := make(chan int)
	go pushToChanFive(ctx, channel)

	for {
		select {
		case val, ok := <-channel:
			if ok == false {
				fmt.Println(val, ok, "<-- zero value")
			} else {
				fmt.Println(val, ok)
			}
		case <-ctx.Done():
			fmt.Printf("Context cancelled: %v\n", ctx.Err())
			return
		}

	}
}
