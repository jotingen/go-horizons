package horizons

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

import (
	"github.com/ziutek/telnet"
)

type Horizons struct {
	Conn *telnet.Conn
}

func New() (Horizons, error) {
	conn, err := telnet.Dial("tcp", "horizons.jpl.nasa.gov:6775")
	if err != nil {
		return Horizons{}, err
	}
	_, err = conn.ReadBytes('>')
	if err != nil {
		return Horizons{}, err
	}
	return Horizons{
		Conn: conn,
	}, nil
}

type Body struct {
	ID          string
	Name        string
	Designation string
	Other       string
}

func (h *Horizons) MajorBodyList() []Body {
	var list []Body

	//Request major body list
	h.Conn.Write([]byte("MB\n"))
	data, _ := h.Conn.ReadUntil("Select ... [F]tp, [M]ail, [R]edisplay, ?, <cr>:")

	//Process response for bodies
	reEntry := regexp.MustCompile(`^\s*-?\d+\s`)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if reEntry.MatchString(line) {

			id := line[0:9]
			name := line[11:45]
			designation := line[46:58]
			other := line[59:len(line)]

			id = strings.TrimSpace(id)
			name = strings.TrimSpace(name)
			designation = strings.TrimSpace(designation)
			other = strings.TrimSpace(other)

			list = append(list, Body{
				ID:          id,
				Name:        name,
				Designation: designation,
				Other:       other,
			})
		}
	}

	//Move prompt to main
	h.Conn.Write([]byte("\n"))
	data, _ = h.Conn.ReadUntil("Horizons>")

	return list
}
