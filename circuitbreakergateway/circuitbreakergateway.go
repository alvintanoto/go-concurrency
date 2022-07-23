package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	cb := gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:        "cb",
			MaxRequests: 5,
			Timeout:     30 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.TotalFailures > 5
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		},
	)

	for {
		body, err := cb.Execute(func() (interface{}, error) {
			resp, err := http.Get("http://0.0.0.0:3000/")
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			return body, err
		})

		if err != nil {
			fmt.Println(err)
		}

		if body != nil {
			fmt.Println(string(body.([]byte)))
		}

		time.Sleep(time.Millisecond * 500)
	}

}
