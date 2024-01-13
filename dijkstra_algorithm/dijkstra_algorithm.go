package dijkstraalgorithm

import (
	"math"

	"golang.org/x/exp/constraints"
)

func getMapKeys[T constraints.Ordered](m map[string]T) []string {
	i := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func findLowestCostNode(costsMap map[string]float64, verified map[string]bool) string {
	lowestCost := math.Inf(1)
	var lowestCostNode string

	for node := range costsMap {
		cost := costsMap[node]
		if cost < lowestCost && !verified[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func FindLowestCostPath(parentsMap map[string]string, costsMap map[string]float64, graph map[string]map[string]float64) (map[string]string, map[string]float64) {
	verified := map[string]bool{}

	node := findLowestCostNode(costsMap, verified)

	for node != "" {
		cost := costsMap[node]
		neighbors := graph[node]

		keys := getMapKeys[float64](neighbors)
		for _, n := range keys {
			newCost := cost + neighbors[n]
			if costsMap[n] > newCost {
				costsMap[n] = newCost
				parentsMap[n] = node
			}
		}

		verified[node] = true
		node = findLowestCostNode(costsMap, verified)
	}

	return parentsMap, costsMap
}
