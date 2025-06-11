package datastructures

import (
	"math"
	"slices"
)

type VertexDegree struct{ total, in, out uint }

type Vertex struct {
	Label                           string
	AllowsLoops, AllowParallelEdges bool
	MaxDegree                       uint
	X, Y                            float32
	edges                           []*edge
}

func (v *Vertex) getNearestUnlinked(vertices []*Vertex) *Vertex {
	var unlinkedVertices []*Vertex

	for i := 0; i < len(vertices); i++ {
		if !slices.Contains(v.GetLinked(), vertices[i]) && v != vertices[i] {
			unlinkedVertices = append(unlinkedVertices, vertices[i])
		}
	}

	if len(unlinkedVertices) == 0 {
		return nil
	}

	var nearestUnlinked struct {
		value    *Vertex
		distance float32
	}

	for i := 0; i < len(unlinkedVertices); i++ {
		d := unlinkedVertices[i].CalcDistance(v.X, v.Y)

		if d <= nearestUnlinked.distance {
			nearestUnlinked.value = unlinkedVertices[i]
			nearestUnlinked.distance = d
		}
	}

	return nearestUnlinked.value
}

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

func (v *Vertex) GetLinked() []*Vertex {
	var siblings []*Vertex

	for i := 0; i < len(v.edges); i++ {
		if v.edges[i].tail == v {
			siblings = append(siblings, v.edges[i].head)
		}

		if v.edges[i].head == v {
			siblings = append(siblings, v.edges[i].tail)
		}
	}

	return siblings
}

func (v *Vertex) CalcDistance(dX, dY float32) float32 {
	if v.X == dX && v.Y == dY {
		return .0
	}

	if v.Y == dY {
		minimum, maximum := min(v.X, dX), max(v.X, dX)

		if minimum < .0 {
			maximum = maximum + (minimum * -1)
			minimum = .0
		}

		return maximum - minimum
	}

	if v.X == dX {
		minimum, maximum := min(v.Y, dY), max(v.Y, dY)

		if minimum < .0 {
			maximum = maximum + (minimum * -1)
			minimum = .0
		}

		return maximum - minimum
	}

	var coords struct {
		min struct {
			x float32
			y float32
		}

		max struct {
			x float32
			y float32
		}
	}

	coords.min.x, coords.max.x = min(v.X, dX), max(v.X, dX)

	if coords.min.x < .0 {
		coords.max.x = coords.max.x + (coords.min.x * -1)
		coords.min.x = .0
	}

	coords.min.y, coords.max.y = min(v.Y, dY), max(v.Y, dY)

	if coords.min.y < .0 {
		coords.max.y = coords.max.y + (coords.min.y * -1)
		coords.min.y = .0
	}

	squaredWidth := math.Pow(float64(coords.max.x), 2)
	squaredHeight := math.Pow(float64(coords.max.y), 2)

	return float32(math.Sqrt(squaredWidth + squaredHeight))
}

type edge struct {
	head, tail *Vertex
	directed   bool
	weight     float32
}

type Graph struct {
	vertices, selected []*Vertex
	edges              []*edge
	active             *Vertex
}

func (g *Graph) SelectLabeled(label string) bool {
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
	g.ClearSelection()
	g.selected = append(g.selected, g.active)
	g.active = v
	g.Link()
}

func (g *Graph) Link() {
	next := g.active.getNearestUnlinked(g.selected)
	if next == nil {
		return
	}

	newEdge := &edge{
		head:     next,
		tail:     g.active,
		directed: false,
		weight:   0,
	}

	g.active.edges = append(g.active.edges, newEdge)
	next.edges = append(next.edges, newEdge)
	g.edges = append(g.edges, newEdge)
	g.active = next
	g.Link()
}

func (g *Graph) Unlink() {

}

func (g *Graph) Diameter() {

}

func (g *Graph) Path() {

}
