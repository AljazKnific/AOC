package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		chars := strings.Split(line, " ")

		if correctSequence(chars) {
			res++
			continue
		}

		for i := 0; i <= len(chars)-1; i++ {
			x := RemoveIndex(chars, i)
			if correctSequence(x) {
				res++
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Result: %d\n", res)

}

func correctSequence(seq []string) bool {
	highLow := false

	fir, err := strconv.Atoi(seq[0])

	if err != nil {
		fmt.Println("error")
	}

	sec, err := strconv.Atoi(seq[1])

	if err != nil {
		fmt.Println("error")
	}

	if fir > sec {
		highLow = true
	}

	for i := 0; i < len(seq)-1; i++ {

		a, err := strconv.Atoi(seq[i])

		if err != nil {
			fmt.Println("Error at solving the first number")
		}

		b, err := strconv.Atoi(seq[i+1])

		if err != nil {
			fmt.Println("Error at solving the second number")
		}
		if highLow {
			tmp := a - b
			if tmp < 1 || tmp > 3 {
				return false
			}

		} else {

			tmp := b - a

			if tmp < 1 || tmp > 3 {
				return false
			}
		}

		if i == len(seq)-2 {
			return true
		}
	}

	return false
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
