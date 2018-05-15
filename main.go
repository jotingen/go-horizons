package main

import (
	"fmt"
)

import (
	"github.com/jotingen/go-horizons/horizons"
)

func main() {
	h, _ := horizons.New()

	list := h.MajorBodyList()
	for _, body := range list {
		fmt.Printf("%+v\n", body)
	}
}
