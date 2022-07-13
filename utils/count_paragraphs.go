package utils

import (
	"unicode"
)
func CountParagraphs(str string,ch chan int ){
	// create new line
	length := 0
	for i,_ := range str {
		next := i + 1
		nextAgain := i + 2
		if(i == len(str) - 2){
			break
		}
		//fmt.Println(string(str[i]))
		if string(str[i]) == "." && unicode.IsSpace(rune(str[next])) && unicode.IsSpace(rune(str[nextAgain])){
			length++
		}
	}
	if length == 0 {
		length = 1
	}
	ch <- length
}
