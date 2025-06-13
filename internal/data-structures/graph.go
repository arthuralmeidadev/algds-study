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

type VertexState struct {
	value string
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

func (s *VertexState) Unlinked() *VertexState {
	s.value = "unliked"
	return s
}

func (s *VertexState) Linked() *VertexState {
	s.value = "linked"
	return s
}

func (s *VertexState) Any() *VertexState {
	s.value = "any"
	return s
}

func (s VertexState) Value() string {
	if s.value == "" {
		return "any"
	}

	return s.value
}

func (v *Vertex) GetNearest(
	vertices []*Vertex,
	vertexState *VertexState,
) *Vertex {
	if len(vertices) == 0 {
		return nil
	}

	slices.DeleteFunc(vertices, func(delV *Vertex) bool {
		return delV == v
	})

	var nearest struct {
		value    *Vertex
		distance float32
	}

	for i := 0; i < len(vertices); i++ {
		vertexStateName := vertexState.Value()
		isLinked := slices.Contains(v.GetLinked(), vertices[i])
		checkVertex :=
			(vertexStateName == "linked" && isLinked) ||
				(vertexStateName == "unliked" && !isLinked) ||
				vertexStateName == "any"

		if checkVertex {
			d := vertices[i].CalcDistance(v.X, v.Y)

			if d <= nearest.distance {
				nearest.value = vertices[i]
				nearest.distance = d
			}
		}
	}

	return nearest.value
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

func (g *Graph) SelectLabeled(label string) bool {
	var v *Vertex
	for i := 0; i < len(g.vertices); i++ {
		if g.vertices[i].Label == label {
			v = g.vertices[i]
			break
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
	g.Link(false, .0)
}

func (g *Graph) Link(directed bool, weight float32) {
	vertexState := VertexState{}
	next := g.active.GetNearest(g.selected, vertexState.Unlinked())
	if next == nil {
		return
	}

	newEdge := &edge{
		head:     next,
		tail:     g.active,
		directed: directed,
		weight:   weight,
	}

	g.active.edges = append(g.active.edges, newEdge)
	next.edges = append(next.edges, newEdge)
	g.edges = append(g.edges, newEdge)
	g.active = next
	g.Link(directed, weight)
}

func (g *Graph) Unlink() {
	vertexState := VertexState{}
	next := g.active.GetNearest(g.selected, vertexState.Linked())

	if next == nil {
		return
	}

	for i := 0; i < len(g.edges); i++ {
		head, tail := g.edges[i].head, g.edges[i].tail
		fromActiveToNext := head == g.active && tail == next
		fromNextToActive := tail == g.active && head == next

		if fromActiveToNext || fromNextToActive {
			g.edges[i] = nil
			slices.Delete(g.edges, i, i+1)
		}
	}

	g.active = next
	g.Unlink()
}

func (g *Graph) Diameter() {

}

func (g *Graph) Path() {

}
