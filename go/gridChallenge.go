package main

import (
	"fmt"
	"sort"
	"strings"
)

func gridChallenge(grid []string) string {
	colOrder := true
	returnValue := "NO"
	sortStrings(grid)
	for i := 0; i < len(grid)-1; i++ {
		if colOrder {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] > grid[i+1][j] {
					colOrder = false
					break
				}
			}
		} else {
			break
		}
	}

	if colOrder {
		returnValue = "YES"
	}

	return returnValue
}

func sortStrings(grid []string) {
    for i := 0; i < len(grid); i++ {
        gridString := strings.Split(grid[i], "")
        sort.Slice(gridString, func (i, j int) bool {
            return gridString[i] < gridString[j]    
        })
        grid[i] = strings.Join(gridString, "")
    }
}

func main() {
	fmt.Println(gridChallenge([]string{"abc", "ade", "efg"}))
}