package ABC

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func MakeUrl(x int, R string) string {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	max := x
	var min int
	if R == "e" || R == "f" {
		min = 126
	} else if R == "g" || R == "h" || R == "ex" {
		if R == "ex" {
			R = "h"
		}
		min = 212
	} else if R == "a" || R == "b" || R == "c" || R == "d" {
		min = 0
	} else {
		return "そ、そんな問題は用意できないっぴ、、、"
	}
	var tmp int = rand.Intn(max-min) + min
	ran := fmt.Sprintf("%d", tmp)
	if tmp < 100 && tmp > 10 {
		ran = "0" + ran
	}
	if tmp < 10 {
		ran = "00" + ran
	}
	//ran := string(rand.Intn(max-min) + min)
	url := "https://atcoder.jp/contests/abc" + ran + "/tasks/abc" + ran + "_" + R
	//https://atcoder.jp/contests/abc258/tasks/abc258_b
	return url
}
