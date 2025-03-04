package panick

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// If panics if the condition is true
func If(condition bool, message string) {
	if condition {
		panic(message)
	}
}

// Sometimes panics if the probability is greater than the random number generated between 0 and 100
func Sometimes(probability uint64, message string) {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	rnd := n.Uint64()
	if rnd < probability {
		panic(message)
	}
}

func OnError(err error) {
	panic(err)
}

// Room creates a Panick Room and executes a function in a safe environment where
// panics are contained and returns the result or an error if a panic occurred
func Room(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panick Room: %v", r)
		}
	}()
	f()
	return err
}

// RandomDelay chooses a random duration up to max and sleeps for that duration in
// a goroutine. It immediately returns the chosen duration.
func RandomDelay(max time.Duration, v any) {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	r := time.Duration(n.Int64())
	Delay(r, fmt.Sprintf("panick: RandomDelay %v", v))
}

// Delay sleeps for d duration and then panics.
func Delay(d time.Duration, v ...any) {
	time.Sleep(d)
	panic(fmt.Sprintln(v...))
}
