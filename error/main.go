package main

import (
	"fmt"
	"log"
)

func main() {
	err := f()

	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Print("Hi there...")
}

func f() error {
	return fmt.Errorf("- %d - %s",1,"Oh no...")
	// return errors.New("Hi new error :)")
}
