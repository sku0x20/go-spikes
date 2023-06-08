package main

// https://www.golinuxcloud.com/golang-monitor-background-process/#Method-1_Using_Wait_method
import (
	"context"
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	wait()
	commandContext()
	goroutineWithChannels()
	goroutineWithWaitGroup()
}

// using exec.Command()
func wait() {
	cmd := exec.Command("echo", "from echo")
	outstream, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process started with pid: %d\n", cmd.Process.Pid)

	buffer := make([]byte, 1024)
	for {
		_, err = outstream.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer))
	}

	// wait for the process to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("process completed.")
}

// exec.CommandContext()
func commandContext() {
	// Create a background context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the process in the background
	cmd := exec.CommandContext(ctx, "sleep", "10")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error, %v\n", err)
		return
	}
	fmt.Printf("process started with pid: %d\n", cmd.Process.Pid)

	// Wait for the process to complete
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Process Completed")
}

// goroutine
func goroutineWithChannels() {
	// Start the process in a goroutine
	done := make(chan error)
	go func() {
		cmd := exec.Command("sleep", "60")
		err := cmd.Start()
		if err != nil {
			done <- err
			return
		}
		fmt.Printf("Process started with pid: %d\n", cmd.Process.Pid)
		done <- cmd.Wait()
	}()

	// Wait for the goroutine to finish
	err := <-done
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Process completed")
}

// sync.WaitGroup()
func goroutineWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("sleep", "10")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Process started with pid: %d\n", cmd.Process.Pid)
		err = cmd.Wait()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}()
	wg.Wait()
	fmt.Printf("process completed")
}
