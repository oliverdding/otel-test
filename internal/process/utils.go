package process

import (
	"math/rand"
	"time"
)

var RNG = rand.New(rand.NewSource(time.Now().UnixNano()))

func SleepWithoutSpan() {
	var sleep int64

	switch modulus := time.Now().Unix() % 5; modulus {
	case 0:
		sleep = RNG.Int63n(111)
	case 1:
		sleep = RNG.Int63n(15)
	case 2:
		sleep = RNG.Int63n(213)
	case 3:
		sleep = RNG.Int63n(87)
	case 4:
		sleep = RNG.Int63n(64)
	}

	time.Sleep(time.Duration(sleep) * time.Millisecond)
}
