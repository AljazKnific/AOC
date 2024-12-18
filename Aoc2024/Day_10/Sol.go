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
	var starts []pair
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)

	}

	resArray := make([][]int, len(table))
	for i := range resArray {
		resArray[i] = make([]int, len(table[0]))
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			resArray[i][j] = int(table[i][j] - '0')
			
			if table[i][j]  == '0' {
				starts = append(starts, pair{x: i, y: j})
			}

		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res = solutionFirst(resArray, starts)
	res2 = solutionSecond(resArray, starts)

	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func solutionFirst(table [][]int, starts []pair) int {
	res := 0
	for _, v := range starts {
		var finishes []pair
		res += rec(table, v.x, v.y, &finishes)
	}

	return res

}

func rec(table [][]int, x int, y int, finishes *[]pair) int {
	res := 0

	// break condition
	if table[x][y] == 9 {
		for _, el := range *finishes {
			if el.x == x && el.y == y {
				return 0
			}
		}
		*finishes = append(*finishes, pair{x: x, y: y})
		return 1
	}


	// check all directions
	//UP
	if x - 1 >= 0 && table[x - 1][y] - table[x][y] == 1 {
		res += rec(table, x-1, y, finishes)
	}

	//LEFT
	if y - 1 >= 0 && table[x][y - 1] - table[x][y] == 1 {
		res += rec(table, x, y - 1, finishes)
	}

	//RIGHT
	if y + 1 < len(table) && table[x][y + 1] - table[x][y] == 1 {
		res += rec(table, x, y + 1, finishes)
	}

	//DOWN
	if x + 1 < len(table) && table[x + 1][y] - table[x][y] == 1 {
		res += rec(table, x+1, y, finishes)
	}

	return res
}

func solutionSecond(table [][]int, starts []pair) int {
	res := 0
	for _, v := range starts {
		res += recSecond(table, v.x, v.y)
	}

	return res

}

func recSecond(table [][]int, x int, y int) int {
	res := 0

	// break condition
	if table[x][y] == 9 {
		return 1
	}

	// check all directions
	//UP
	if x - 1 >= 0 && table[x - 1][y] - table[x][y] == 1 {
		res += recSecond(table, x-1, y)
	}

	//LEFT
	if y - 1 >= 0 && table[x][y - 1] - table[x][y] == 1 {
		res += recSecond(table, x, y - 1)
	}

	//RIGHT
	if y + 1 < len(table) && table[x][y + 1] - table[x][y] == 1 {
		res += recSecond(table, x, y + 1)
	}

	//DOWN
	if x + 1 < len(table) && table[x + 1][y] - table[x][y] == 1 {
		res += recSecond(table, x+1, y)
	}

	return res
}



