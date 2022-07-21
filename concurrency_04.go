// channel!
package main

import (
	"fmt"
	"time"
)

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}

	//menghindari deadlock
	close(c)
}

// Waitgroup untuk mencegah isu tadi.
func main() {
	c := make(chan string)
	go say("A", c)

	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	//simpler solution
	for msg := range c {
		fmt.Println(msg)
	}
}
