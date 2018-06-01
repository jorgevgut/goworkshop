package examples

import (
	"fmt"
)

func printNumber(number int) {
	fmt.Println(number)
}

// GoRoutines example illustrate how to run a goroutine
func GoRoutines() {
	for i := 0; i < 50; i++ {
		go printNumber(i)
	}
}

// GoRoutinesBlocked example illustrate how to run a goroutine
func GoRoutinesBlocked() {
	c := make(chan int, 1)
	for i := 0; i < 50; i++ {
		go printNumber(i)
	}
	<-c
}

// Channels is for trying out this feat
func Channels() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// ChannelsBuffered is for trying out this feat
func ChannelsBuffered() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
}
