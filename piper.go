package piper

import (
	"bufio"
	"log"
	"os/exec"
)

// MustPipe returns stdout and stderr line scanners (bufio.Scanner) for the
// command created from the given argument. It calls log.Fatal() if the command
// could not be started, or if the stdout or stderr plumbings failed
func MustPipe(exe string, args ...string) (*bufio.Scanner, *bufio.Scanner) {
	cmd := exec.Command(exe, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(exe, err)
	}

	return bufio.NewScanner(stdout), bufio.NewScanner(stderr)
}
