package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountDownOperationSpy struct {
	Calls []string
}

func (s *CountDownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountDownOperationSpy) Write(_ []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &CountDownOperationSpy{}

		CountDown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spy := &CountDownOperationSpy{}

		CountDown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("got %v want %v", spy.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
