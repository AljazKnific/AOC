package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	var table []int
	var table2 []int
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		tmpTable := strings.Split(line, "")

		check := true
		num := 0
		for _, el := range tmpTable {
			tmp, err := strconv.Atoi(el)

			
			if err != nil {
				panic(err)
			}
			//add numbers to array
			if check {
				for i := 0; i < tmp; i++ {
					table2 = append(table2, num)
				}
				num++
				check = false
			} else {
				for i := 0; i < tmp; i++ {
					table2 = append(table2, -1)
				}
				check = true
			}
			table = append(table, tmp)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//res = solFirstAttempt2(table2)
	res2 = solSecond(table, table2)
	
	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}


func solFirstAttempt2(table []int) int {
	res := 0

	posSpace := 0
	lastDisk := len(table) - 1

	for {

		//get position of the firstFreeSpace
		for {
			if table[posSpace] == -1 {
				break
			}
			posSpace++
		}

		if lastDisk < posSpace {
			break
		}

		table[posSpace] = table[lastDisk]
		table[lastDisk] = -1

		for {
			if table[lastDisk] != -1 {
				break
			}
			lastDisk--
		}

	}


	for i := 0; i < len(table); i++ {
		if table[i] == -1 {
			continue
		}
		res += i * table[i]
	}

	return res
}

func solSecond(table []int, table2 []int) int {
	res := 0

	var indexTable []int
	for _, el := range table {
		indexTable = append(indexTable, el)
	}

	for i := len(table) - 1; i > 1; i -= 2 {
		files := table[i]

		// check all of the space
		for j := 1; j < i; j += 2 {
			//space for files
			if files <= indexTable[j] {
				//rewrite
				startIndex := 0
				for k:= 0; k < j; k++ {
					startIndex += table[k]
				}
				startIndex += (table[j] - indexTable[j])

				for k:= 0; k < files; k++ {
					table2[startIndex + k] = i / 2
				}
				indexTable[j] -= files

				//delete
				lastIndex := 0
				for k := len(table) - 1; k >= i; k-- {
					lastIndex += table[k]
				}

				start := len(table2) - lastIndex
				for k:= start; k < start + files; k++ {
					table2[k] = -1
				}
				break
			}
		}
	}

	//get the result
	for i := 0; i < len(table2); i++ {
		if table2[i] == -1 {
			continue
		}
		res += i * table2[i]
	}

	return res
}