package main

import "fmt"

func main()  {

	f()

	// os.Exit(1)
}

func f()  {
	defer func() {
		r := recover()
		fmt.Printf("recover --> %+v",r)
	}()
	panic("oh no0o...")
}