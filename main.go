package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	numbers := [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	channel := make(chan string, 10)

	waiter := &sync.WaitGroup{}
	waiter.Add(10)

	for i := 0; i < 10; i++ {
		go func(y int) {
			defer waiter.Done()

			for a := range channel {
				fmt.Println(y, a)
			}
			time.Sleep(5 * time.Second)
		}(i)
	}

	for _, i := range numbers {
		channel <- i
	}
	close(channel)

	waiter.Wait()

}
