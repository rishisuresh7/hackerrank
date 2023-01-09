package main

import "fmt"

func plusMinus(arr []int32) {
    length := float64(len(arr))
    counts := map[int]float64{
        0: 0,
        1: 0,
        -1: 0,
    }
    for _, value := range arr {
        if value > 0 {
            counts[1]++
        }   else if value < 0 {
            counts[-1]++
        } else {
            counts[0]++
        }
    }

    fmt.Printf("%.6f\n", counts[1]/length)
    fmt.Printf("%.6f\n", counts[-1]/length)
    fmt.Printf("%.6f\n", counts[0]/length)    
}

func main() {
	plusMinus([]int32{1, 1, 0, -1, -1})
}