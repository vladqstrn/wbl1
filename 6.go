package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины
func Six() {
	quit := make(chan int)
	data := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-data:
				fmt.Println(<-data)

			case <-quit:
				fmt.Println("quit")
				return
			case <-ctx.Done():
				fmt.Println("context done!")
				return
			default:
				time.Sleep(100 * time.Microsecond)
			}
		}
	}()
	// горутина записывает в канал случайное число, ждет пока что-то запишут в канал  quite или отменяется по контексту
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			case <-ctx.Done():
				fmt.Println("context done!")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				data <- rand.Intn(100)
			}
		}
	}()
	quit <- 5
	//time.Sleep(7 * time.Second)
	close(data)
	close(quit)
}
