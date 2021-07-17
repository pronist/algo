package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func tr(src io.Reader, dst io.Writer, errDst io.Writer) error {
	cmd := exec.Command("tr", "a-z", "A-Z")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		_, err := io.Copy(stdin, src)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
		} else if err != nil {
			log.Println("failed to write to STDIN", err)
		}
		stdin.Close()
		wg.Done()
	}()

	go func() {
		io.Copy(dst, stdout)
		stdout.Close()
		wg.Done()
	}()

	go func() {
		io.Copy(errDst, stderr)
		stderr.Close()
		wg.Done()
	}()
	wg.Wait()

	return cmd.Wait()
}

func main() {
	tr(os.Stdin, os.Stdout, os.Stderr)
}
