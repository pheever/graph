package graph

//Path is a set of nodes
type Path []*Node

//Valid if each consecutive pair in the sequence is connected by an edge
func (s *Path) Valid() bool {
	for i, n := range (*s)[:len((*s))-1] {
		if !n.IsNeighbor((*s)[i+1]) {
			return false
		}
	}
	return true
}

//IsCyclic is a path with at least three edges, in which the first and last nodes are the same
func (s *Path) IsCyclic() bool {
	edges := len((*s)) > 2
	isValid := (*s).Valid()
	l2f := (*s)[len((*s))-1].IsNeighbor((*s)[0])
	return isValid && edges && l2f
}

//Distance of the sequence of the nodes
func (s *Path) Distance() float64 {
	d := 0.0
	if s.Valid() {
		for i, n := range (*s)[:len((*s))-1] {
			d += n.EdgeWeight((*s)[i+1])
		}
	}
	return d
}
