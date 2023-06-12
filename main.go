package main

import (
	"fmt"
	/* "github.com/Kkamerunn/gofromscratch/variables" */
	"github.com/Kkamerunn/gofromscratch/exercises"
	// "runtime"
)

func main() {
	/* state, text := variables.ConvertToText(1330)
	fmt.Println(state)
	fmt.Println(text)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("This is not windows, is", os)
	} else {
		fmt.Println("This is windows")
	} */

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("This is Linux")
	// case "darwin":
	// 	fmt.Println("This is Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	num, text := exercises.ConvertToIntAndShow("100")
	fmt.Println(num)
	fmt.Println(text)
}