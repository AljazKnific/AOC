package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type pairTable struct {
	letter string
	checked bool
}

type pairResult struct {
	numOfLetters int
	fence int
	pairs []pair
}

type pair struct {
	x float64
	y float64
}

type resultSecond struct {
	numOfLetters int
	sol []pair
}



func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	var table [][]pairTable
	


	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		tmpTable := strings.Split(line, "")

		tmpRow := []pairTable{}

		for _, char := range tmpTable {
			tmpRow = append(tmpRow, pairTable{letter: char, checked: false})
		}

		table = append(table, tmpRow)

	}

	//fmt.Println(table)

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var positions [][]pair

	res, positions = solFirst(table)
	fmt.Println(positions)
	resetTable(table)
	res2 = solSecond(table, positions)
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}


func solFirst(table [][]pairTable) (int, [][]pair) {
	res := 0
	var m []pairResult
	var positions [][]pair

	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if !table[i][j].checked {
				var ttt []pair
				m = append(m, pairResult{fence: 0, numOfLetters: 0})
				rec(table, i, j, &m, &ttt)
				positions = append(positions, ttt)
				
			}
		}
	}


	for _,el := range m {
		//fmt.Printf("Letter: %s, fence: %d, numOfLetters: %d, *: %d\n", v, el.fence, el.numOfLetters, el.fence * el.numOfLetters)
		res += el.fence * el.numOfLetters 
	}

	return res, positions
}


func rec(table [][]pairTable, x int, y int, m *[]pairResult, pairs *[]pair) {

	if table[x][y].checked {
		return
	}

	fences := 4
	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter) {
		fences -= 1

	}

	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter) {
		fences -= 1

	}

	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter) {
		fences -= 1
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter) {
		fences -= 1
	}

	tmp := (*m)[len((*m)) - 1]
	tmp.numOfLetters++
	tmp.fence += fences
	(*m)[len((*m)) - 1] = tmp

	table[x][y].checked = true
	*pairs = append(*pairs, pair{x: float64(x), y: float64(y)})

	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter && !table[x - 1][y].checked) {
		rec(table, x - 1, y, m, pairs)
	}

	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter && !table[x][y - 1].checked) {
		rec(table, x, y - 1, m, pairs)
	}

	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter && !table[x + 1][y].checked) {
		rec(table, x + 1, y, m, pairs)
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter && !table[x][y + 1].checked) {
		rec(table, x, y + 1, m, pairs)
	}

}

func resetTable(table [][]pairTable) {
	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j].checked = false
		}
	} 
}

func solSecond(table [][]pairTable, positions [][]pair) int {

	f, err := os.Create("output1.txt")

	if err != nil {
		fmt.Println(err)
		return -1
	}

	defer f.Close()

	res := 0

	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			var results = resultSecond{
				numOfLetters: 0,
				sol: []pair{},
			}
			if !table[i][j].checked {
				//call function
				for _,idx := range positions {
					if int(idx[0].x) == i && int(idx[0].y) == j {
						rec2(table, i, j, &results, idx)

					}
				}
				//fmt.Println(results.sol)
				//fmt.Printf("Letter: %s, Region size: %d, Sides: %d\n", table[i][j].letter, results.numOfLetters , len(results.sol))

				//write into a file

				for _, idx := range results.sol {
					_, err = f.WriteString(fmt.Sprintf("(%.1f, %.1f) ", idx.x, idx.y))

					if err != nil {
						panic(err)
					}
				}

				_, err = f.WriteString("\n")

				if err != nil {
					panic(err)
				}

				_, err = f.WriteString(fmt.Sprintf("Letter: %s, Region size: %d, Sides: %d\n", table[i][j].letter, results.numOfLetters , len(results.sol)))
				
				if err != nil {
					panic(err)
				}




				res += results.numOfLetters * len(results.sol)
			}
		}
	}


	return res
}

func rec2(table [][]pairTable, x int, y int, m *resultSecond, positions []pair) {
	if table[x][y].checked {
		return
	}

	var sides []int

	if(x + 1 >= len(table) || table[x][y].letter != table[x + 1][y].letter) {
		sides = append(sides, 0)
	}
	
	if(x - 1 < 0 || table[x][y].letter != table[x - 1][y].letter) {
		sides = append(sides, 2)
	}

	if(y + 1 >= len(table) || table[x][y].letter != table[x][y + 1].letter) {
		sides = append(sides, 1)
	}


	if(y - 1 < 0 || table[x][y].letter != table[x][y - 1].letter) {
		sides = append(sides, 3)

	}

	m.numOfLetters++

	table[x][y].checked = true
	//fmt.Printf("%s x: %d, y: %d, sides: %d\n", table[x][y].letter, x , y, len(sides))
	switch len(sides) {
	case 4:
		for k := 0; k < 4; k++ {
				m.sol = append(m.sol, pair{x: float64(x), y: float64(y)})
			}
			//fmt.Println("Kocka je sama")
			
			return
			case 3: 
			//fmt.Println(sides)
			var tmpSide []int
			tmpSide = append(tmpSide, 0, 1, 2, 3)
			for _, el := range sides {
				for x, v := range tmpSide {
					if el == v {
						tt := remove(tmpSide, x)
						tmpSide = tt
						break
					}
				}
			}
			//fmt.Printf("K:= %d\n", tmpSide[0])
			checkCornersThreeSides(x,y, tmpSide[0], table, m, positions)
			break
		case 2:
			checkCornerAndAdd(x, y, sides, table, m, positions)
			break
		case 1:
			//x int, y int, side int, table [][]pairTable, m *resultSecond
			checkCornerOneSide(x, y, sides[0], table, m, positions)
			break
			
			default: 
			checkCornersNoSides(x,y,table,m, positions)
			//fmt.Println("Prispel sem na default -> nimam sosedov :(")
			break
			
		}
		
		//fmt.Println(m.sol)

	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter && !table[x - 1][y].checked) {
		rec2(table, x - 1, y, m, positions)
	}

	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter && !table[x][y - 1].checked) {
		rec2(table, x, y - 1, m, positions)
	}

	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter && !table[x + 1][y].checked) {
		rec2(table, x + 1, y, m, positions)
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter && !table[x][y + 1].checked) {
		rec2(table, x, y + 1, m, positions)
	}
}

func addNumAndCheck(num1 float64, num2 float64, m *resultSecond, table [][]pairTable, positions []pair) {
		var check bool
		check = false
		for _,el := range m.sol {
			if el.x == num1 && el.y == num2 {
				if !checkCross(num1, num2, table, positions) {
					check = true
					break

				}
			}
		}

		if !check {
			temp := m.sol
			temp = append(temp, pair{x: num1, y: num2})
			m.sol = temp
		}
}

func checkCross(x float64, y float64, table [][]pairTable, positions []pair) bool {
	numNWx, numNWy := int(x - 0.5), int(y - 0.5)
	numNEx, numNEy := int(x - 0.5), int(y + 0.5)
	numSWx, numSWy := int(x + 0.5), int(y - 0.5)
	numSEx, numSEy := int(x + 0.5), int(y + 0.5)

	if numNWx < 0 || numNWy < 0  {
		return false
	}
	if numNEx < 0 || numNEy >= len(table)  {
		return false
	}
	if  numSWy < 0 || numSWx >= len(table) {
		return false
	}
	if  numSEx >= len(table) || numSEy >= len(table) {
		return false
	}

	//check the diagonal elements if they are in the same area
	if isInPositions(positions, numNWx, numNWy) && isInPositions(positions,numSEx, numSEy) && !isInPositions(positions, numNEx, numNEy) && !isInPositions(positions,numSWx, numSWy) && !table[numSEx][numSEy].checked {
		return true
	}

	//check the diagonal elements if they are in the same area
	if !isInPositions(positions, numNWx, numNWy) && !isInPositions(positions,numSEx, numSEy) && isInPositions(positions, numNEx, numNEy) && isInPositions(positions,numSWx, numSWy) && !table[numSWx][numSWy].checked {
		return true
	}

	/*
	if table[numNWx][numNWy].letter == table[numSEx][numSEy].letter && table[numNEx][numNEy].letter == table[numSWx][numSWy].letter && table[numNWx][numNWy].letter != table[numNEx][numNEy].letter {
		return true
	}
		*/

	return false
}

func isInPositions(positions []pair, x int, y int) bool {
	for _, k := range positions {
		if int(k.x) == x && int(k.y) == y {
			return true
		}
	}
	return false
}

func checkCornersIfTheSameLetter(x int, y int, letter string, table [][]pairTable) bool {
	if x < 0 || y < 0 || x >= len(table) || y >= len(table) {
		return false
	}

	if table[x][y].letter == letter {
		return true
	}

	return false

}

func checkCornerAndAdd(x int, y int, sides []int, table [][]pairTable, m *resultSecond, positions []pair) {
	// sides are two

	//parallel
	if sides[0] - sides[1] == 2 || sides[1] - sides[0] == 2 {
		//NW
		if checkCornersIfTheSameLetter(x - 1, y - 1, table[x][y].letter, table) {
			addNumAndCheck(float64(x) - 0.5, float64(y) - 0.5, m, table, positions)
		}

		//NE
		if checkCornersIfTheSameLetter(x - 1, y + 1, table[x][y].letter, table) {
			addNumAndCheck(float64(x) - 0.5, float64(y) + 0.5, m, table, positions)
		}

		// SW
		if checkCornersIfTheSameLetter(x + 1, y - 1, table[x][y].letter, table) {
			addNumAndCheck(float64(x) + 0.5, float64(y) - 0.5, m, table, positions)
		}

		//SE
		if checkCornersIfTheSameLetter(x + 1, y + 1, table[x][y].letter, table) {
			addNumAndCheck(float64(x) + 0.5, float64(y) + 0.5, m, table, positions)
		}
		
	//corner
	} else {
		
		if sides[0] == 0 {
			if sides[1] == 1 {
				addNumAndCheck(float64(x) + 0.5, float64(y) + 0.5, m, table, positions)
			} else if sides[1] == 3 {
				addNumAndCheck(float64(x) + 0.5, float64(y) - 0.5, m, table, positions)
			}
		} else if sides[0] == 2 {
			if sides[1] == 1 {
				addNumAndCheck(float64(x) - 0.5, float64(y) + 0.5, m, table, positions)
			} else if sides[1] == 3 {
				addNumAndCheck(float64(x) - 0.5, float64(y) - 0.5, m, table, positions)
			}
		}

			checkCornerOneSide(x,y,sides[0], table, m, positions)
			checkCornerOneSide(x,y,sides[1], table, m, positions)

	} 
}

func checkCornerOneSide(x int, y int, side int, table [][]pairTable, m *resultSecond, positions []pair) {
	numx1, numy1 := 0, 0
	numx2, numy2 := 0, 0

	switch side {
		case 0:
			numx1, numy1 = 1, 1
			numx2, numy2 = 1, -1
			break
		case 1:
			numx1, numy1 = -1, 1
			numx2, numy2 = 1, 1
			break
		case 2:
			numx1, numy1 = -1, 1
			numx2, numy2 = -1, -1
			break
		case 3:
			numx1, numy1 = -1, -1
			numx2, numy2 = 1, -1
			break
		default:
			fmt.Println("CheckCornerOneSide switch case default")
	}

	if checkCornersIfTheSameLetter(x + numx1, y + numy1, table[x][y].letter, table) {
		addNumAndCheck(float64(x) + (float64(numx1) / 2), float64(y) + (float64(numy1) / 2), m, table, positions)

	}
	if checkCornersIfTheSameLetter(x + numx2, y + numy2, table[x][y].letter, table) {
		addNumAndCheck(float64(x) + (float64(numx2) / 2), float64(y) + (float64(numy2) / 2), m, table, positions)

	}
}

func checkCornersThreeSides(x int, y int, side int, table [][]pairTable, m *resultSecond, positions []pair) {
	numx1, numy1 := 0.0, 0.0
	numx2, numy2 := 0.0, 0.0

	switch side {
		case 0:
			numx1, numy1 = -0.5, -0.5
			numx2, numy2 = -0.5, 0.5
			addNumAndCheck(float64(x) + numx1, float64(y) + numy1, m, table, positions)
			addNumAndCheck(float64(x) + numx2, float64(y) + numy2, m, table, positions)
			break
		case 1:
			numx1, numy1 = 0.5, -0.5
			numx2, numy2 = -0.5, -0.5
			addNumAndCheck(float64(x) + numx1, float64(y) + numy1, m, table, positions)
			addNumAndCheck(float64(x) + numx2, float64(y) + numy2, m, table, positions)
			break
		case 2: 
			numx1, numy1 = 0.5, -0.5
			numx2, numy2 = 0.5, 0.5
			addNumAndCheck(float64(x) + numx1, float64(y) + numy1, m, table, positions)
			addNumAndCheck(float64(x) + numx2, float64(y) + numy2, m, table, positions)
			break
		case 3:
			numx1, numy1 = 0.5, 0.5
			numx2, numy2 = -0.5, 0.5
			addNumAndCheck(float64(x) + numx1, float64(y) + numy1, m, table, positions)
			addNumAndCheck(float64(x) + numx2, float64(y) + numy2, m, table, positions)
			break
		default:
			fmt.Println("Check corners three sides -> no side detected")
			return
	}

	checkCornerOneSide(x, y, side, table, m, positions)
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


func checkCornersNoSides(x int, y int, table [][]pairTable, m *resultSecond, positions []pair) {

	if(table[x][y].letter != table[x - 1][y - 1].letter) {
		addNumAndCheck(float64(x) - 0.5, float64(y) - 0.5, m, table, positions)
	}

	if(table[x][y].letter != table[x + 1][y + 1].letter) {
		addNumAndCheck(float64(x) + 0.5, float64(y) + 0.5, m, table, positions)
	}

	if(table[x][y].letter != table[x - 1][y + 1].letter) {
		addNumAndCheck(float64(x) - 0.5, float64(y) + 0.5, m, table, positions)
	}

	if(table[x][y].letter != table[x + 1][y - 1].letter) {
		addNumAndCheck(float64(x) + 0.5, float64(y) - 0.5, m, table, positions)
	}
}
	