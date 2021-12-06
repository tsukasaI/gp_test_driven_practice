package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func CountDown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1)
	fmt.Fprint(out, finalWord)
}

func main() {
	CountDown(os.Stdout)
}
