package main

import (
	"errors"
	"log"
	"os"
)

func main()  {
	file,err := os.Open("./dose-not-exists.file")
	if err != nil{
		if errors.Is(err,os.ErrNotExist){
			log.Fatal("Our file doesn't exists!")
		}else{
			panic(err)
		}
	}

	defer file.Close()
}
