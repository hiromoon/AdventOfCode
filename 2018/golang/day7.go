package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	input := readInput("day7")
	fmt.Println("part1: ", part1(input))
}

// Node is ...
type Node struct {
	ID       string
	children []*Node
}

// Graph is ..
type Graph struct {
	Start *Node
}

func (g *Graph) add(id string, node *Node) bool {
	if g.Start == nil {
		g.Start = &Node{
			id,
			[]*Node{node},
		}
		return true
	}

	if node.ID == g.Start.ID {
		tmp := g.Start
		g.Start = &Node{
			id,
			[]*Node{tmp},
		}
		return true
	}

	if n := g.find(id); n != nil {
		n.children = append(n.children, node)
		return true
	}

	return false
}

func (g *Graph) find(id string) *Node {
	if g.Start.ID == id {
		return g.Start
	}
	return g.Start.find(id)
}

func (n *Node) find(id string) *Node {
	if n.ID == id {
		return n
	}

	for _, node := range n.children {
		if tn := node.find(id); tn != nil {
			return tn
		}
	}

	return nil
}

func part1(input []string) string {
	graph := construct(input)
	fmt.Println(graph)
	return ""
}

func (g *Graph) printGraph() {
	g.Start.printNode()
}

func (n *Node) printNode() {
	fmt.Printf("->%s", n.ID)
	for _, node := range n.children {
		node.printNode()
	}
}

func hasKey(target map[string][]int, key string) bool {
	for k := range target {
		if k == key {
			return true
		}
	}
	return false
}

func construct(txt []string) *Graph {
	graph := &Graph{nil}
	cache := map[string][]*Node{}
	for _, i := range txt {
		id, node := parse(i)
		if !graph.add(id, node) {
			cache[id] = append(cache[id], node)
		}
	}

	graph.printGraph()

	cacheLength := len(cache)
	for cacheLength != 0 {
		removeList := map[string][]int{}
		for id, nodes := range cache {
			for index, node := range nodes {
				if graph.add(id, node) {
					if !hasKey(removeList, id) {
						removeList[id] = []int{}
					}
					removeList[id] = append(removeList[id], index)
				}
			}
		}

		if len(removeList) != 0 {
			for id, indexies := range removeList {
				sort.Slice(indexies, func(i, j int) bool { return indexies[i] > indexies[j] })
				for _, index := range indexies {
					cache[id] = append(cache[id][:index], cache[id][index+1:]...)
				}
				if len(cache[id]) == 0 {
					delete(cache, id)
				}
			}
		}

		cacheLength = len(cache)
		fmt.Println(cache)
	}

	return graph
}

func parse(txt string) (string, *Node) {
	r := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	result := r.FindStringSubmatch(txt)
	parentID := result[1]
	node := &Node{result[2], []*Node{}}
	return parentID, node
}
