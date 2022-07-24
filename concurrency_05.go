// buffered channel
package main

import "fmt"

func main() {
	// tambahin parameter kedua (capacity) untuk menghindari deadlock
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	// deadlock, karena line 7 tidak akan
	// lanjut apabila tidak ada yang menerima channelnya!

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
