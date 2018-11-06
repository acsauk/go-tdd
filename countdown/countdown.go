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

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalword)
}

func main() {
	Countdown(os.Stdout)
}