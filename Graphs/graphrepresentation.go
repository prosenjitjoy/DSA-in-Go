package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex adds a Vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		newVertex := &Vertex{key: k}
		g.vertices = append(g.vertices, newVertex)
	}
}

// Add Edges add an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%d ", v.key)
		}
		fmt.Println()
	}
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(0)
	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.Print()
}
