package main

import (
	"errors"
	"fmt"
)

func main() {
	i, _ := calculate(1, 2, "/")
	fmt.Println(i)
}

func calculate(a int, b int, operate string) (int, error) {
	switch operate {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	}
	errors.New("error operate: " + operate)
	return 0, nil
}
