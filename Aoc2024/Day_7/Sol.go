package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
)

type pair struct {
	num int
	equation []int
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var table []pair
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " ")

		var tempArray []int

		for i := 1; i < len(tmp); i++ {
			if j, err := strconv.Atoi(tmp[i]); err == nil {
				tempArray = append(tempArray, j)
				
			}
		}
		x := strings.Split(tmp[0], ":")
		if num, err := strconv.Atoi(x[0]); err == nil {
			table = append(table, pair{num: num, equation: tempArray})
			
		}


	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}

	res = solutionFirst(table)
	res2 = solutionSecond(table)

	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)

}


func createEquations(size int) [][]string {
	// Initialize the result array
	resArray := make([][]string, int(math.Pow(2, float64(size))))
	for i := range resArray {
		resArray[i] = make([]string, size)
	}

	// Calculate the number of equations
	numOfEquations := int(math.Pow(2, float64(size)))

	for i := 0; i < size; i++ {
		yPos := i
		xPos := 0

		x := int(math.Pow(2, float64(i)))

		for k := 0; k < x; k++ {
			// Fill with "*"
			for j := 0; j < (numOfEquations / (2 * x)); j++ {
				resArray[xPos][yPos] = "*"
				xPos++
			}

			// Fill with "+"
			for j := 0; j < (numOfEquations / (2 * x)); j++ {
				resArray[xPos][yPos] = "+"
				xPos++
			}
		}
	}

	return resArray
}

func createEquations2(size int) [][]string {
	// Initialize the result array
	resArray := make([][]string, int(math.Pow(3, float64(size))))
	for i := range resArray {
		resArray[i] = make([]string, size)
	}

	// Calculate the number of equations
	numOfEquations := int(math.Pow(3, float64(size)))

	for i := 0; i < size; i++ {
		yPos := i
		xPos := 0

		x := int(math.Pow(3, float64(i)))

		for k := 0; k < x; k++ {
			// Fill with "*"
			for j := 0; j < (numOfEquations / (3 * x)); j++ {
				resArray[xPos][yPos] = "*"
				xPos++
			}

			// Fill with "+"
			for j := 0; j < (numOfEquations / (3 * x)); j++ {
				resArray[xPos][yPos] = "+"
				xPos++
			}

			// Fill with "|"
			for j := 0; j < (numOfEquations / (3 * x)); j++ {
				resArray[xPos][yPos] = "|"
				xPos++
			}
		}
	}

		
	return resArray
}

func checkTheEquation(p pair, equation [][]string) int {
	for _, el := range equation {
		answ := p.equation[0]

		for i := 1; i < len(p.equation); i++ {
			if el[i-1] == "*" {
				answ *= p.equation[i]
			} else {
				answ += p.equation[i]
			}
		} 

		if answ == p.num {
			return p.num
		}
	}
	return 0
}

func checkTheEquation2(p pair, equation [][]string) int {
	for _, el := range equation {
		answ := p.equation[0]

		for i := 1; i < len(p.equation); i++ {
			if el[i-1] == "*" {
				answ *= p.equation[i]
			} else if el[i-1] == "+" {
				answ += p.equation[i]
			} else {
				concatStr := strconv.Itoa(answ) + strconv.Itoa(p.equation[i])
				tmp, err := strconv.Atoi(concatStr)
				answ = tmp
				if err != nil {
					fmt.Println("Error converting to int:", err)
					return -42
				}
			}
		} 
		
		if answ == p.num {
			return p.num
		}
	}
	return 0
}

func solutionFirst(input []pair) int {
	res := 0
	for _, el := range input {
		var tableEq [][]string
		tableEq = createEquations(len(el.equation) - 1)
		res += checkTheEquation(el, tableEq)
	}

	return res

}

func solutionSecond(input []pair) int {
	res := 0
	for _, el := range input {
		var tableEq [][]string
		tableEq = createEquations2(len(el.equation) - 1)
		res += checkTheEquation2(el, tableEq)
	}

	return res

}