package partition

import (
	"math/rand"
	"time"
)

var randGen = rand.New(rand.NewSource(time.Now().Unix()))

func randomInRange(min, max int) int {
	return randGen.Intn(max-min) + min
}
