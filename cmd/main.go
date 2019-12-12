package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	fmt.Println("building site at test-blog/")

	f, err := os.Open("./test-blog/test.md")
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	output := blackfriday.Run(buf)
	fmt.Printf("\n\n%s\n", output)
}
