package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var inputFile = flag.String("inputFile", "6thDecember/input.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := bytes[:len(bytes)-1]
	// First Part
	fmt.Println(check(contents, 4))
	// Second Part
	fmt.Println(check(contents, 14))
}

func check(contents []byte, num int) int {
	for i := 0; i < len(contents)-num; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < num; j++ {
			seen[contents[i+j]] = true
		}
		if len(seen) == num {
			return i + num
		}
	}
	return -1
}
