package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	edges := make(map[string]map[string]struct{})

	// Parse edges
	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, "-")
		if len(nodes) != 2 {
			continue
		}
		addEdge(edges, nodes[0], nodes[1])
		addEdge(edges, nodes[1], nodes[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	result := findTriplets(edges)
	fmt.Println("Number of triplets:", result)
}

func addEdge(edges map[string]map[string]struct{}, a, b string) {
	if _, exists := edges[a]; !exists {
		edges[a] = make(map[string]struct{})
	}
	edges[a][b] = struct{}{}
}

func findTriplets(edges map[string]map[string]struct{}) int {
	triplets := make(map[string]struct{})

	for x, neighbors := range edges {
		for y := range neighbors {
			for z := range edges[y] {
				if x != z && hasEdge(edges, x, z) {
					triplet := []string{x, y, z}
					sort.Strings(triplet)
					tripletsKey := strings.Join(triplet, ",")
					triplets[tripletsKey] = struct{}{}
				}
			}
		}
	}


	count := 0
	for triplet := range triplets {
		if triplet[0] == 't' || triplet[3] == 't' || triplet[6] == 't' {
			count++
		}
	}
	return count
}

func hasEdge(edges map[string]map[string]struct{}, a, b string) bool {
	_, exists := edges[a][b]
	return exists
}
