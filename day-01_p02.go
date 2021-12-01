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

	var samples []int64 = make([]int64, 4)

	var inc int = 0

	for i, line := range strings.Split(string(data), "\n") {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		samples[i%len(samples)] = value

		if i >= 3 {
			current := sumAt(samples, i, 3)
			last := sumAt(samples, i-1, 3)
			if current > last {
				inc += 1
			}
		}
	}

	log.Printf("There is %d measurements langer than the previous one\n", inc)
}

func sumAt(samples []int64, pos int, window int) int64 {
	var sum int64 = 0
	for i := 1; i <= window; i++ {
		sample_pos := (pos - window + i) % len(samples)
		sum += samples[sample_pos]
	}
	return sum
}
