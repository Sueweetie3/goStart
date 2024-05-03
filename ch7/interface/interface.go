package main

import (
	"fmt"
	"time"
)

type Custom struct {}

func (c *Custom) String() string {
	return "Custom"
}

// define an interface
type Artifact interface {
    Title() string
    Creators() []string
    Created() time.Time
}

func main(){
	// var _ = Custom{}.String() // Error: cannot call pointer method on Custom literal, because of Custom literal is not addressable
	var c Custom
	var _ = c.String()
	var _ fmt.Stringer = &c
	// var _ fmt.Stringer = c // Error: cannot use c (type Custom) as interface fmt.Stringer in assignment, because c doesnt implement String function
}