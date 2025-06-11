package datastructures

type Vertex struct {
	Label                           string
	AllowsLoops, AllowParallelEdges bool
	MaxDegree                       uint
	X, Y                            float32
	edges                           []*edge
}

type VertexDegree struct{ total, in, out uint }

func (v *Vertex) Degree() *VertexDegree {
	d := new(VertexDegree)

	for i := 0; i < len(v.edges); i++ {
		if v.edges[i].head == v {
			d.in += 1
			d.total += 1
		}

		if v.edges[i].tail == v {
			d.out += 1
			d.total += 1
		}
	}

	return d
}

type edge struct {
	head, tail *Vertex
	directed   bool
	weight     float32
}

type Graph struct {
	vertices, selected []*Vertex
	edges              []edge
	active             *Vertex
}

func (g *Graph) SelectVertex(label string) bool {
	var v *Vertex
	for i := 0; i < len(g.vertices); i++ {
		if g.vertices[i].Label == label {
			v = g.vertices[i]
		}
	}

	if v == nil {
		return false
	}

	g.selected = append(g.selected, v)
	return true
}

func (g *Graph) GetSelected() []*Vertex {
	return g.selected
}

func (g *Graph) ClearSelection() {
	g.selected = make([]*Vertex, 0)
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)

	if len(g.selected) > 0 {
		head := v

		for i := 0; i < len(g.selected); i++ {
			tail := g.selected[i]
			isHeadAtMaxDeg := head.Degree().total == head.MaxDegree
			isTailAtMaxDeg := tail.Degree().total == tail.MaxDegree

			if isHeadAtMaxDeg || isTailAtMaxDeg {
				break
			}

			newEdge := edge{
				head: head,
				tail: tail,
			}
			v.edges = append(v.edges, &newEdge)
			g.edges = append(g.edges, newEdge)
		}
	}

	g.ClearSelection()
	g.selected = append(g.selected, v)
	g.active = v
}

func (g *Graph) Connect() {

}

func (g *Graph) Diameter() {

}

func (g *Graph) Path() {

}
