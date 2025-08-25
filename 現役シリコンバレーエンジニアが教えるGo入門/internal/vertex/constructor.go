package vertex

type Vertex struct {
	X, Y int
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}