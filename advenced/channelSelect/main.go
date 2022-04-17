package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	v := make(chan int, 1)
	go worker("ali", wg, v)
	go worker("veli", wg, v)
	go worker("cengiz", wg, v)
	go worker("kamil", wg, v)
	go worker("yavuz", wg, v)
	go worker("sevket", wg, v)
	go worker("bekir", wg, v)
	manager(v)
	wg.Wait()
}

func worker(workerName string, wg *sync.WaitGroup, value chan int) {
	for {
		select {
		case number := <-value:
			wg.Add(1)
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(rnd.Int()%5) * time.Second)
			fmt.Println(workerName, " ", number, " işini yaptı sonuç ", faktoriyel(number))
			wg.Done()
		default:
			time.Sleep(time.Second)
		}
	}
}

func faktoriyel(number int) int {
	if number == 1 {
		return 1
	}
	return faktoriyel(number-1) * number
}

func manager(value chan int) {
	for i := 1; i < 20; i++ {
		value <- i
	}

}
