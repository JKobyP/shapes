package shapes

const (
	WindowHeight = 400
	WindowWidth  = 600
)

type Shape interface {
	Element
	Area() int
}

type Element interface {
	Location() Point
}

// A Window represents a visible portion of the screen.
type Window struct {
	Area     *Rectangle
	Elements []Element
}

// AddElement adds a new element to the Window
func (w *Window) AddElement(s Element) {
	w.Elements = append(w.Elements, s)
}

func (w *Window) Height() int {
	return w.Area.Height
}

func (w *Window) Width() int {
	return w.Area.Width
}

// A Point is a 2d point in euclidean space.
type Point struct {
	X int
	Y int
}

// A Rectangle represents a rectangle with an origin point at its top left.
type Rectangle struct {
	origin Point
	Height int
	Width  int
}

// Includes returns true if p is located within r
func (r *Rectangle) Includes(p Point) bool {
	return p.X < r.Height-r.Location().X && p.X > 0 && p.Y < r.Width-r.Location().Y && p.Y > 0
}

func (r Rectangle) Location() Point {
	return r.origin
}

func (r Rectangle) Area() int {
	return r.Height * r.Width
}

// A Circle is a circle with an origin point at its center.
type Circle struct {
	origin Point
	Radius int
}

func (c *Circle) Diameter() int {
	return c.Radius * 2
}

func (c *Circle) Move(offset Point) {
	c.origin.X += offset.X
	c.origin.Y += offset.Y
}

func (c Circle) Location() Point {
	return c.origin
}

func (c Circle) Area() int {
	const pi float64 = 3.141592653589793238
	return int(pi * float64(c.Radius*c.Radius))
}

func InitWindow(width, height int) *Window {
	var origin Point = Point{0, 0}
	rectangle := &Rectangle{origin, width, height}
	shapeSlice := make([]Element, 0)
	return &Window{rectangle, shapeSlice}
}
