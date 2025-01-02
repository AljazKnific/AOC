package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const MOD = 16777216

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	x :=  make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		var t []int
		for _ = range 2000 {
			t = append(t, num % 10)
			num = step(num)
		}

		seen := make(map[string]bool)
		for i := 0; i < len(t) - 4; i++ {
			nums := t[i:i + 5]
			a, b, c, d := nums[1] - nums[0], nums[2] - nums[1], nums[3] - nums[2], nums[4] - nums[3]
			str := fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)
			_, ok := seen[str]
			if ok {
				continue
			}
			seen[str] = true
			_, ok2 := x[str]
			if !ok2{
				x[str] = 0
			}
			x[str] += nums[4]
		}

		res += num
		
	}

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", getResult(x))
}

func step(num int) int {
	num = (num ^ (num * 64)) % MOD
	num = (num ^ (num / 32)) % MOD
	num = (num ^ (num * 2048)) % MOD
	return num
}

func getResult(x map[string]int) int {
	res := 0

	for _,v := range x {
		if v > res {
			res = v
		}
	}
	return res
}
