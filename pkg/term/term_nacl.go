// +build nacl

package term

import (
	"errors"
	"io"
	"os"
	"os/signal"
)

var (
	ErrInvalidState = errors.New("Invalid terminal state")
)

type State struct {
}

type Winsize struct {
	Height uint16
	Width  uint16
	x      uint16
	y      uint16
}

func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}

func GetFdInfo(in interface{}) (uintptr, bool) {
	var inFd uintptr
	var isTerminalIn bool
	if file, ok := in.(*os.File); ok {
		inFd = file.Fd()
		isTerminalIn = IsTerminal(inFd)
	}
	return inFd, isTerminalIn
}

func GetWinsize(fd uintptr) (*Winsize, error) {
	ws := &Winsize{}
	return ws, nil
}

func SetWinsize(fd uintptr, ws *Winsize) error {
	return nil
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd uintptr) bool {
	return true
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func RestoreTerminal(fd uintptr, state *State) error {
	if state == nil {
		return ErrInvalidState
	}
	return nil
}

func SaveState(fd uintptr) (*State, error) {
	var oldState State
	return &oldState, nil
}

func DisableEcho(fd uintptr, state *State) error {
	handleInterrupt(fd, state)
	return nil
}

func SetRawTerminal(fd uintptr) (*State, error) {
	oldState := State{}
	return &oldState, nil
}

func handleInterrupt(fd uintptr, state *State) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		_ = <-sigchan
		RestoreTerminal(fd, state)
		os.Exit(0)
	}()
}
