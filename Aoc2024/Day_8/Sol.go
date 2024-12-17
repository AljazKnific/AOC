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
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}

	
	var m map[string][]pair
	m = getPositions(table)
	
	resArray := make([][][]string, len(table))
	for i := range resArray {
		resArray[i] = make([][]string, len(table[0]))
		for j := range resArray[i] {
			resArray[i][j] = make([]string, 0)
		}
		
	}

	res = solutionFirst(m, resArray)
	res2 = solutionSecond(m, resArray, table)
	

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func getPositions(table []string) map[string][]pair {
	m := make(map[string][]pair)

	for i:= 0; i< len(table); i++ {
		for j:= 0; j < len(table[i]); j++ {
			if table[i][j] != '.' {
				key := string(table[i][j])
				m[key] = append(m[key], pair{x: i, y: j})
			} 
		}
	}
	return m
}

func checkPosition(resArray [][][]string, x int, y int, letter string) int {
	if x < 0 || y < 0 || x >= len(resArray) || y >= len(resArray) {
		return 0
	}

	var check bool
	var check1 bool
	check = false
	check1 = false
	if len(resArray[x][y]) > 0 {
		check = true
	}
	

	for _, el := range resArray[x][y] {
		if el == letter {
			check1 = true
			break
		}
	}

	if !check1 {
		resArray[x][y] = append(resArray[x][y], letter)
	} 

	if check {
		return 0
	}
	return 1
}

func checkPosition2(x1 int, y1 int, x2 int, y2 int, resArray [][][]string, letter string) {
	//position of new antinode
	x := 2 * x1 - x2
	y := 2 * y1 - y2

	if x < 0 || y < 0 || x >= len(resArray) || y >= len(resArray) {
		return 
	}

	var check1 bool
	check1 = false


	for _, el := range resArray[x][y] {
		if el == letter {
			check1 = true
			break
		}
	}

	if !check1 {
		resArray[x][y] = append(resArray[x][y], letter)
	}
	
	checkPosition2(x,y,x1,y1,resArray,letter)

}

func getResult2(resArray [][][]string, table []string) int {
	res := 0
	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(resArray[i]); j++ {
			if table[i][j] != '.' || len(resArray[i][j]) > 0 {
				res++
			}
		}
	}
	return res
}

func solutionFirst(m map[string][]pair, resArray [][][]string) int {

	res := 0

	for v, el := range m {
		for i := 0; i < len(el) - 1; i++ {
			for j:= i + 1; j < len(el); j++ {

				x1, y1 := el[i].x, el[i].y
				x2, y2 := el[j].x, el[j].y

				r1x := 2 * x1 - x2
				r1y := 2 * y1 - y2

				r2x := 2 * x2 - x1
				r2y := 2 * y2 - y1

				res += checkPosition(resArray, r1x, r1y, v)
				res += checkPosition(resArray, r2x, r2y, v)
			}
		}
	}

	return res
}

func solutionSecond(m map[string][]pair, resArray [][][]string, table []string) int {

	for v, el := range m {
		for i := 0; i < len(el) - 1; i++ {
			for j:= i + 1; j < len(el); j++ {

				x1, y1 := el[i].x, el[i].y
				x2, y2 := el[j].x, el[j].y

				checkPosition2(x1,y1,x2,y2,resArray,v)
				checkPosition2(x2,y2,x1,y1,resArray,v)
			}
		}
	}

	return getResult2(resArray, table)

}

