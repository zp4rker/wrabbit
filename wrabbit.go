package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wrabbit <command> [args]")
		os.Exit(1)
	}

	var cmd *exec.Cmd
	if len(os.Args) > 2 {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	} else {
		cmd = exec.Command(os.Args[1])
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process!")
		os.Exit(1)
	}

	fmt.Printf("Process started with PID %v\n", cmd.Process.Pid)

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Encountered an error while waiting for process!")
		os.Exit(1)
	}

	fmt.Printf("Process exited with status %v\n", cmd.ProcessState.ExitCode())
}
