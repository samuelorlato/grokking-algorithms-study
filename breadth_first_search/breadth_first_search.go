package breadthfirstsearch

import "golang.org/x/exp/constraints"

// O(V+A)
func BreadthFirstSearch[T constraints.Ordered](initialVertice T, graph map[T][]T, f func(k T) bool) *T {
	queue := []T{}
	queue = append(queue, graph[initialVertice]...)

	verified := map[T]bool{}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if !verified[element] {
			if f(element) {
				return &element
			} else {
				queue = append(queue, graph[element]...)
				verified[element] = true
			}
		}
	}

	return nil
}