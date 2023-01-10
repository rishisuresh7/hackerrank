package main

import "fmt"

func lonelyinteger(a []int32) int32 {
    counts := map[int32]int{}
    for _, value := range a {
        if _, ok := counts[value]; !ok {
            counts[value] = 1
            continue
        }
        
        counts[value]++
    }

	index := int32(-1)
    for key, value := range counts {
        if value == 1 {
			index = key
			break
        }
    }

    return index
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1}))
}