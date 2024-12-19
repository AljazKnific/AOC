package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var cache map[string] int

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var table []int
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		tmpTable := strings.Split(line, " ")

		for i := 0; i < len(tmpTable); i++ {
			x, err := strconv.Atoi(tmpTable[i])

			if err != nil {
				panic(err)
			}

			table = append(table, x)
		}
	}

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//res = firstSolution(table, 25)
	cache = make(map[string] int)
	res2 = solutionSecond(table, 75)

	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func firstSolution(table []int, blinks int) int {

	for i := 0; i < blinks; i++ {
		table = processBlinks(table)
	}


	return len(table)
}

// Helper to process blinks
func processBlinks(table []int) []int {
	loop := len(table)

	for j := 0; j < loop; j++ {
		el, updatedTable := table[0], table[1:]
		table = updatedTable

		if el == 0 {
			table = append(table, 1)
		} else {
			table = processElement(el, table)
		}
	}

	return table
}

// Helper to process individual elements
func processElement(el int, table []int) []int {
	s := strconv.Itoa(el)

	if len(s)%2 == 0 {
		res1, res2 := splitEvenDigits(s)

		num1, err := strconv.Atoi(res1)

		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(res2)

		if err != nil {
			panic(err)
		}

		table = append(table, num1)
		table = append(table, num2)

	} else {
		table = append(table, el*2024)
	}

	return table
}

// Helper to split a string with even digits
func splitEvenDigits(s string) (string, string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func solutionSecond(table []int, blinks int) int {
	res := 0
	for _, el := range table {
		res += rec2(blinks, el)
	}

	return res
}

func rec2(blinks int, num int) int {
	if blinks == 0 {
		return 1
	}

	key := strconv.Itoa(blinks) + " + " + strconv.Itoa(num)

	if cache[key] != 0 {
		return cache[key]
	}

	if num == 0 {
		return rec2(blinks -1, 1)
	} else {
		s := strconv.Itoa(num)

		if len(s) % 2 == 0 {
			res1, res2 := splitEvenDigits(s)

			num1, err := strconv.Atoi(res1)
	
			if err != nil {
				panic(err)
			}
	
			num2, err := strconv.Atoi(res2)
	
			if err != nil {
				panic(err)
			}

			return rec2(blinks - 1, num1) + rec2(blinks - 1, num2)

		} else {
			stones := rec2(blinks - 1, num * 2024)
			cache[key] = stones
			return stones
		}
	}
}




