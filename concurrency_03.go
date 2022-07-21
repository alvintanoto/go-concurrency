// wait group
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}

// Waitgroup untuk mencegah isu tadi.
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		say("A")
		wg.Done()
	}()

	wg.Wait()
}
