package main

import (
	"fmt"
	"github.com/jkobyp/shapes"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type config struct {
	height      int
	width       int
	numElements int
	locations   []*shapes.Point
	direction   []*shapes.Point
	iterations  int
}

func main() {
	conf := config{}
	outfile := readArgs(&conf, os.Args)
	w := shapes.InitWindow(conf.height, conf.width)

	// Initialize a Window
	sdlWindow := shapes.NewSdlWindow(w)

	for i := 0; i < conf.numElements; i++ {
		w.AddElement(shapes.Circle{*conf.locations[i], 10})
	}

	shapes.Update(sdlWindow, w)
	// Main loop: slides points accross the screen
	for j := 0; j < conf.iterations; j++ {
		shapes.Update(sdlWindow, w)
		for i := 0; i < len(w.Elements); i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i, " of ", len(w.Elements), ": ", w.Elements)
			c := w.Elements[i].(shapes.Circle)
			c.Move(*conf.direction[i])
			fmt.Println(c)
			w.Elements[i] = c
			fmt.Println(w.Elements)
			fmt.Println()
			if !w.Area.Includes(w.Elements[i].Location()) {
				fmt.Printf("%v is not in area %v, so is breaking.\n", w.Elements[i].Location(), w.Area)
				break
			}
		}
	}
	writeOut(fmt.Sprintf("%v\n", w.Serialize()), outfile)
}

func writeOut(out string, outfile string) {
	var bytes []byte
	for _, b := range out {
		bytes = append(bytes, byte(b))
	}
	ioutil.WriteFile(outfile, bytes, 0644)
}

func readArgs(c *config, args []string) (out string) {
	if len(args) == 1 {
		panic("Arguments expected")
	}
	if len(args) > 1 {
		out = args[1]
	} else {
		out = "out.txt"
	}
	var err error
	c.height, err = strconv.Atoi(args[2])
	c.width, err = strconv.Atoi(args[3])
	c.iterations, err = strconv.Atoi(args[4])
	c.numElements, err = strconv.Atoi(args[5])
	if err != nil {
		panic("ahh!")
	}
	for _, arg := range args[6 : 6+c.numElements] {
		p, _ := shapes.DeserializePoint(arg)
		c.locations = append(c.locations, &p)
	}
	for _, arg := range args[6+c.numElements : len(args)] {
		p, _ := shapes.DeserializePoint(arg)
		c.direction = append(c.direction, &p)
	}
	return out
}
