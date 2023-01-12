package main

import "fmt"

func caesarCipher(s string, k int32) string {
    res := ""
    for i :=0 ; i < len(s); i++ {
        in := int32(s[i])
        re := in + (k%26)
        if(in >= 65 && in <=90) {
            if re > 90 {
                re = (re - 90) + 65 - 1
            }
            res += fmt.Sprintf("%c", re)
        } else if (in >= 97 && in <= 122) {
            if re > 122 {
                re = (re - 122) + 97 - 1
            }
            res += fmt.Sprintf("%c", re)
        } else {
            res += fmt.Sprintf("%c", in)
        }
    }

    return res
}

func main() {
	fmt.Println(caesarCipher("middle-Outz", 2))
}