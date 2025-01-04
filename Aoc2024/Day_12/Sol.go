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
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var positions [][]pair

	res, positions = solFirst(table)
	res2 = solutionSecond(table, positions)
	
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


func solutionSecond(table [][]pairTable, positions [][]pair) int {
	res := 0
	var dirTable []pair
	var dirX []pair
	var dirY []pair
	dirTable = append(dirTable, pair{x: -1, y: 0}, pair{x: 0, y: 1},pair{x: 1, y: 0},pair{x: 0, y: -1})
	dirX = append(dirX, pair{x: -1, y: 0}, pair{x: 1, y: 0})
	dirY = append(dirY, pair{x: 0, y: 1}, pair{x: 0, y: -1})

	//check first
	for _, el := range positions {
		edgeCounter := 0
		for i := 0; i < len(el); i++ {
			currElement := el[i]

			// convex edge
			for k := 0; k < 4; k++ {
				fir := dirTable[k]
				sec := dirTable[(k + 1) % 4]

				if !inSameArea(el, fir.x + currElement.x, fir.y + currElement.y) && !inSameArea(el, sec.x + currElement.x, sec.y + currElement.y) {
					edgeCounter++
				}

			}

			//concav edge
			for _, x1 := range dirX {
				for _, y1  := range dirY {

					if inSameArea(el, x1.x + currElement.x, currElement.y) && inSameArea(el, currElement.x, y1.y + currElement.y) && !inSameArea(el, x1.x + currElement.x, y1.y + currElement.y) {
						edgeCounter++
					}

				}
			}

		}
		res += len(el) * edgeCounter
	}


	return res
}

func inSameArea(pos []pair, x float64, y float64) bool {
	for _, el := range pos {
		if el.x == x && el.y == y {
			return true
		}
	}
	return false
}