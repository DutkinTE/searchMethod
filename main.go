package main

import (
	"fmt"

	"github.com/DutkinTE/bfs/bfs"
)

func main() {
	graph := bfs.NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	distances := graph.BFS(1)

	fmt.Println("Кратчайшие расстояния от вершины 1:")
	for vertex, distance := range distances {
		fmt.Printf("Вершина %d: %d\n", vertex, distance)
	}
}
