package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Signal struct {
	outChan chan int
}

func main() {
	signal := Signal{outChan: make(chan int, 10)}

	go signal.flash()
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-signal.getOutChan():
			fmt.Println("Get message from FLASH")
		case <-ticker.C:
			fmt.Println("Get message from TICKER")
		}
	}
}

func (s *Signal) flash() {
	for {
		time.Sleep(time.Second * 2)
		s.outChan <- rand.Int()
	}
}

func (s *Signal) getOutChan() <-chan int {
	return s.outChan
}
