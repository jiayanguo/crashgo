package main

import (
	"io"
	"os"
	"fmt"
)

func main(){
	fileName := os.Args[1]
	file, err := os.Open(fileName) 
	
	if err != nil {
		fmt.Println("Read file error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}