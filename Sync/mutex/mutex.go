package main

import (
	. "fmt"
	"sync"
	"time"
)

func main() {
	// Sync 구현
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			Println(i, "start")
			time.Sleep(time.Second)
			Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	Scanln(&input)
}
