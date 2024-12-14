package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	maxX := len(table) - 1
	maxY := len(table[0]) - 1
	
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if right(i,j,maxX, maxY, table) {
				res ++
			}
			if left(i,j,maxX, maxY, table) {
				res ++
			}
			
			if up(i,j,maxX, maxY, table) {
				res ++
			}
			if down(i,j,maxX, maxY, table) {
				res ++
			}
			if NE(i,j,maxX, maxY, table) {
				res ++
			}
			if NW(i,j,maxX, maxY, table) {
				res ++
				}
				
			if SW(i,j,maxX, maxY, table) {
				res ++
			}
				
			if SE(i,j,maxX, maxY, table) {
				res ++
				}

			if table[i][j] == 65 && MAS(i,j,maxX, maxY, table) {
				res2++
			}
					
				
		}
	}
		

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)

}


func right(x int, y int, maxx int, maxy int, field []string) bool {
	if(y + 3 > maxy) {
		return false
		}
		
		return field[x][y] == 88 && field[x][y+1] == 77 && field[x][y+2] == 65 && field[x][y + 3] == 83
		}
		

func left(x int, y int, maxx int, maxy int, field []string) bool {
	if(y - 3 < 0) {
		return false
	}

	return field[x][y] == 88 && field[x][y-1] == 77 && field[x][y-2] ==65 && field[x][y - 3] == 83
}

func up(x int, y int, maxx int, maxy int, field []string) bool {
	if(x - 3 < 0) {
		return false
	}

	return field[x][y] == 88 && field[x-1][y] == 77 && field[x-2][y] ==65 && field[x-3][y] == 83
}

func down(x int, y int, maxx int, maxy int, field []string) bool {
	if(x + 3 > maxx) {
		return false
	}

	return field[x][y] == 88 && field[x+1][y] == 77 && field[x+2][y] ==65 && field[x+3][y] == 83
}

func NE(x int, y int, maxx int, maxy int, field []string) bool {
	if(y + 3 > maxy || x - 3 < 0) {
		return false
		}

	return field[x][y] == 88 && field[x-1][y + 1] == 77 && field[x-2][y + 2] ==65 && field[x-3][y + 3] == 83
}

func NW(x int, y int, maxx int, maxy int, field []string) bool {
	if(x - 3 < 0 || y - 3 < 0) {
		return false
	}

	return field[x][y] == 88 && field[x-1][y - 1] == 77 && field[x-2][y - 2] ==65 && field[x-3][y - 3] == 83
}

func SE(x int, y int, maxx int, maxy int, field []string) bool {
	if(x + 3 > maxx || y + 3 > maxy) {
		return false
	}

	return field[x][y] == 88 && field[x+1][y + 1] == 77 && field[x+2][y + 2] ==65 && field[x+3][y + 3] == 83
}

func SW(x int, y int, maxx int, maxy int, field []string) bool {
	if(x + 3 > maxx || y - 3 < 0) {
		return false
	}

	return field[x][y] == 88 && field[x+1][y - 1] == 77 && field[x+2][y - 2] ==65 && field[x+3][y - 3] == 83
}

func MAS(x int, y int, maxx int, maxy int, field []string) bool {
	if(x + 1 > maxx || y - 1 < 0 || y + 1 > maxy || x - 1 < 0) {
		return false
	}

	numOfMas := 0

	if field[x-1][y-1] == 77 && field[x+1][y+1] == 83 {
		numOfMas++
	}

	if field[x-1][y + 1] == 77 && field[x+1][y-1] == 83 {
		numOfMas++
	}

	if field[x+1][y-1] == 77 && field[x-1][y+1] == 83 {
		numOfMas++
	}

	if field[x+1][y+1] == 77 && field[x-1][y-1] == 83 {
		numOfMas++
	}

	if numOfMas == 2 {
		return true
	}

	return false
}
