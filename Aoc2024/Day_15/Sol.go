package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	
)

type pair struct {
	xl int
	yl int
	xr int
	yr int
}

type normalPair struct {
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

	var table [][]string
	var paths [][]string
	var table2 [][]string
	res := 0
	res2 := 0

	check := false

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, "")
		
		if len(tmp) == 0 {
			check = true
			continue
		}

		if !check {
			table = append(table, tmp)

		} else {
			paths = append(paths, tmp)
		}
			
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(table); i++ {
		var t []string
		for j := 0; j < len(table[i]); j++ {
			switch table[i][j] {
			case "#":
				t = append(t, "#", "#")
				break
			case "O":
				t = append(t, "[", "]")
				break
			case ".":
				t = append(t, ".", ".")
				break
			case "@":
				t = append(t, "@", ".")
				break
			default:
			}
		}
		table2 = append(table2, t)
	}

	res = firstSolution(paths, table)
	res2 = secondSolution(paths, table2)
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}


func secondSolution(paths [][]string, table [][]string) int {
	res := 0
	startX, startY := startingPosition(table)


	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			startX, startY = changePositionSecond(startX, startY, paths[i][j], table)
		}
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == "[" {
				res += (i * 100 + j)
			}
		}
	}


	return res
}


func changePositionSecond(x int, y int, ch string, table [][]string) (int, int) {
	numx, numy := 0, 0

	switch ch {
	case "^":
		numx = -1
		break
	case "<":
		numy = -1
		break
	case ">":
		numy = 1
		break
	case "v":
		numx = 1
		break
	default:
		fmt.Println("Unknown character in paths array")	
	}

	switch table[x + numx][y + numy] {

	case "#":
		return x, y
	case ".":
		times, lastX, lastY := checkForWalls(x,y,numx,numy, table)
		return moveTheRobot(times, lastX, lastY, numx, numy, table)
	default:

		//left and right covered
		if numy == 1 || numy == -1 {
			times, lastX, lastY := checkForWalls(x,y,numx,numy, table)

			if times == -42 {
				return x, y
			} else {
				return moveTheRobot(times, lastX, lastY, numx, numy, table)
			}

		// implement for up and down movements	
		} else {
			//var lastPairs []normalPair
			var pairs []pair
			pairs = append(pairs, pair{xl: x, yl: y, xr: -1, yr: -1})

			if getInfo(x + numx, y + numy, numx, numy, table, &pairs) {
				moveTheRobotTwo(pairs, numx, numy, table)
				return x + numx, y + numy
			}

		}

		return x,y

	}

	return -42, -42
}

func moveTheRobotTwo(pairs []pair, dirX int, dirY int, table [][]string) {
	
	for i:= len(pairs) - 1; i > -1; i-- {
		el := pairs[i]

		if el.xl != -1 || el.yl != -1 {
			table[el.xl + dirX][el.yl + dirY] =  table[el.xl][el.yl]
			table[el.xl][el.yl] = "."

		}

		if el.xr != -1 || el.yr != -1 {
			table[el.xr + dirX][el.yr + dirY] =  table[el.xr][el.yr]
			table[el.xr][el.yr] = "."

		}

	}
}

// create  a queue for a 
func getInfo(x int, y int, dirX int, dirY int, table [][]string, pairs *[]pair) bool {

	nX, nY := x, y
	var queue []pair

	if table[x][y] == "]" {
		nY = nY - 1
		queue = append(queue, pair{xl: nX, yl: nY, xr: x, yr: y })
	} else {
		nY = nY + 1
		queue = append(queue, pair{xl: x, yl: y, xr: nX, yr: nY })
	}


	for {

		var tmpEl pair
		tmpEl, queue = queue[0], queue[1:]

		leftX, leftY := tmpEl.xl, tmpEl.yl
		rightX, rightY := tmpEl.xr, tmpEl.yr
		cc := false
		for _, el := range *pairs {
			if el.xl == leftX && el.yl == leftY && el.xr == rightX && el.yr == rightY {
				cc = true
				break
			}
		}

		if cc {
			if len(queue) == 0 {
				break
			}
			continue
		}

		if table[leftX + dirX][leftY + dirY] == "#" || table[rightX + dirX][rightY + dirY] == "#" {
			return false
		}

		if table[leftX + dirX][leftY + dirY] == "." && table[rightX + dirX][rightY + dirY] == "." {	

			t3 := pair{xl: leftX, yl: leftY, xr: rightX, yr: rightY}
			*pairs = append(*pairs, t3)

			
		}

		if table[leftX + dirX][leftY + dirY] == "." && table[rightX + dirX][rightY + dirY] == "[" {	

			t3 := pair{xl: leftX, yl: leftY, xr: rightX, yr: rightY}
			*pairs = append(*pairs, t3)

			//add to queue
			t4 := pair{xl: leftX + dirX, yl: leftY + dirY + 1, xr: rightX + dirX, yr: rightY + dirY + 1}
			queue = append(queue, t4)
			
		}

		if table[leftX + dirX][leftY + dirY] == "]" && table[rightX + dirX][rightY + dirY] == "." {	
			t3 := pair{xl: leftX, yl: leftY, xr: rightX, yr: rightY}
			*pairs = append(*pairs, t3)

			//add to queue
			t4 := pair{xl: leftX + dirX, yl: leftY + dirY - 1, xr: rightX + dirX, yr: rightY + dirY - 1}
			queue = append(queue, t4)
			
		}

		if table[leftX + dirX][leftY + dirY] == "]" && table[rightX + dirX][rightY + dirY] == "[" {	

			t3 := pair{xl: leftX, yl: leftY, xr: rightX, yr: rightY}
			*pairs = append(*pairs, t3)

			//add to queue
			t4 := pair{xl: leftX + dirX, yl: leftY + dirY - 1, xr: rightX + dirX, yr: rightY + dirY - 1}
			t5 := pair{xl: leftX + dirX, yl: leftY + dirY + 1, xr: rightX + dirX, yr: rightY + dirY + 1}
			queue = append(queue, t4, t5)
			
		}

		// same [ -> [ or ] -> ]
		if table[leftX][leftY] == table[leftX + dirX][leftY + dirY] {
			t3 := pair{xl: leftX, yl: leftY, xr: rightX, yr: rightY}
			*pairs = append(*pairs, t3)

			t4 := pair{xl: leftX + dirX, yl: leftY + dirY, xr: rightX + dirX, yr: rightY + dirY}
			queue = append(queue, t4)
		}

		if len(queue) == 0 {
			break
		}
	}

	return true
}

func startingPosition(table [][]string) (int, int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == "@" {
				return i, j
			}
		}
	}

	return -1, -1
}

//Function returns (int, int, int) -> (how many chars need to be moved, x-pos of the last element, y-pos of the last element)
func checkForWalls(x int, y int, dirX int, dirY int, table [][]string) (int, int, int) {
	x, y = x + dirX, y + dirY
	times := 0
	for {
		if table[x][y] == "#" {
			return -42, -42, -42
		} 

		if table[x][y]  == "." {
			times++
			return times, x, y
		}

		times++
		x, y = x + dirX, y + dirY 
	}

	return -1, -1, -1
}

func moveTheRobot(times int, x int, y int, dirX int, dirY int, table [][]string) (int, int) {

	for _ = range times {
		tmp := table[x - dirX][y - dirY]
		table[x][y] = tmp
		table[x- dirX][y - dirY] = "."

		x -= dirX
		y -= dirY
	}

	return x + dirX, y + dirY
}

func changePosition(x int, y int, ch string, table [][]string) (int, int) {
	numx, numy := 0, 0

	switch ch {
	case "^":
		numx = -1
		break
	case "<":
		numy = -1
		break
	case ">":
		numy = 1
		break
	case "v":
		numx = 1
		break
	default:
		fmt.Println("Unknown character in paths array")	
	}

	switch table[x + numx][y + numy] {

	case "#":
		return x, y
	default:
		times, lastX, lastY := checkForWalls(x,y,numx,numy, table)

		if times == -42 {
			return x, y
		} else {
			return moveTheRobot(times, lastX, lastY, numx, numy, table)
		}
	}
}

func firstSolution(paths [][]string, table [][]string) int {
	res := 0
	startX, startY := startingPosition(table)


	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			startX, startY = changePosition(startX, startY, paths[i][j], table)
		}
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == "O" {
				res += (i * 100 + j)
			}
		}
	}

	return res
}
