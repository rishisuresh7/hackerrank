package main

import "fmt"

func countingSort(arr []int32) []int32 {
    result := make([]int32, 100)
    for i := 0; i < len(arr); i++ {
        result[arr[i]]++
    }

    return result
}

func main() {
	fmt.Println(countingSort([]int32{1, 1, 3, 2, 1}))
}