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

// Maybe panics if the probability is greater than the random number generated between 0 and 100
func Maybe(probability uint64, message string) {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	rnd := n.Uint64()
	if rnd < probability {
		panic(message)
	}
}

func Error(err error) {
	panic(err)
}

// Func returns a function that panics after calling f.
func Func(f func()) func() {
	return func() {
		f()
		panic("panick: Func")
	}
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

func PanicWithMessage()
