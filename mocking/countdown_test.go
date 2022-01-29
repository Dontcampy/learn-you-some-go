package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type ObservableCountdownOperations struct {
	Calls []string
}

func (s *ObservableCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *ObservableCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		observableSleeper := &ObservableCountdownOperations{}

		Countdown(buffer, observableSleeper)

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
		observableSleeperPrinter := &ObservableCountdownOperations{}
		Countdown(observableSleeperPrinter, observableSleeperPrinter)

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

		if !reflect.DeepEqual(want, observableSleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, observableSleeperPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	observableTime := &ObservableTime{}
	sleeper := ConfigurableSleeper{sleepTime, observableTime.Sleep}
	sleeper.Sleep()

	if observableTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, observableTime.durationSlept)
	}
}
