package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
)

type packet struct {
	fir string
	oper string
	sec string
	res string
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make(map[string]int)
	var operations []packet

	check := false

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " ")
		if len(line) == 0 {
			check = true
			continue
		}

		if !check {
			t := tmp[0]
			tt := strings.Split(t, ":")
			x, _ := strconv.Atoi(tmp[1])
			values[tt[0]] = x
		} else {
			operations = append(operations, packet{fir: tmp[0], oper: tmp[1], sec: tmp[2], res: tmp[4]})
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	res := firstSolution(operations, values)

	fmt.Printf("Result: %d\n", res)
}

func firstSolution(operations []packet, values map[string] int) int {
	temp := operations
	res := 0

	for {
		if len(temp) ==  0 {
			break
		}

		for i, el := range temp {
			v1, ok1 := values[el.fir]

			v2, ok2 := values[el.sec]

			if ok1 && ok2 {
				values[el.res] = equationRes(el.oper, v1, v2)
				temp = append(temp[:i], temp[i+1:]...)
				break
			}

		}
	}

	counter := 0
	for {
		str := "z"
		if counter < 10 {
			str += "0"
		}	

		strNum := strconv.Itoa(counter)

		str += strNum

		v, ok := values[str]

		if ok {
			res += v * int(math.Pow(2, float64(counter)))
		} else {
			break
		}

		counter++

	}

	return res
}

func equationRes(op string, num1 int, num2 int) int {
	res := 0
	switch op {
	case "AND":
		if num1 == 1 && num2 == 1 {
			res = 1
		}
		break
	case "OR":
		if num1 == 1 || num2 == 1 {
			res = 1
		}
		break
	case "XOR":
		if num1 != num2 {
			res = 1
		}
		break
	default:
		fmt.Println("Equation res unknown operation")
		
	}

	return res
}