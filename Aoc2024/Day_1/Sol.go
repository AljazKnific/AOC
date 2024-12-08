package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Section 1
	file, err := os.Open("Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstArray := []int{}
	secondArray := []int{}

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "   ")

		fir, err := strconv.Atoi(chars[0])

		if err != nil {
			fmt.Println("Error at the first number")
		}

		firstArray = append(firstArray, fir)

		sec, err := strconv.Atoi(chars[1])

		if err != nil {
			fmt.Println("Error at the second number")
		}

		secondArray = append(secondArray, sec)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	sort.Slice(firstArray, func(i, j int) bool {
		return firstArray[i] < firstArray[j]
	})

	sort.Slice(secondArray, func(i, j int) bool {
		return secondArray[i] < secondArray[j]
	})

	for i := 0; i < len(secondArray); i++ {
		tmp := secondArray[i] - firstArray[i]

		if tmp < 0 {
			tmp *= -1
		}

		res += tmp
	}

	res2 := 0

	for _, element := range firstArray {
		num := 0
		for i := 0; i < len(secondArray); i++ {
			if secondArray[i] == element {
				num++
			}
		}
		res2 += (element * num)
	}

	fmt.Printf("Result: %d", res)
	fmt.Printf("Result: %d", res2)
}
