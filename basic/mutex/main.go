package main

import (
	"sync"
)

var depo = map[int]int{}

func main() {
	mutex := new(sync.RWMutex)
	for i := 0; i < 100; i++ {
		deneme2(i, mutex)
	}
}

func deneme2(deger int, mutex *sync.RWMutex) int {
	mutex.RLock()
	defer mutex.RUnlock()
	if val, ok := depo[deger]; ok {
		return val
	}
	return 0
}
