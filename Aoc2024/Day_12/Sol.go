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
}

type pair struct {
	x float
	y float
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

	res = solFirst(table)
	resetTable(table)
	res2 = solSecond()
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}


func solFirst(table [][]pairTable) int {
	res := 0
	var m []pairResult

	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if !table[i][j].checked {
				m = append(m, pairResult{fence: 0, numOfLetters: 0})
				rec(table, i, j, &m)
				
			}
		}
	}

	for _,el := range m {
		//fmt.Printf("Letter: %s, fence: %d, numOfLetters: %d, *: %d\n", v, el.fence, el.numOfLetters, el.fence * el.numOfLetters)
		res += el.fence * el.numOfLetters 
	}

	return res
}


func rec(table [][]pairTable, x int, y int, m *[]pairResult) {

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

	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter && !table[x - 1][y].checked) {
		rec(table, x - 1, y, m)
	}

	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter && !table[x][y - 1].checked) {
		rec(table, x, y - 1, m)
	}

	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter && !table[x + 1][y].checked) {
		rec(table, x + 1, y, m)
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter && !table[x][y + 1].checked) {
		rec(table, x, y + 1, m)
	}

}

func resetTable(table [][]pairTable) {
	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j].checked = false
		}
	} 
}

func solSecond(table [][]pairTable) int {
	res := 0

	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			var results := resultSecond{
				numOfLetters: 0,
				sol: []pair{}
			}
			if !table[i][j].checked {
				//call function
				rec2(table, i, j, &results)

				res += results.numOfLetters * len(results.sol)
			}
		}
	}


	return res
}

func rec2(table [][]pairTable, x int, y int, m *resultSecond) {
	if table[x][y].checked {
		return
	}

	var sides []int

	
	
	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter) {
		fences -= 1
		sides = append(sides, 0)
	}
	
	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter) {
		fences -= 1
		sides = append(sides, 2)
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter) {
		fences -= 1
		sides = append(sides, 1)
	}


	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter) {
		fences -= 1
		sides = append(sides, 3)

	}

	m.numOfLetters++

	table[x][y].checked = true
	switch len(sides) {
		case 4:
			for k: 0; k < 4; k++ {
				m.sol = append(m.sol, pair{x: x, y: y})
			}
			fmt.Println("Kocka je sama")
			return
		case 3: 
			addNumAndCheck(x - 0.5, y - 0.5, m)
			addNumAndCheck(x - 0.5, y + 0.5, m)
			addNumAndCheck(x + 0.5, y - 0.5, m)
			addNumAndCheck(x + 0.5, y + 0.5, m)
			break
		case 2:
			break
		case 1:
			break

		default: 
			fmt.Println("Prispel sem na default -> nimam sosedov :(")
			break

	}
	

	if(x - 1 > -1 && table[x][y].letter == table[x - 1][y].letter && !table[x - 1][y].checked) {
		rec2(table, x - 1, y, m)
	}

	if(y - 1 > -1 && table[x][y].letter == table[x][y - 1].letter && !table[x][y - 1].checked) {
		rec2(table, x, y - 1, m)
	}

	if(x + 1 < len(table) && table[x][y].letter == table[x + 1][y].letter && !table[x + 1][y].checked) {
		rec2(table, x + 1, y, m)
	}

	if(y + 1 < len(table) && table[x][y].letter == table[x][y + 1].letter && !table[x][y + 1].checked) {
		rec2(table, x, y + 1, m)
	}
}

func addNumAndCheck(num1 float, num2 float, m *resultSecond) {
		var check bool
		check = false
		for _,el := range m.sol {
			if el.x == num1 && el.y == num2 {
				check = true
				break
			}
		}

		if !check {
			temp := m.sol
			temp = append(temp, pair{x: num1, y: num2})
			m.sol = temp
		}
}

func checkCornerAndAdd(x int, y int, sides []int, m *resultSecond) {
	// sides are two

	//parallel
	if sides[0] - sides[1] == 2 || sides[1] - sides[0] == 2 {	
		
	//corner
	} else {
		
		if sides[0] == 0 {
			if sides[1] == 1 {
				addNumAndCheck(x + 0.5, y + 0.5, m)
			} else if sides[1] == 3 {
				addNumAndCheck(x + 0.5, y - 0.5, m)
			}
		} else if sides[0] == 2 {
			if sides[1] == 1 {
				addNumAndCheck(x - 0.5, y + 0.5, m)
			} else if sides[1] == 3 {
				addNumAndCheck(x - 0.5, y - 0.5, m)
			}
		}
	} 
}

