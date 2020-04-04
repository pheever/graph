package graph

import (
	"errors"
	"math"
)

//Node represents a graph node
type Node struct {
	data  interface{}
	graph *Graph
	edges map[*Node]*Edge
}

//NewNode creates a new node
func NewNode(data interface{}) *Node {
	return &Node{
		data: &data,
	}
}

//Data of the node
func (n *Node) Data() interface{} {
	return n.data
}

func (n *Node) addEdge(e *Node, weight float64) error {
	if n.graph == nil {
		return errors.New("Node needs to be part of a graph before adding edges")
	}
	if n.edges == nil {
		n.edges = make(map[*Node]*Edge, 0)
	}
	if n.IsNeighbor(e) {
		return nil
	}
	n.edges[e] = &Edge{weight: weight}
	return nil
}

//Neighbors of ths node
func (n *Node) Neighbors() []*Node {
	if n.edges == nil || len(n.edges) == 0 {
		return []*Node{}
	}
	i := 0
	neighbors := make([]*Node, len(n.edges))
	for e := range n.edges {
		neighbors[i] = e
		i++
	}
	return neighbors
}

//HasEdges if there is at least 1 edge
func (n *Node) HasEdges() bool {
	return n.edges != nil && len(n.edges) > 0
}

func (n *Node) removeEdge(e *Node) error {
	if n.edges == nil {
		return errors.New("Node has no edges")
	}
	if _, isNeighbor := n.edges[e]; isNeighbor {
		delete(n.edges, e)
		return nil
	}
	return errors.New("Edge not found")
}

//IsNeighbor verifies edge to node
func (n *Node) IsNeighbor(node *Node) bool {
	if n.edges == nil {
		return false
	}
	_, isNeighbor := n.edges[node]
	return isNeighbor
}

//Breadth first search
func (n *Node) bfs() (visited map[*Node]int) {
	visited = make(map[*Node]int)
	fifo := []*Node{n}
	d := 0
	for len(fifo) > 0 {
		visited[fifo[0]] = d
		for _, e := range fifo[0].Neighbors() {
			if _, isVisited := visited[e]; !isVisited {
				fifo = append(fifo, e)
			}
		}
		fifo = fifo[1:]
		d++
	}
	return
}

//Dijkstra algorithm
func (n *Node) Dijkstra(to *Node) (dist map[*Node]float64, prev map[*Node]*Node) {
	dist = make(map[*Node]float64)
	prev = make(map[*Node]*Node)
	q := make(map[*Node]bool)
	for v := range n.graph.nodes {
		q[v] = true
		prev[v] = nil
		dist[v] = math.MaxFloat32
	}
	dist[n] = 0
	for len(q) > 0 {
		var u *Node = nil
		for v := range q {
			if u == nil || dist[v] < dist[u] {
				u = v //find u
			}
		}
		if u == to {
			return
		}
		delete(q, u) //remove u from q
		for v, e := range u.edges {
			if _, found := q[v]; !found {
				continue
			}
			alt := dist[u] + e.weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return
}

//PathTo other nodes
func (n *Node) PathTo(node *Node) Path {
	_, p := n.Dijkstra(node)

	path := Path{node}
	for u := p[node]; u != nil; u = p[u] {
		path = append(Path{u}, path...)
	}
	return path
}

//EdgeWeight of a valid edge, 0 other wise
func (n *Node) EdgeWeight(node *Node) float64 {
	if !n.IsNeighbor(node) {
		return 0
	}
	return n.edges[node].weight
}
