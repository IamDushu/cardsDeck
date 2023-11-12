package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// var data []byte --> Refer notes; Try doing with this
	data := make([]byte, 99999)

	dataRead, _ := resp.Body.Read(data)

	fmt.Println(string(data)) //----> if we just write Println(data) it just prints a slice.
	fmt.Println(dataRead)     //------> no. of bytes read by read func
}
