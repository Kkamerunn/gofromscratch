package exercises

import "strconv"

func ConvertToIntAndShow(str string) (int, string) {
	myInteger, err := strconv.Atoi(str)
	if err != nil {
		return 0, "There was an error " + err.Error()
	}
	if myInteger > 100 {
		return myInteger, "Es mayor a 100"
	} else if  myInteger == 100 {
		return myInteger, "Es igual a 100"
	} else {
		return myInteger, "Es menor a 100"
	}
}