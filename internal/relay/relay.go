package relay

import (
	"fmt"
	"github.com/lishimeng/go-libs/log"
	"github.com/lishimeng/go-libs/stream/serial"
	"io"
	"net"
)

type Worker struct {

	socks io.ReadWriteCloser

	ser io.ReadWriteCloser

	server net.Listener

	listen uint16

	Ser serial.Config

	bufSize int
}

func New(serialConf serial.Config, listen uint16) (w *Worker, err error) {

	w = &Worker{
		Ser: serialConf,
		listen: listen,
		bufSize: 1024,
	}
	conn := serial.New(&serialConf)
	err = conn.Connect()
	if err != nil {
		return
	}
	w.ser = conn.Ser
	return
}

func (w *Worker) Start() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", w.listen))
	if err != nil {
		return
	}
	log.Info("start listen: %s", lis.Addr().String())
	w.server = lis
	for {
		conn, err := w.server.Accept()
		if err != nil {
			return
		}
		w.socks = conn
		w.run()
	}
}

func (w *Worker) Close() {
	if w.server != nil {
		_ = w.server.Close()// stop accept new connection
	}

	if w.socks != nil {
		_ = w.socks.Close()
	}

	if w.ser != nil {
		_ = w.ser.Close()
	}
}
