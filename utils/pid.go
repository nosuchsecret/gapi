package utils

import (
	"os"
	"errors"
	"strconv"
	"io/ioutil"
)

func SetPid(pidfile string) error {
	pid := os.Getpid()
	_, err := os.Lstat(pidfile)
	if err == nil {
		return errors.New("Pid file exists")
	}
	f, err := os.Create(pidfile)
	if err != nil {
		return err
	}
	defer f.Close()
	err = ioutil.WriteFile(pidfile, []byte(strconv.Itoa(pid)), 0777)
	if err != nil {
		return err
	}

	//add remove signal callback
	cb := func(arg interface{}) {
		if file, ok := arg.(string); ok {
			os.Remove(file)
		}
	}

	SignalCbAdd(cb, pidfile)

	return nil
}

func RemovePid(pidfile string) {
	os.Remove(pidfile)
}
