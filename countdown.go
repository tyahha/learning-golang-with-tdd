package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
