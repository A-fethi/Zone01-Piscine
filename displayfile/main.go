package main

import (
	"fmt"
	"os"
	//"io/ioutil"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		data, _ := os.ReadFile(os.Args[1])
		os.Stdout.Write(data)
	}
}
