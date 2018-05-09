package horizons

import (
	"github.com/reiver/go-telnet"
)

type Horizons struct {
	Caller telnet.Caller
}

func (h *Horizons) Call() {
	h.Caller = telnet.StandardCaller

	telnet.DialToAndCall("horizons.jpl.nasa.gov:6775", h.Caller)
}
