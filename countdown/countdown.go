package main

import (
	"time"
	"os"
	"io"
	"fmt"
)

const(
	finalword = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	
	sleeper.Sleep()
	fmt.Fprint(out, finalword)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}