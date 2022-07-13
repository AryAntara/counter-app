package utils

import (
	"strings"
)
func CountWord(str string,ch chan int) {
	length := len(strings.Fields(str))
	ch <- length
}

