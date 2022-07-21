package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go say("sheep", c)

	// disini proses akan STOP hingga mendapat balikan dari channel yang di set
	// for {
	// 	msg, open := <-c

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// simpler solution!
	for msg := range c {
		fmt.Println(msg)
	}

	// ini akan deadlock karena program akan menunggu hasil goroutine selamanya
}

func say(s string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}

	// close channelnya setelah process selesai (untuk menghindari deadlock)
	close(c)
}
