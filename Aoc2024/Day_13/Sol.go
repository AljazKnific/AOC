package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"math"
)

type inputs struct {
	ax float64
	ay float64
	bx float64
	by float64
	px float64
	py float64
}


func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var table []inputs
	res := 0
	res2 := 0

	counter := 0
	var tmpInput inputs
	for scanner.Scan() {
		if counter == 3 {
			table = append(table, tmpInput)
			counter = 0
			continue
		}
		line := scanner.Text()
		re := regexp.MustCompile(`[-]?\d+`)

		matches := re.FindAllString(line, -1)
	
		var numbers []float64
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err == nil {
				numbers = append(numbers, float64(num))
			}
		}
		if counter == 0 {
			tmpInput = inputs{ax: numbers[0], ay: numbers[1]}
		} else if counter == 1 {
			tmpInput.bx = numbers[0]
			tmpInput.by = numbers[1]
		} else if (counter == 2) {
			tmpInput.px = numbers[0] + 10000000000000
			tmpInput.py = numbers[1] + 10000000000000
		}

		counter++
	}
	table = append(table, tmpInput)

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//res = naiveFirstSolution(table)
	res2 = smartSolution(table)
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func naiveFirstSolution(input []inputs) int {
	res := 0
	for _, item := range input {
		minRez := 0
		for i := 1; i < 101; i++ {
			for j := 1; j < 101; j++ {
				//fmt.Println(float64(i) * item.ax + float64(j) * item.bx == item.px)
				if float64(float64(i) * item.ax + float64(j) * item.bx) == item.px && float64(float64(i) * item.ay + float64(j) * item.by) == item.py {
					num := 3 * i + j
					if minRez > num || minRez == 0 {
						minRez = num
					}
				}
			}
		}

		res += minRez
	}	

	return res
}

//Algebraic approach, where we search for the intersection of two lines
func smartSolution(input []inputs) int {
	res := 0

	for _, item := range input {
		//s -> A button
		//t -> B button
		var s float64
		var t float64
		s = float64(item.px * item.by - item.py * item.bx) / float64(item.ax * item.by - item.ay * item.bx)
		t = float64(item.px - s * item.ax) / float64(item.bx)
		
		if s == math.Trunc(s) && t == math.Trunc(t) {
			//fmt.Printf("S: %d, T: %d\n", s,t)
			res += int(s * 3 + t)
		}
			
	}

	return res
}




