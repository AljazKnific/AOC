package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	
)

type pair struct {
	x int
	y int
}

type pathInfo struct {
	//visited []vis
	pos pair
	cost int
	dir int
}

type vis struct {
	pos pair
	dir int
}

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var table [][]string
	res := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, "")
		
		table = append(table, tmp)
	
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Sx, Sy := findString(table, "S")
	Ex, Ey := findString(table, "E")

	S := pair{x: Sx, y: Sy}
	fmt.Println(S)
	E := pair{x: Ex, y: Ey}
	fmt.Println(E)
	res = firstSolution(table, S, E)

	fmt.Printf("Result: %d\n", res)
	fmt.Printf("Result2: %d\n", res2)
}

func findString(table [][]string, s string) (int, int) {
	for i := 0; i < len(table); i++ {
		for j:= 0; j < len(table[i]); j++ {
			if table[i][j] == s {
				return i, j
			}
		}
	}
	return -42, -42
}

func firstSolution(table [][]string, S pair, E pair) int {
	var queue []pathInfo
	var visited []vis

	t := pathInfo{pos: S, cost: 0, dir: 1}
	queue = append(queue, t)

	check := false
	counter := 0
	max := 0
	
	for {

		
		var curr pathInfo
		
		curr, queue = queue[0], queue[1:]
		

		if check {
			if curr.cost > max {
				fmt.Println(counter)
				//fmt.Println(queue)
				return max
			}

			if curr.pos.x == E.x && curr.pos.y == E.y && max == curr.cost {
				counter++
			}

		}

		if curr.pos.x == E.x && curr.pos.y == E.y {
			check = true
			counter++
			max = curr.cost
			//return curr.cost
		}

		//check if current position and dir is in visited
		if inVisited(curr, visited) {
			continue
		}

		visited = append(visited, vis{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: curr.dir})

		dirX, dirY := getDirection(curr.dir)		

		//can we go forward
		if checkIfExists(curr, dirX, dirY, visited) && table[curr.pos.x + dirX][curr.pos.y + dirY] != "#" {
			
			//add current position to visited, update with the new one 
			//tmpV := curr.visited
			//tmpV = append(tmpV, vis{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: curr.dir})
			
			//insert into the right position based on the cost of the path
			//newPathInfo := pathInfo{visited: tmpV, pos: pair{x: curr.pos.x + dirX, y: curr.pos.y + dirY}, dir: curr.dir, cost: curr.cost + 1}
			newPathInfo := pathInfo{pos: pair{x: curr.pos.x + dirX, y: curr.pos.y + dirY}, dir: curr.dir, cost: curr.cost + 1}
			queue = placeInTheRightPlace(queue, newPathInfo)
			
		}

		//rotate -90 -> check if the position is already in the visited
		dirT := curr.dir - 1
		if dirT < 0 {
			dirT = 3
		} 
		
		//add current position to visited, update with the new one 
		//tmpV2 := curr.visited
		//tmpV2 = append(tmpV2, vis{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: curr.dir})
		
		//insert into the right position based on the cost of the path
		//newPathInfo2 := pathInfo{visited: tmpV2, pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: dirT, cost: curr.cost + 1000}
		newPathInfo2 := pathInfo{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: dirT, cost: curr.cost + 1000}
		queue = placeInTheRightPlace(queue, newPathInfo2)


		//rotate 90 -> check if the position if already in visited
		dirTT := (curr.dir + 1) % 4

		//add current position to visited, update with the new one 
		//tmpV3 := curr.visited
		//tmpV3 = append(tmpV3, vis{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: curr.dir})
		
		//insert into the right position based on the cost of the path
		//newPathInfo3 := pathInfo{visited: tmpV3, pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: dirTT, cost: curr.cost + 1000}
		newPathInfo3 := pathInfo{pos: pair{x: curr.pos.x, y: curr.pos.y}, dir: dirTT, cost: curr.cost + 1000}
		queue = placeInTheRightPlace(queue, newPathInfo3)

		
	}
	
	return -1
}

func inVisited(el pathInfo, table []vis) bool {
	for _, i := range table {
		if el.pos.x == i.pos.x && el.pos.y == i.pos.y && el.dir == i.dir {
			return true
		}
	}

	return false
}

func placeInTheRightPlace(queue []pathInfo, pI pathInfo) []pathInfo {
	for i := 0; i < len(queue); i++ {
		
		if queue[i].cost > pI.cost {
			
			return append(queue[:i], append([]pathInfo{pI}, queue[i:]...)...)
		}
	}

	return append(queue, pI)
}


func checkIfExists(el pathInfo, dirX int, dirY int, table []vis) bool {
	for _, i := range table {
		//same coords and direction of walking
		if i.pos.x == el.pos.x + dirX && i.pos.y == el.pos.y + dirY && i.dir == el.dir {
			return false
		}
	}

	return true
}


func getDirection(dir int) (int, int) {
	dirX, dirY := 0, 0
	switch dir {
	case 0:
		dirX = 1
		break
	case 1:
		dirY = 1
		break
	case 2:
		dirX = -1
		break
	case 3:
		dirY = -1
		break
	default:
		fmt.Println("Wrong direction")
		break
	}

	return dirX, dirY
}