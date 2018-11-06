package main

import (
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
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalword)
}

func main() {
	Countdown(os.Stdout)
}