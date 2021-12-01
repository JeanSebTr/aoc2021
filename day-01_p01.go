package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var last int64

	var inc int = 0

	for i, line := range strings.Split(string(data), "\n") {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if i > 0 {
			if value > last {
				inc += 1
			}
		}
		last = value
	}

	log.Printf("There is %d measurements langer than the previous one\n", inc)
}
