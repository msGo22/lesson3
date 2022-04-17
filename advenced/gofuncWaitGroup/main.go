package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	key := "a3e"
	difficult := 3
	regularWay(key, difficult, 35)
	WithGoroutine(key, difficult, 35)
}

func md5Miner(input string, i int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	var rnd *rand.Rand
	var generatedKey string
	for len(strings.Split(generatedKey, input)) < i {
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
		generatedKey = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rnd.Int()))))
	}
	fmt.Print(generatedKey, " ,")
}

func regularWay(key string, difficult int, count int) {
	regularStarted := time.Now()
	for i := 0; i < count; i++ {
		md5Miner(key, difficult, nil)
	}
	fmt.Println("Regular work time :", time.Now().Sub(regularStarted))
}

func WithGoroutine(key string, difficult int, count int) {
	withGoStarted := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go md5Miner(key, difficult, wg)
	}
	wg.Wait()
	fmt.Println("Goroutine work time :", time.Now().Sub(withGoStarted))
}
