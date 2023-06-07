package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func RestVariables() {
	Name = "Peter"
	State = true
	Salary = 1577.66
	Date = time.Now()
	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToText(num int) (bool, string) {
	text := strconv.Itoa(num)
	return true, text
}