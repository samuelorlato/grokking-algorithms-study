package main

import (
	"fmt"
	"strings"

	breadthfirstsearch "github.com/samuelorlato/grokking-algorithms-study/breadth_first_search"
)

func startsWithZ(s string) bool {
	return strings.Split(s, "")[0] == "z" || strings.Split(s, "")[0] == "Z"
}

func main() {
	graph := map[string][]string{}
	graph["arthur"] = []string{"alice", "ana", "clara"}
	graph["alice"] = []string{"jonas", "clara", "arthur"}
	graph["ana"] = []string{"samuel", "arthur"}
	graph["clara"] = []string{"matheus", "alice", "arthur"}
	graph["jonas"] = []string{}
	graph["samuel"] = []string{"zaqueu"}
	graph["matheus"] = []string{}
	graph["zaqueu"] = []string{"leandro"}
	graph["leandro"] = []string{}

	firstNameThatStartsWithZ := breadthfirstsearch.BreadthFirstSearch[string]("arthur", graph, startsWithZ)
	if firstNameThatStartsWithZ != nil {
		fmt.Println(*firstNameThatStartsWithZ)
	}
}
