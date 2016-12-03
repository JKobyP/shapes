package shapes

import (
	"bytes"
	"encoding/json"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	"strings"
)

const (
	WindowHeight = 400
	WindowWidth  = 600
)

type Shape Element

type Element interface {
	Location() Point
}

func NewSdlWindow(w *Window) *sdl.Window {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w.Area.Width, w.Area.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	window.UpdateSurface()
	return window
}

func Update(window *sdl.Window, w *Window) bool {
	surface, _ := window.GetSurface()

	//background
	var x = int32(w.Area.Location().X)
	var y = int32(w.Area.Location().Y)
	rect := sdl.Rect{x, y, int32(w.Area.Width), int32(w.Area.Height)}
	surface.FillRect(&rect, 0x00000000)

	for _, s := range w.Elements {
		switch e := s.(type) {
		case Circle:
			var x = int32(e.Location().X)
			var y = int32(e.Location().Y)
			rect := sdl.Rect{X: x, Y: y, W: 10, H: 10}
			surface.FillRect(&rect, 0xffff0000)
		case Rectangle:
			var x = int32(e.Location().X)
			var y = int32(e.Location().Y)
			rect := sdl.Rect{x, y, x + int32(e.Width), y + int32(e.Height)}
			surface.FillRect(&rect, 0xffff0000)
		}
	}
	window.UpdateSurface()
	return true
}

func Exit(window *sdl.Window) {
	window.Destroy()
	sdl.Quit()
}

// A Window represents a visible portion of the screen.
type Window struct {
	Area     *Rectangle
	Elements []Shape
}

func InitWindow(x, y int) *Window {
	rectangle := &Rectangle{Point{0, 0}, x, y}
	return &Window{rectangle, make([]Shape, 0)}
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

func (w *Window) Serialize() string {
	var b = new(bytes.Buffer)
	e := json.NewEncoder(b)
	for _, elem := range w.Elements {
		e.Encode(elem)
	}
	return b.String()
}

// A Point is a 2d point in euclidean space.
type Point struct {
	X int
	Y int
}

func (p Point) Location() Point {
	return p
}

func (p *Point) Move(offset Point) {
	p.X += offset.X
	p.Y += offset.Y
}

func DeserializePoint(p string) (Point, bool) {
	if len(p) < 3 {
		return Point{}, false
	}
	s := strings.Split(p, ",")
	x, err := strconv.Atoi(s[0])
	y, err := strconv.Atoi(s[1])
	if err != nil {
		return Point{}, false
	}
	return Point{x, y}, true
}

// A Rectangle represents a rectangle with an origin point at its top left.
type Rectangle struct {
	Origin Point
	Height int
	Width  int
}

func (r *Rectangle) Includes(p Point) bool {
	return p.X <= r.Height && p.X >= 0 && p.Y <= r.Width && p.Y >= 0
}

func (r Rectangle) Location() Point {
	return r.Origin
}

func (r Rectangle) Area() int {
	return r.Height * r.Width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

// A Circle is a circle with an origin point at its center.
type Circle struct {
	Origin Point
	Radius int
}

func (c *Circle) Diameter() int {
	return c.Radius * 2
}

func (c *Circle) SetRadius(v int) {
	c.Radius = v
}

func (c *Circle) Move(p Point) {
	c.Origin.X += p.X
	c.Origin.Y += p.Y
}

func (c Circle) Location() Point {
	return c.Origin
}

func (c Circle) Area() int {
	const pi float64 = 3.141592653589793238
	return int(pi * float64(c.Radius*c.Radius))
}
