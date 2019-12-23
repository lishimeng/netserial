package relay

import (
	"github.com/lishimeng/go-libs/log"
	"io"
)

func (w *Worker) run() {
	go w.rx()
	go w.tx()
}

func trans(to io.Writer, from io.Reader, bufSize int) (err error) {
	buf := make([]byte, bufSize)
	for {
		_, err = io.CopyBuffer(to, from, buf)
		if err != nil {
			return
		}
	}
}

func (w *Worker) rx() {
	defer func() {
		if err := recover(); err != nil {
			log.Info(err)
		}
	}()
	err := trans(w.socks, w.ser, w.bufSize)
	if err != nil {
		log.Info(err)
	}
}

func (w *Worker) tx() {
	defer func() {
		if err := recover(); err != nil {
			log.Info(err)
		}
	}()
	err := trans(w.ser, w.socks, w.bufSize)
	if err != nil {
		log.Info(err)
	}
}
