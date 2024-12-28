package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	
)

type inputs struct {
	px int
	py int 
	vx int
	vy int
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

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`[-]?\d+`)

		matches := re.FindAllString(line, -1)
	
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[0])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[2])

		inpTmp := inputs{px: px, py: py, vx: vx, vy: vy}

		table = append(table, inpTmp)
			
	}


	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//res = firstSolution(table, 100, 103, 101)
	res2 = secondSolution(table, 103, 101)

	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func secondSolution(table []inputs, boundsX int, boundsY int) int {
	counter := 1

	for {

		for idx := range table {
			table[idx].px = outOfBounds(table[idx].px, table[idx].vx, boundsX)
			table[idx].py = outOfBounds(table[idx].py, table[idx].vy, boundsY)
		}

		if differentResults(table) {
			break
		}

		counter++
	}

	return counter
}

func differentResults(table []inputs) bool {

	for i := 0; i < len(table) - 1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i].px == table[j].px && table[i].py == table[j].py {
				return false
			}
		}
	}

	return true
}


func firstSolution(table []inputs, rounds int, boundsX int, boundsY int) int {
	for i := 0 ; i < rounds; i++ {
		for idx := range table {
			table[idx].px = outOfBounds(table[idx].px, table[idx].vx, boundsX)
			table[idx].py = outOfBounds(table[idx].py, table[idx].vy, boundsY)
		}
	}
	
	
	return findResult(table, boundsX, boundsY)
}

func outOfBounds(p int, v int, bounds int) int {
	num := p + v

	if num < 0 {
		return bounds + num
	}

	if num >= bounds {
		return num - bounds
	}

	return num
}
func findResult(table []inputs, boundsX int, boundsY int) int {
	var tmpTable []int
	tmpTable = append(tmpTable, 0,0,0,0)
	crossX := (boundsX - 1) / 2
	crossY := (boundsY - 1) / 2
	for _, el := range table {
		if el.px == crossX || el.py == crossY {
			continue
		}

		if el.px < crossX && el.py < crossY {
			tmpTable[0]++
		}

		if el.px < crossX && el.py > crossY {
			tmpTable[1]++
		}

		if el.px > crossX && el.py < crossY {
			tmpTable[2]++
		}

		if el.px > crossX && el.py > crossY {
			tmpTable[3]++
		}
	}

	return tmpTable[0] * tmpTable[1] * tmpTable[2] * tmpTable[3] 
}



