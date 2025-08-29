package cntr

import (
	"math/rand"
	"sync"
	"time"
)

var randPool sync.Pool = sync.Pool{
	New: func() any {
		source := rand.NewSource(time.Now().UnixNano())
		return rand.New(source)
	},
}

func applyRand(def func(randGenerator *rand.Rand)) {
	randGenerator := randPool.Get().(*rand.Rand)
	defer randPool.Put(randGenerator)
	def(randGenerator)
}

// 返回隨機 int
var RandInt = func(n int) int {
	var value int
	applyRand(func(randGenerator *rand.Rand) {
		value = randGenerator.Intn(n)
	})
	return value
}
