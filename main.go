package main

import (
	"github.com/gofiber/fiber/v2"
	"counter/utils"
	"encoding/json"
	"strings"
	//"fmt"
	//"time"
)

func main(){
	app := fiber.New(fiber.Config{
		BodyLimit:50 * 1024 * 1024,
	}) 
	app.Post("/count",counting)
	app.Listen(":8888")
}

type counters struct {
	Word int `json:"word"`
	Char int 	`json:"character"`
	CharOnly int `json:"character without space"`
	Sentences int `json:"sentences"`
	Paragraph int `json:"paragraph"`
	SylLabels int	`json:"syllabels"`
}
func counting(c *fiber.Ctx) error {
	//start := time.Now()
	str := strings.ReplaceAll(string(c.Body()),"%OD%OA","\n")
	// channel
	word := make(chan int,1)
	char := make(chan int,1)
	charWithoutSpace := make(chan int,1)
	syllables := make(chan int,1)
	sentences := make(chan int,1)
	paragraph := make(chan int,1)
	// core go routine
	go utils.CountChar(str,char)
	go utils.CountCharWithOutSpace(str,charWithoutSpace)
	go utils.CountWord(str,word)
	go utils.CountSyllables(str,syllables)
	go utils.CountSentences(str,sentences)
	go utils.CountParagraphs(str,paragraph)
  //end := time.Now()
	//defer fmt.Println("start",start,"\ndone",end)
	allCounter := counters{}
	allCounter.Word = <- word
	allCounter.Char = <- char
	allCounter.CharOnly = <- charWithoutSpace
	allCounter.SylLabels = <- syllables
	allCounter.Paragraph = <- paragraph
	allCounter.Sentences = <- sentences
	strs,_ := json.Marshal(allCounter)
	c.Send(strs)
	// close channel
	close(word)
	close(char)
	close(charWithoutSpace)
	close(syllables)
	close(paragraph)
	close(sentences)
	return nil
}


