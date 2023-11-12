package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//f implements Reader Interface, it has Read Function
	//os.Stdout implements the Writer Interface, it has Write Function
	io.Copy(os.Stdout, f)

}
