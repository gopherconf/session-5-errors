package main

import (
	"fmt"
	"strconv"
)

type customErr struct {
	Description string
	Code        int
}

func (e customErr) Error() string {
	return e.Description + " with code:" + strconv.Itoa(e.Code)
}

func main() {
	_, err := doSomething(0)
	if err != nil {
		fmt.Print("1st Error\n")
		panic(err)
	}

	_, err = doSomething(2)
	if err != nil {

		// read other fields from our custom error struct
		// cast err to customErr by `.(customErr)` and check cast result with `ok` flag
		//Notice: this block added after video record
		if cr, ok := err.(customErr); ok {
			fmt.Printf("\n ---- %+v ----- \n", cr)
		}

		fmt.Print("2nd Error\n")
		panic(err)
	}

}

func doSomething(i int) (string, error) {
	if i == 0 {
		return "OK!", nil
	}

	return "", customErr{
		Description: "i is not equal to 0",
		Code:        -1,
	}
}
