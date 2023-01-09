package main

import (
	"fmt"
)

func minimumBribes(q []int32) {
	swaps := 0
	l := int32(len(q))
	for i, j := l-1, l-2; i >= 0; {
		if i+1 == q[i] {
			i--
			j = i - 1
			continue
		}

		if i-j > 2 {
			fmt.Println("Too chaotic")
			return
		}

		q[i], q[j] = q[j], q[i]
		j--
		swaps++
	}

	fmt.Println(swaps)
}

func main() {
	// 3 Bribes
	minimumBribes([]int32{2, 1, 5, 3, 4})
	// Chaotic
	minimumBribes([]int32{2, 5, 1, 3, 4})
}