// Package codetop provides rand10 with rand7
package codetop

import (
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(7)
}

func rand10() int {
	for {
		a := rand7()
		b := rand7()
		idx := (a-1)*7 + b
		if idx <= 40 {
			return 1 + idx%10
		}

		a = idx - 40
		b = rand7()
		idx = (a-1)*7 + b
		if idx <= 60 {
			return 1 + idx%10
		}

		a = idx - 60
		b = rand7()
		idx = (a-1)*7 + b
		if idx <= 20 {
			return 1 + idx%10
		}
	}
}
