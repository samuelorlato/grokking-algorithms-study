package main

import (
	"fmt"
	"math"

	dijkstraalgorithm "github.com/samuelorlato/grokking-algorithms-study/dijkstra_algorithm"
)

func main() {
	graph := map[string]map[string]float64{}
	graph["start"] = map[string]float64{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = map[string]float64{}
	graph["a"]["end"] = 1
	graph["b"] = map[string]float64{}
	graph["b"]["a"] = 3
	graph["b"]["end"] = 5
	graph["end"] = map[string]float64{}

	costsMap := map[string]float64{}
	costsMap["a"] = 6
	costsMap["b"] = 2
	costsMap["end"] = math.Inf(1)

	parentsMap := map[string]string{}
	parentsMap["a"] = "start"
	parentsMap["b"] = "start"
	parentsMap["end"] = ""

	newParents, newCosts := dijkstraalgorithm.FindLowestCostPath(parentsMap, costsMap, graph)
	fmt.Println(newParents)
	fmt.Println(newCosts)
}
