package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int][]int)
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			x := strings.Split(line, "|")
			if i, err := strconv.Atoi(x[0]); err == nil {
				if j, err := strconv.Atoi(x[1]); err == nil {
					m[i] = append(m[i], j)
				}
			}



			
		} else if len(line) > 0 {
			temp := checkString(m, line) 
			if temp > 0 {
				res += temp
			} else {
				res2 += checkString2(m, line)
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)

}

//pass values of key and array before the number
func correct(values []int, before []int) bool {
	for _,v := range values {
		for _, v2 := range before {
			if v == v2 {
				return false
			}
		}
	}
	return true
}

//pass values of key and array before the number
func correct2(m map[int][]int, before []int, array []int, num int) {

	for j,v := range before {
		for _, v2 := range m[array[len(array) - 1]] {
			if v == v2 {
				tmp := array[j]
				array[j] = array[len(array) - 1]
				array[len(array) - 1] = tmp
				break
			}
		}
	}
}

func checkString(m map[int][]int, line string) int {
	tempArray := strings.Split(line, ",")
	var intArray []int
	for _,str := range tempArray {
		if j, err := strconv.Atoi(str); err == nil {
			intArray = append(intArray, j)
			if !correct(m[intArray[len(intArray) - 1]],intArray[:len(intArray) - 1]) {
				return 0
			}

		}
	}
	return intArray[(len(intArray) - 1) / 2]
}

func checkString2(m map[int][]int, line string) int {
	tempArray := strings.Split(line, ",")
	var intArray []int
	for _,str := range tempArray {
		if j, err := strconv.Atoi(str); err == nil {
			intArray = append(intArray, j)
			correct2(m,intArray[:len(intArray) - 1], intArray, intArray[len(intArray) - 1]) 

		}
	}
	return intArray[(len(intArray)) / 2]
}
