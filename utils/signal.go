package utils

import (
	"os"
	"syscall"
	"os/signal"
)

type SigCb func(arg interface{})
type SigItem struct {
	Cb SigCb
	Arg interface{}
}

var sigcb []SigItem

func SignalCbAdd(cb SigCb, arg interface{}) {
	scb := SigItem{Cb: cb, Arg: arg}
	sigcb = append(sigcb, scb)
}

func SignalInit() {
	go SignalHandler()
}

func SignalHandler() {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			for _, i := range sigcb {
				i.Cb(i.Arg)
			}
			os.Exit(1);
		}
	}
}

