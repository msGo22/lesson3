package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	v := make(chan int, 1)
	go worker("ali", wg, v, ctx)
	go worker("veli", wg, v, ctx)
	go worker("cengiz", wg, v, ctx)
	go worker("kamil", wg, v, ctx)
	go worker("yavuz", wg, v, ctx)
	go worker("sevket", wg, v, ctx)
	go worker("bekir", wg, v, ctx)
	go manager(v, cancel)
	<-ctx.Done()
	wg.Wait()
	close(v)
}

func worker(workerName string, wg *sync.WaitGroup, value chan int, ctx context.Context) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case number := <-value:
			fmt.Println(workerName, " ", number, " işini yaptı")
		case <-ctx.Done():
			fmt.Println(workerName, " ayrıldı")
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func manager(value chan int, cancel context.CancelFunc) {
	i := 1
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			cancel()
			return
		default:
			i++
			value <- i
		}
	}

}
