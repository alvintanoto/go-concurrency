// simple concurenccy negative case
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

// Program akan exit dan goroutine tidak akan jalan karena main function selesai.
func main() {
	go say("A")
	go say("B")
}
