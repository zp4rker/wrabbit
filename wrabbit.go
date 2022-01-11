package main

import (
	"fmt"
	"os"
	"os/exec"

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

	// start command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process!")
		os.Exit(1)
	}
	fmt.Printf("Process started with PID %v\n", cmd.Process.Pid)

	// prepare statfile
	sf, err := wrabbit.PrepareStatfile()
	if err != nil {
		fmt.Println("Failed to prepare statfile! Killing process...")
		cmd.Process.Kill()
		os.Exit(1)
	}

	// fill statfile
	sf.Data.ProcId = cmd.Process.Pid
	sf.Data.Args = cmd.Args
	sf.Data.Running = true

	// wait for process to exit
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Encountered an error while waiting for process!")
		os.Exit(1)
	}
	fmt.Printf("Process exited with status %v\n", cmd.ProcessState.ExitCode())
}
