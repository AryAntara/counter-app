package utils

import (
	"unicode"
)
func CountCharWithOutSpace(str string,ch chan int) { 
	length := 0
	for _,s := range str {
		if unicode.IsSpace(s) == false {
			length++
		}
	}
	ch <- length
}
