package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ReallyLongAndExpensiveOperation(channel chan int) {
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		i := rand.Intn(1000)
		time.Sleep(time.Duration(i) * time.Millisecond)
		channel <- i
	}
}

func main() {
	channel := make(chan int)
	go ReallyLongAndExpensiveOperation(channel)
	for {
		fmt.Println("I waited", <-channel, "ms for this message.")
	}
}
