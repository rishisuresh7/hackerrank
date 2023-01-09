package main

import "fmt"

func miniMaxSum(arr []int32) {
    sum := int64(0)
    min := arr[0]
    max := arr[0]
    for i := 0; i < 5; i++  {
        if min > arr[i] {
            min = arr[i]
        }

        if arr[i] > max {
            max = arr[i]
        }
        
        sum += int64(arr[i])
    }
    
    fmt.Printf("%d %d\n", sum - int64(max), sum - int64(min))
}

func main() {
	miniMaxSum([]int32{1, 3, 5, 7, 9})
}