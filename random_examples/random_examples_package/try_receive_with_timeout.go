package random_examples_package

import (
	"time"
)

func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration):
		return 0, true, false
	}
}
