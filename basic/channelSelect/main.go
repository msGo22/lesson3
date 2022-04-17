package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	pingKanal := make(chan bool, 1)
	pongKanal := make(chan bool, 1)
	bekletme := make(chan bool, 1)
	go pingPong(ctx, pingKanal, pongKanal, bekletme)
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
		fmt.Println("Cancel ile bitti")
	}()
	pingKanal <- true
	<-bekletme
}

func pingPong(ctx context.Context, pi, po, bekletme chan bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Oyun bitti")
			close(bekletme)
			return
		case <-pi:
			fmt.Println("ping")
			po <- true
		case <-po:
			fmt.Println("pong")
			pi <- true
		}
	}
}
