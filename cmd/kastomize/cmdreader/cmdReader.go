package cmdreader

import (
	"bufio"
	"github.com/pieterclaerhout/go-log"
	"os/exec"
)

func ReadOutputToSlice(cmd *exec.Cmd) []string {
	lines := make([]string, 0)
	// Get a pipe to read from standard out
	r, _ := cmd.StdoutPipe()
	// Use the same pipe for standard error
	cmd.Stderr = cmd.Stdout
	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})
	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)
	// Use the scanner to scan the output line by line
	// It's running in a goroutine so that it doesn't block
	go func() {
		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			log.Debug(line)
			lines = append(lines, line)
		}
		// We're all done, unblock the channel
		done <- struct{}{}
	}()

	// Start the command and check for errors
	err := cmd.Start()
	log.CheckError(err)

	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
	err = cmd.Wait()
	log.CheckError(err)

	return lines
}
