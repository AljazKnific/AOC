package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("Input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()

		res += getNumber(FindMulExpression(line))

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Res: %d\n", res)

}

func FindMulExpression(line string) [][]string {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return mulRegex.FindAllStringSubmatch(line, len(line))
}

func getNumber(array [][]string) int {
	res := 0
	for _, el := range array {

		numOne, err := strconv.Atoi(el[1])

		if err != nil {
			panic(err)
		}

		numTwo, err := strconv.Atoi(el[2])

		if err != nil {
			panic(err)
		}

		res += numOne * numTwo
	}

	return res
}
