package main

import (
	"fmt"

	"github.com/pheever/graph/graph"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("hello graph")
	loadTest()
}

func loadTest() {
	g, e := graph.LoadGraph("test-graphs/graph1.csv")
	if e != nil {
		fmt.Println("failed to load graph ", e.Error())
		return
	}
	fmt.Println("graph loaded")
	fmt.Println("total nodes", len(g.Nodes()))
	a, _ := g.Nodes()[0].Dijkstra(nil)
	fmt.Println(a)
}

func gtest() {
	g := &graph.Graph{}
	nodes := []*graph.Node{
		&graph.Node{}, //0 f
		&graph.Node{}, //1 g
		&graph.Node{}, //2 h
		&graph.Node{}, //3 i
		&graph.Node{}, //4 j
		&graph.Node{}, //5 k
		&graph.Node{}, //6 l
		&graph.Node{}, //7 m
	}
	for _, n := range nodes {
		g.AddNode(n)
	}
	g.AddEdgeDefaultWeight(nodes[0], nodes[1])
	g.AddEdgeDefaultWeight(nodes[0], nodes[2])
	g.AddEdgeDefaultWeight(nodes[1], nodes[4])
	g.AddEdgeDefaultWeight(nodes[1], nodes[3])
	g.AddEdgeDefaultWeight(nodes[2], nodes[4])
	g.AddEdgeDefaultWeight(nodes[2], nodes[6])
	g.AddEdgeDefaultWeight(nodes[2], nodes[7])
	g.AddEdgeDefaultWeight(nodes[4], nodes[5])
	g.AddEdgeDefaultWeight(nodes[3], nodes[5])
	g.AddEdgeDefaultWeight(nodes[5], nodes[6])
	g.AddEdgeDefaultWeight(nodes[6], nodes[7])
	d := g.Distance(nodes[0], nodes[6])
	fmt.Println("distance f -> l:", d)

	dd, p := nodes[0].Dijkstra(nil)
	fmt.Println("distances:", dd)
	fmt.Println("prev:", p)
	fmt.Println("distance f -> l:", dd[nodes[5]])

	s := graph.Path{
		nodes[0],
		nodes[1],
		nodes[4],
		nodes[5],
		nodes[6],
	}
	fmt.Printf("path is valid: %t\n", s.Valid())
	fmt.Printf("path is cyclic: %t\n", s.IsCyclic())
	fmt.Printf("path distance: %.2f\n", s.Distance())
	fmt.Printf("path %+v\n", s)

	s = g.Nodes()[0].PathTo(g.Nodes()[6])
	fmt.Printf("path %+v\n", s)
	fmt.Printf("path is valid: %t\n", s.Valid())
	fmt.Printf("path is cyclic: %t\n", s.IsCyclic())
	fmt.Printf("path distance: %.2f\n", s.Distance())
}
