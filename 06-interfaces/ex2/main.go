package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWrite struct{}

func main() {
	//resp, err := http.Get("http://google.com")
	resp, err := http.Get("https://api.github.com/users/octocat")
	if err != nil {
		fmt.Printf("error %+v", err)
		os.Exit(1)
	}

	//create a byte slice to be populated with the info that read return
	//bs := make([]byte, 99999)

	//Read populate the byte slice
	//resp.Body.Read(bs)
	//print the byte slice
	//fmt.Println(bs)

	// parse the byteslice into string and will print a json
	//fmt.Println(string (bs))


	//one line
	//io.Copy(os.Stdout, resp.Body)

	lw := logWrite{}
	io.Copy(lw, resp.Body)

}

// Write Custome implementation of Write
func(logWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this bytes: ", len(bs))
	return len(bs), nil
}