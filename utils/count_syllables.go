package utils

import ( 
	"github.com/mtso/syllables" 
	"strings"
	//"sync"
	"fmt"
)

func CountSyllables(str string,ch chan int) {
	counter := make(chan int)
	tocount := 0
	words := []string{}

	if len(strings.Split(str," ")) <= 10000 {
		words = strings.Split(str,"")
	} else if len(strings.Split(str,". ")) <= 10000 {
		words = strings.Split(str,". ")
	} else if len(strings.Split(str,"\n")) <= 10000 {
		words = strings.Split(str,"\n")
	} else {
		ch <- syllables.In(str)
		close(counter)
		return 
	}

	length := 0	
	for _,word := range words {
		go func(str string,ch chan int){
			ch <- syllables.In(str)
		}(word, counter)
		tocount++
	}
	for tocount > 0{
		length += <-counter
		tocount--
	}
	defer close(counter)
	ch <- length
}

