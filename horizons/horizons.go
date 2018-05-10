package horizons

import (
	"fmt"
)

import (
	"github.com/reiver/go-telnet"
)

type Horizons struct {
	Conn *telnet.Conn
}

func New() (Horizons, error) {
	conn, err := telnet.DialToAndCall("horizons.jpl.nasa.gov:6775")
	if err != nil {
		return Horizons{}, err
	}
	return Horizons{
		Conn: conn,
	}, nil
}

func (h *Horizons) Dial() ([]byte, error) {
	var buf []byte

	h.Conn.Write([]byte("?"))
	h.Conn.Write([]byte("\n"))

	length, err := h.Conn.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	fmt.Printf("%d: %s\n", length, buf)
	length, err = h.Conn.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	fmt.Printf("%d: %s\n", length, buf)
	length, err = h.Conn.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	fmt.Printf("%d: %s\n", length, buf)
	return buf, nil
}
