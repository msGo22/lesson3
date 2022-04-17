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

var list = map[string]bool{}
var mutex = new(sync.RWMutex)

func main() {
	key := "a3"
	difficult := 3
	withGoroutine(key, difficult, 100)
}

func md5Miner(input string, i int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	var rnd *rand.Rand
	var generatedKey string
	for {
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
		generatedKey = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(rnd.Int()))))
		if len(strings.Split(generatedKey, input)) > i {
			mutex.Lock()
			if _, ok := list[generatedKey]; !ok {
				defer mutex.Unlock()
				list[generatedKey] = true
				return
			}
			mutex.Unlock()
		}
	}
}

func withGoroutine(key string, difficult int, count int) {
	wg := new(sync.WaitGroup)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go md5Miner(key, difficult, wg)
	}
	wg.Wait()
	fmt.Println(list)
	fmt.Println(len(list))
}
