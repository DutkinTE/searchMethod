package bfs

import (
	"testing"
)

func TestBFS(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	distances := graph.BFS(1)

	expected := map[int]int{
		1: 0,
		2: 1,
		3: 1,
		4: 2,
		5: 3,
	}

	for vertex, exp := range expected {
		if distances[vertex] != exp {
			t.Errorf("Для вершины %d ожидается %d, получено %d", vertex, exp, distances[vertex])
		}
	}

	if _, exists := distances[6]; exists {
		t.Errorf("Вершина 6 не должна быть достижима из вершины 1")
	}
}
