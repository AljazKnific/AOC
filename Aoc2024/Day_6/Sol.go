package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pair struct {
	x int 
	y int
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var table []string

	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}

	x, y := findStart(table)
	if x == -42 && y == -42 {
		fmt.Println("No start position found")
		return
	}

	startDir := 0
	resArray := make([][]uint8, len(table))
	for i := range resArray {
		resArray[i] = make([]uint8, len(table[0]))
	}

	for {
		var check bool
		x, y, check = drawPath(startDir, table, x, y, resArray)

		if check {
			break
		}
		startDir++
		startDir %= 4
	}

	res := firstSolution(resArray)
	fmt.Printf("Result: %d\n", res)
	res2 := solutionTwo(resArray, table)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result2: %d\n", res2)
}

// drawPath moves in a given direction and returns updated coordinates and status.
func drawPath(direction int, table []string, startX int, startY int, resArray [][]uint8) (int, int, bool) {
	xdir := 0
	ydir := 0
	switch direction {
	// UP
	case 0:
		xdir = -1
	// RIGHT
	case 1:
		ydir = 1
	// DOWN
	case 2:
		xdir = 1
	// LEFT
	case 3:
		ydir = -1
	}

	for {
		startX += xdir
		startY += ydir

		// Check bounds
		if startX < 0 || startX >= len(table) || startY < 0 || startY >= len(table[0]) {
			return startX, startY, true
		}

		// Stop if hitting #
		if table[startX][startY] == '#' {
			return startX - xdir, startY - ydir , false
		}

		// Mark visited
		resArray[startX][startY] = 1
	}
}

// findStart locates the start position marked by '^'.
func findStart(table []string) (int, int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '^' {
				return i, j
			}
		}
	}
	return -42, -42
}

// firstSolution calculates the total count of visited positions.
func firstSolution(table [][]uint8) int {  
	res := 0
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

// Get indices of all visited positions
func findStart2(table [][]uint8) []pair {
	var res []pair

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 1 {
				res = append(res, pair{x: i,y: j})
			}
		}
	}
	return res
}

func drawPath2(direction uint8, table []string, startX int, startY int, resArray [][]uint8, avoidX int, avoidY int) (int, int, int) {
	xdir := 0
	ydir := 0
	switch direction {
	// UP
	case 1:
		xdir = -1
	// RIGHT
	case 2:
		ydir = 1
	// DOWN
	case 3:
		xdir = 1
	// LEFT
	case 4:
		ydir = -1
	}

	for {
		startX += xdir
		startY += ydir

		
		// Check bounds
		if startX < 0 || startX >= len(table) || startY < 0 || startY >= len(table[0]) {
			return startX, startY, 1
		}

		if direction == resArray[startX][startY] {
			return startX, startY, 2
		}

		// Stop if hitting # or obstacle
		if table[startX][startY] == '#' || (avoidX == startX && avoidY == startY) {
			return startX - xdir, startY - ydir , 0
		}

		// Mark visited
		resArray[startX][startY] = direction
	}
}

func solutionTwo(resArray [][]uint8, table []string) int {

	pairs := findStart2(resArray)
	res := 0

	mainX, mainY := findStart(table)
	x := mainX
	y := mainY

	for _,el := range pairs {
		var startDir uint8
		startDir = 1
		for i := range resArray {
			resArray[i] = make([]uint8, len(table[0]))
		}
		
	
		for {
			var check int
			x, y, check = drawPath2(startDir, table, x, y, resArray, el.x, el.y)
	
			if check == 1 {
				break
			} else if check == 2 {
				res++
				break
			}
			startDir++
			startDir %= 5
			if startDir == 0 {
				startDir++
			}
		}
		x = mainX
		y = mainY
	}

	return res
}