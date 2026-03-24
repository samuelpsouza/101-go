package main

import (
	"os"
	"fmt"
)

func main() {
	caller := os.Args[0]

	fmt.Println("Executed Command: " + caller)
	
	for index, arg := range os.Args[1:] {
		fmt.Println(index, "->" ,arg) 
	}

}