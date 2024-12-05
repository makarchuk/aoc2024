package dag

import "slices"

type Dag[Node comparable] struct {
	Nodes map[Node]struct{}
	Edges map[Node][]Node
}

func NewDag[Node comparable]() *Dag[Node] {
	return &Dag[Node]{
		Nodes: map[Node]struct{}{},
		Edges: map[Node][]Node{},
	}
}

func (d *Dag[Node]) AddNode(node Node) {
	d.Nodes[node] = struct{}{}
}

func (d *Dag[Node]) AddEdge(from, to Node) {
	d.Edges[from] = append(d.Edges[from], to)
}

func (d *Dag[Node]) Linearize() []Node {
	root := d.findRoot()
	return d.linearizeSubgraph(root)
}

func (d *Dag[Node]) linearizeSubgraph(root Node) []Node {
	linearized := []Node{}

	for _, node := range d.Edges[root] {
		linarizedSubgraph := d.linearizeSubgraph(node)
		for _, node := range linarizedSubgraph {
			if !slices.Contains(linearized, node) {
				linearized = append(linearized, node)
			}
		}
	}
	linearized = append(linearized, root)

	return linearized
}

func (d *Dag[Node]) findRoot() Node {

	pointedNodes := map[Node]struct{}{}

	for _, edges := range d.Edges {
		for _, node := range edges {
			pointedNodes[node] = struct{}{}
		}
	}

	for node := range d.Nodes {
		if _, ok := pointedNodes[node]; !ok {
			return node
		}
	}

	panic("no root node found")
}
