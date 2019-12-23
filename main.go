package main

import (
	"github.com/lishimeng/go-libs/log"
	"github.com/lishimeng/go-libs/stream/serial"
	"github.com/lishimeng/netserial/internal/cmd"
	"github.com/lishimeng/netserial/internal/relay"
	"time"
)

func main() {

	var err error
	log.SetLevelAll(log.DEBUG)// default log level

	var worker *relay.Worker
	err = cmd.Exec(func(p cmd.Params) {

		worker, err = relay.New(serial.Config{Baud: p.Baud, Name: p.SerialName}, p.Port)
		if err != nil {
			log.Info(err)
		} else {
			worker.Start()
		}
	})
	if err != nil {
		log.Info(err)
	}
	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		if worker != nil {
			worker.Close()
		}
		return
	}
}
