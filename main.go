package main

import (
	"fmt"
)

func main() {
	tokens := InitTokenizer()
	fmt.Println(AnalyzeData("Datta", tokens))
}
