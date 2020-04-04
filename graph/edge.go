package graph

//Edge represents the relation between 2 nodes
type Edge struct {
	weight float64
}

//Weight of the edge
func (e *Edge) Weight() float64 {
	return e.weight
}
