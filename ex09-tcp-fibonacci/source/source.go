package source

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type Answer struct {
	Answer *big.Int
	Time   time.Duration
}

const IP = "127.0.0.1:80"

func FibonacciResult(n int) *big.Int {
	fn := make(map[int]*big.Int)

	for count := 0; count <= n; count++ {
		var f = big.NewInt(0)
		if count <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f

	}
	return fn[n]
}

func Reader(scan string) int64 {
	input, err := strconv.ParseInt(scan, 10, 64)
	CheckErr(err)
	return input
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error message:", err)
		return
	}
}
