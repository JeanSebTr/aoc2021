package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input/day-2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var horizontal int64 = 0
	var depth int64 = 0
	var aim int64 = 0

	for _, line := range strings.Split(string(data), "\n") {
		var direction string
		var distance int64
		if _, err = fmt.Sscanf(line, "%s %d", &direction, &distance); err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += distance
			depth += aim * distance
			break
		case "up":
			aim -= distance
			break
		case "down":
			aim += distance
			break
		}
	}

	log.Printf("We moved a distance of %d units\n", depth*horizontal)
}
