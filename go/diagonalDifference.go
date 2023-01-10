package main

import "fmt"

func abs(a int32) int32 {
    if a < 0 {
        return -1 *a
    }

    return a
}

func diagonalDifference(arr [][]int32) int32 {
    d1Sum := int32(0)
    d2Sum := d1Sum
    j := len(arr) - 1
    for i := 0; i < len(arr); i++ {
        d1Sum += arr[i][i]
        
        d2Sum += arr[i][j-i]
    }

    return abs(d1Sum - d2Sum)
}

func main() {
	matrix := [][]int32{
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}
	fmt.Println(diagonalDifference(matrix))
}