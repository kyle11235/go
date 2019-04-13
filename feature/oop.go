package main

import (
	"fmt"
)

// package encapsulation by lower case
type creature struct {
	Name string
	Real bool
}

// encapsulation by lower case
func (c creature) dump() {
	fmt.Printf("Name: %s, Real: %t\n", c.Name, c.Real)
}

// inheritance
type FlyingCreature struct {
	// nameless/anonymous - child has direct access to attributes and methods
	// deadly diamond of death
	creature
	Wings int
}

// inheritance
type WalkingCreature struct {
	creature
	Legs int
}

// interface
type Actor interface {
	Move()
}

func (c *FlyingCreature) Move() {
	c.dump()
	fmt.Println(c.Name + " is flying")
}

func (c *WalkingCreature) Move() {
	c.dump()
	fmt.Println(c.Name + " is walking")
}

// oop - package encapsulation/inheritance/polymorphism
// decoupled type and interface, no need to declare any relationship ahead of time
func main() {

	// polymorphism
	var a1 Actor = &FlyingCreature{
		creature{"Dragon", false},
		2,
	}
	var a2 Actor = &WalkingCreature{
		creature{"Dog", true},
		4,
	}

	a1.Move()
	a2.Move()

}
