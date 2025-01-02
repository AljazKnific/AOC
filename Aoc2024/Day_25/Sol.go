package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var keys [][]int
	var locks [][]int
	counter := 0

	lock := true
	tempInt := resetArray()

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, "")

		if len(strings.TrimSpace(line)) == 0 {
			if lock {
				locks = append(locks, tempInt)
			} else {
				keys = append(keys, tempInt)
			}

			tempInt = resetArray()
			counter = 0
			continue
		}

		if counter == 0 {
			if tmp[0] == "#" {
				lock = true
			} else {
				lock = false
			}
		}

		if (counter > 0 && lock) || (!lock && counter < 6) {
			for i := 0; i < 5; i++ {
				if tmp[i] == "#" {
					tempInt[i]++
				}
			}
		}

		counter++
	}

	if counter > 0 {
		if lock {
			locks = append(locks, tempInt)
		} else {
			keys = append(keys, tempInt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", solutionFirst(keys, locks))
}

func solutionFirst(keys [][]int, locks [][]int) int {
	res := 0
	for _,k:= range keys {
		for _, i := range locks {
			check := true
			for j := 0; j < len(k); j++ {
				if k[j] + i[j] > 5 {
					check = false
					break
				} 
			}

			if check {
				res += 1
			}
		}
	}

	return res
}

func resetArray() []int {
	return []int{0, 0, 0, 0, 0}
}
