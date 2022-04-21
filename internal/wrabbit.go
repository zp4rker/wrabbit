package wrabbit

import (
	"fmt"
	"net"
	"os/exec"
)

type Wrabbit struct {
	Statfile *Statfile
	Listener net.Listener
	Cmd      *exec.Cmd
}

// initialise the unix socket and output
func (wr *Wrabbit) Init() error {
	addr := fmt.Sprintf("%v/api.sock", wr.Statfile.Dir)

	// create listener
	l, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}

	wr.Listener = l
	return nil
}

func (wr *Wrabbit) Cleanup() {
	// close listener
	if wr.Listener != nil {
		wr.Listener.Close()
	}
}
