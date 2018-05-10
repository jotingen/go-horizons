package main

import (
	"fmt"
)

import (
	"github.com/jotingen/go-horizons/horizons"
)

func main() {
	h, _ := horizons.New()

	fmt.Printf("%+v\n", h)
	h.Dial()
}
