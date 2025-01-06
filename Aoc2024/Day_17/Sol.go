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

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make(map[int]int)
	var out string
	var solTwo []int
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			counter = -1
			continue
		}
		tmp := strings.Split(line, " ")

		if counter > -1 {
			el, _ := strconv.Atoi(tmp[2])
			values[counter] = el
			counter++
			

		} else {
			ins := strings.Split(tmp[1], ",")

			for i := 0; i < len(ins); i++ {
				ele, _ :=  strconv.Atoi(ins[i])
				solTwo = append(solTwo, ele)

			}

			for i := 0; i < len(ins) - 1; i+=2 {

				operand, _ := strconv.Atoi(ins[i+1])

				switch ins[i] {
				case "0":
					values[0] = values[0] / int(math.Pow(2, float64(comboOperand(values, operand))))
					break
				case "1":
					values[1] = values[1] ^ operand
					break
				case "2":
					values[1] = comboOperand(values, operand) % 8
					break
				case "3":
					if values[0] == 0 {
						break
					}
					i = operand - 2
					break
				case "4":
					values[1] = values[1] ^ values[2]
					break
				case "5":
					t := strconv.Itoa(comboOperand(values, operand) % 8)
					out += t
					break
				case "6":
					values[1] = values[0] / int(math.Pow(2, float64(comboOperand(values, operand))))
					break
				case "7":
					values[2] = values[0] / int(math.Pow(2, float64(comboOperand(values, operand))))
					break
				default:
					fmt.Println("Unknown opcode")
				}
			}
		}
		
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tt :=  strings.Split(out, "")
	out = ""
	for i := 0; i < len(tt) - 1; i++ {
		out += tt[i] + ","
	}
	out += tt[len(tt)-1]

	fmt.Println(solTwo)
	//instructions
	// B = A % 8
	// B = B XOR 1
	// C = A / 2^B
	// A = A / 2^3
	// B = B XOR 4
	// B = B XOR C
	// out(B % 8)
	// a != 0 loop


	fmt.Println(out)
	fmt.Println(secSolution(solTwo, 0))
}

func comboOperand(values map[int]int, num int) int {
	if num < 4 {
		return num
	} else {
		return values[num - 4]
	}
}

func secSolution(pro []int, ans int) int {
	fmt.Println(pro)
	if len(pro) == 0 {
		return ans
	}

	for i := 0; i < 8; i++ {
		B := i
		//B = A % 8
		A := (ans << 3) + B
		B = B ^ 1
		C := A >> B
		A = A >> 3
		B = B ^ 4
		B = B ^ C
		fmt.Println(ans,A, B, C)
		if B % 8 == pro[len(pro)-1] {
			sub := secSolution(pro[:len(pro)-1], A)
			if sub == -42 {
				continue
			}
			return  sub

		}
	}

	return -42
}
