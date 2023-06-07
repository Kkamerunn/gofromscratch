package main

import (
	"fmt"
	"github.com/Kkamerunn/gofromscratch/variables"
)

func main() {
	state, text := variables.ConvertToText(1330)
	fmt.Println(state)
	fmt.Println(text)
}