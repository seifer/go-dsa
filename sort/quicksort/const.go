package quicksort

import (
	"math/rand"
	"time"
)

var rnd = rand.New(
	rand.NewSource(
		time.Now().UnixNano(),
	),
)
