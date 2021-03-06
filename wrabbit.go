package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	wrabbit "github.com/zp4rker/wrabbit/internal"
)

func main() {
	// validate args
	if len(os.Args) < 2 {
		fmt.Println("Usage: wrabbit <command> [args]")
		os.Exit(1)
	}

	// setup command
	var cmd *exec.Cmd
	if len(os.Args) > 2 {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	} else {
		cmd = exec.Command(os.Args[1])
	}

	// prepare statfile
	sf, err := wrabbit.PrepareStatfile()
	if err != nil {
		fmt.Println("Failed to prepare statfile! Aborting...")
		os.Exit(1)
	}

	// wrabbit struct
	wr := wrabbit.Wrabbit{Statfile: sf, Cmd: cmd}
	defer wr.Cleanup()

	// initialise wrabbit api and start listening
	err = wr.Init()
	if err != nil {
		fmt.Println("Failed to initialise wrabbit API! Aborting...")
		wr.Cleanup()
		os.Exit(1)
	}
	go wr.Listen()

	// start command
	err = cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process!")
		os.Exit(1)
	}
	fmt.Printf("Process started with PID %v\n", cmd.Process.Pid)

	// fill statfile
	sf.Data.ProcId = cmd.Process.Pid
	sf.Data.Args = cmd.Args
	sf.Data.Running = true
	err = sf.Update()
	if err != nil {
		fmt.Println("Failed to update statfile! Killing process...")
		cmd.Process.Kill()
		os.Exit(1)
	}

	// start polling
	pollStop := make(chan bool)
	go sf.StartPoll(&pollStop)

	// wait for process to exit
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Encountered an error while waiting for process!")
		wr.Cleanup()
		os.Exit(1)
	}
	fmt.Printf("Process exited with status %v\n", cmd.ProcessState.ExitCode())

	// stop polling
	pollStop <- true

	// update statfile
	sf.Data.Running = false
	sf.Data.EndDate = time.Now()
	err = sf.Update()
	if err != nil {
		fmt.Println("Failed to update statfile!")
		wr.Cleanup()
		os.Exit(1)
	}
}
