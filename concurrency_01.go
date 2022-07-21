// simple concurrency
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}

// simple concurrency
func main() {
	go say("A")
	say("B")
}
