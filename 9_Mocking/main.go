package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprint(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
