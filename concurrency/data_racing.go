package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := new(sync.Mutex)

	count := 0

	for i := 1; i <= 5; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {
				lock.Lock()
				count++
				fmt.Println(count)
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 5)
}
