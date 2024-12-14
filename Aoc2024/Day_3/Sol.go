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
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res2 := 0
	use := true

	for scanner.Scan() {
		line := scanner.Text()

		tmp, newUse := getNumber2(
			use,
			FindMulExpression(line),
			FindDoExpression(line),
			FindDontExpression(line),
			FindMulExpressionPosition(line),
		)
		res2 += tmp
		use = newUse 
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Res: %d\n", res2)
}

func getNumber2(use bool, array [][]string, do [][]int, dont [][]int, pos [][]int) (int, bool) {
	res := 0
	for ind, el := range array {
		numOne, err := strconv.Atoi(el[1])
		if err != nil {
			panic(err)
		}

		numTwo, err := strconv.Atoi(el[2])
		if err != nil {
			panic(err)
		}

		position := pos[ind][0]
		doPos, dontPos := -1, -1

		for i := range do {
			if position > do[i][0] {
				doPos = i
			} else {
				break
			}
		}

		for i := range dont {
			if position > dont[i][0] {
				dontPos = i
			} else {
				break
			}
		}

		if doPos != -1 && dontPos != -1 {
			use = do[doPos][0] > dont[dontPos][0]
		} else if doPos != -1 {
			use = true
		} else if dontPos != -1 {
			use = false
		}


		if use {
			res += numOne * numTwo
		}
	}

	return res, use
}

func FindDoExpression(line string) [][]int {
	doRegex := regexp.MustCompile(`do\(\)`)
	return doRegex.FindAllStringIndex(line, -1)
}

func FindDontExpression(line string) [][]int {
	dontRegex := regexp.MustCompile(`don't\(\)`)
	return dontRegex.FindAllStringIndex(line, -1)
}

func FindMulExpressionPosition(line string) [][]int {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return mulRegex.FindAllStringIndex(line, -1)
}

func FindMulExpression(line string) [][]string {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return mulRegex.FindAllStringSubmatch(line, -1)
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
