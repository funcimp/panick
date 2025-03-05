package panick

import (
	"testing"
	"time"
)

func TestRandomDelay(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RandomDelay did not panic")
		}
	}()

	max := 1 * time.Millisecond
	for range 10 {
		RandomDelay(max, "tests")
	}
}
