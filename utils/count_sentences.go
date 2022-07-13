package utils

import (
	"github.com/neurosnap/sentences/english"
)
func CountSentences(str string,ch chan int) {
	token,err := english.NewSentenceTokenizer(nil)
	if err != nil {
		ch <- 0
		return 
	}
	setences := token.Tokenize(str)
	ch <- len(setences)
}
