package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}

	// elves := map[int]int{}
	reader := bufio.NewReader(f)
	i := 0 


	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		_, ok := 
		if len(line) == 0 {

		}
		_, ok := elves[i]
		if !ok {
			elves[i] = calories
		}
		elves[i] += calories
		calories, _ := strconv.Atoi(string(line))

		fmt.Println(calories)

	}
}
