package graph

import (
	"errors"
)

//Graph represents a graph
type Graph struct {
	isDirected bool
	nodes      map[*Node]bool
}

//SetDirection to the graph
func (g *Graph) SetDirection(directed bool) error {
	if g.nodes != nil && len(g.nodes) != 0 {
		return errors.New("Direction cannot be set, nodes already exist")
	}
	g.isDirected = directed
	return nil
}

//Nodes of the graph
func (g *Graph) Nodes() []*Node {
	nodes := make([]*Node, len(g.nodes))
	i := 0
	for n := range g.nodes {
		nodes[i] = n
		i++
	}
	return nodes
}

//AddNode to the graph
func (g *Graph) AddNode(n *Node) {
	if g.nodes == nil {
		g.nodes = make(map[*Node]bool)
	}
	if _, isInGraph := g.nodes[n]; isInGraph {
		return
	}
	n.graph = g
	g.nodes[n] = false
}

//AddEdgeDefaultWeight from node a to node b, and vise versa if the graph is undirected with default weight 1
func (g *Graph) AddEdgeDefaultWeight(a, b *Node) error {
	return g.addEdge(a, b, 1)
}

//AddEdge from node a to node b, and vise versa if the graph is undirected with weight
func (g *Graph) AddEdge(a, b *Node, w float64) error {
	return g.addEdge(a, b, w)
}

func (g *Graph) addEdge(a, b *Node, w float64) error {
	if e := a.addEdge(b, w); e != nil {
		return e
	}
	if !g.isDirected {
		if e := b.addEdge(a, w); e != nil {
			return e
		}
	}
	return nil
}

//RemoveEdge from node a to node b, and vise versa if the graph is undirected
func (g *Graph) RemoveEdge(a, b *Node) error {
	if e := a.removeEdge(b); e != nil {
		return e
	}
	if !g.isDirected {
		if e := b.removeEdge(a); e != nil {
			return e
		}
	}
	return nil
}

//IsConnected if for every pair of nodes, there is a path between them
func (g *Graph) IsConnected() bool {
	if !g.isDirected {
		for n := range g.nodes {
			if !n.HasEdges() {
				return false
			}
		}
		return true
	}
	for n := range g.nodes {
		if len(n.bfs()) == len(g.nodes) {
			return true
		}
	}
	return false
}

//Distance of the path from node a to node b
func (g *Graph) Distance(a, b *Node) int {
	if d, reached := a.bfs()[b]; reached {
		return d
	}
	return -1
}
