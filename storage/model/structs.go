package model

type Circle struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Radius int    `json:"radius"`
	Color  string `json:"color"`
	Exists bool   `json:"exists"`
}

// CREATE YOU STRUCTS HERE!
// YOU CAN USE "Circle" STRUCT AS AN EXAMPLE

// REGISTER YOU STRUCTS IN LIST
// YOU CAN USE DEFAULT STRUCT AS AN EXAMPLE
var Structs = map[string]func() interface{}{
	"circle": func() interface{} { return &Circle{} },
	// "structName": func() interface{} { return &Struct{} },
}
