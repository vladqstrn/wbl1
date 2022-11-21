package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

func worker(wId int, jobs <-chan int, ctx context.Context) {
	for {
		select {
		case <-jobs:
			fmt.Printf("worker: %d got vaulue: %d\n", wId, <-jobs)
		case <-ctx.Done():
			fmt.Printf("worker: %d finished work. context done!", wId)
			return
		}

	}
}

// pushToChan пишет в канал рандомные числа
func pushToChan(jobs chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("termination of data transmission ...")
			return
		default:
			//time.Sleep(500 * time.Millisecond)
			jobs <- rand.Intn(100)
		}

	}
}

func StartFour() {
	var workersNum int
	fmt.Print("Введите колличество воркеров: ")
	fmt.Fscan(os.Stdin, &workersNum)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	jobs := make(chan int, 3)

	for w := 1; w <= workersNum; w++ {
		go worker(w, jobs, ctx)
	}

	pushToChan(jobs, ctx)
}
