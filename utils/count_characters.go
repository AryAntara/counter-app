package utils

func CountChar(str string,ch chan int){ 
	ch <- len(str)
}
