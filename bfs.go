package bfs

import (
	"container/list"
	"math"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g *Graph) BFS(start int) map[int]int {
	distances := make(map[int]int)
	for vertex := range g.adjacencyList {
		distances[vertex] = math.MaxInt32
	}
	distances[start] = 0

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		element := queue.Front()
		u := element.Value.(int)
		queue.Remove(element)

		for _, v := range g.adjacencyList[u] {
			if distances[v] == math.MaxInt32 {
				distances[v] = distances[u] + 1
				queue.PushBack(v)
			}
		}
	}
	return distances
}
