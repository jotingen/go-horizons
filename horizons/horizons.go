package horizons

//import (
//	"os"
//	"time"
//
//	"github.com/ziutek/telnet"
//)
//
//const timeout = 10 * time.Second
//
////func checkErr(err error) {
////	if err != nil {
////		log.Fatalln("Error:", err)
////	}
////}
//
//func expect(t *telnet.Conn, d ...string) error {
//	err := t.SetReadDeadline(time.Now().Add(timeout))
//	if err != nil {
//		return err
//	}
//	err = t.SkipUntil(d...)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func sendln(t *telnet.Conn, s string) error {
//	err := t.SetWriteDeadline(time.Now().Add(timeout))
//	if err != nil {
//		return err
//	}
//	buf := make([]byte, len(s)+1)
//	copy(buf, s)
//	buf[len(s)] = '\n'
//	_, err = t.Write(buf)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//type Horizons struct {
//	Conn *telnet.Conn
//}
//
//func New() (Horizons, error) {
//	conn, err := telnet.Dial("tcp", "horizons.jpl.nasa.gov:6775")
//	if err != nil {
//		return Horizons{}, err
//	}
//	conn.SetUnixWriteMode(true)
//	return Horizons{
//		Conn: conn,
//	}, nil
//}
//
//func (h *Horizons) Dial() error {
//	err := expect(h.Conn, ">")
//	if err != nil {
//		return err
//	}
//	err = sendln(h.Conn, "?")
//	if err != nil {
//		return err
//	}
//	data, _ := h.Conn.ReadBytes('>')
//
//	os.Stdout.WriteString("TEST")
//	os.Stdout.Write(data)
//	os.Stdout.WriteString("TEST")
//	os.Stdout.WriteString("\n")
//	return nil
//}

//import (
//	"fmt"
//)
//
//import (
//	"github.com/reiver/go-telnet"
//)
//
//type Horizons struct {
//	Conn *telnet.Conn
//}
//
//func New() (Horizons, error) {
//	conn, err := telnet.DialTo("horizons.jpl.nasa.gov:6775")
//	if err != nil {
//		return Horizons{}, err
//	}
//	return Horizons{
//		Conn: conn,
//	}, nil
//}
//
//func (h *Horizons) Dial() ([]byte, error) {
//	var buf []byte
//
//	h.Conn.Write([]byte("?"))
//	h.Conn.Write([]byte("\n"))
//
//	length, err := h.Conn.Read(buf)
//	if err != nil {
//		return []byte{}, err
//	}
//	fmt.Printf("%d: %s\n", length, buf)
//	length, err = h.Conn.Read(buf)
//	if err != nil {
//		return []byte{}, err
//	}
//	fmt.Printf("%d: %s\n", length, buf)
//	length, err = h.Conn.Read(buf)
//	if err != nil {
//		return []byte{}, err
//	}
//	fmt.Printf("%d: %s\n", length, buf)
//	return buf, nil
//}

import (
	"io"
	"net"
	"os"
)

type Horizons struct {
	Conn net.Conn
}

func New() (Horizons, error) {
	conn, err := net.Dial("tcp", "horizons.jpl.nasa.gov:6775")
	if err != nil {
		return Horizons{}, err
	}
	return Horizons{
		Conn: conn,
	}, nil
}

func (h *Horizons) Dial() {
	errc := make(chan error)
	go cp(h.Conn, os.Stdin, errc)
	go cp(os.Stdout, h.Conn, errc)
	<-errc
}

func cp(dst io.Writer, src io.Reader, errc chan<- error) {
	_, err := io.Copy(dst, src)
	errc <- err
}
