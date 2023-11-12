package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// // var data []byte --> Refer notes; Try doing with this
	// data := make([]byte, 99999)
	// dataRead, _ := resp.Body.Read(data)
	// fmt.Println(string(data)) //----> if we just write Println(data) it just prints a slice.
	// fmt.Println(dataRead)     //------> no. of bytes read by read func
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
