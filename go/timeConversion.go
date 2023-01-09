package main

import (
	"fmt"
	"strings"
	"strconv"
)

func timeConversion(s string) string {
    timeString := s[0:len(s)-2]
    parts := strings.Split(timeString, ":")    
    if s[len(s)-2:] == "AM" {
        if parts[0] == "12" {
            parts[0] = "00"
        }
    } else {
        integer, _ := strconv.Atoi(parts[0])
        if integer < 12 {
            integer += 12
        }
        if integer > 23 {
            integer = 0
        }
        
        parts[0] = fmt.Sprintf("%d", integer)
    }

	return strings.Join(parts, ":")
}

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}