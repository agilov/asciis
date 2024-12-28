package game

import (
	"os"
	"os/exec"
	"time"
)

const arrowUp byte = 1
const arrowDown byte = 2
const arrowRight byte = 3
const arrowLeft byte = 4
const enter byte = 13

type input struct {
	keyChan   chan byte
	ioReadBuf []byte
}

func newInput() *input {
	in := &input{
		keyChan:   make(chan byte, 1),
		ioReadBuf: make([]byte, 1),
	}
	go in.listen()
	return in
}

// Function to enable raw mode on the terminal
func terminalRaw() {
	cmd := exec.Command("stty", "--save")
	cmd.Stdin = os.Stdin // error otherwise: stty: 'standard input': Inappropriate ioctl for device
	if err := cmd.Run(); err != nil {
		panic("failed to save terminal raw input: " + err.Error())
	}
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "-echo", "min", "1").Run()
}

// Restore terminal
func terminalRestore() {
	cmd := exec.Command("stty", "sane")
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func (in *input) readStd() {
	in.ioReadBuf = in.ioReadBuf[:1]
	os.Stdin.Read(in.ioReadBuf)
}

func (in *input) listen() {
	for {
		in.readStd()
		if in.ioReadBuf[0] == '\x1b' {
			in.readArrow()
		} else {
			in.keyChan <- in.ioReadBuf[0]
		}
	}
}

func (in *input) readArrow() {
	in.readStd()
	if in.ioReadBuf[0] != '[' {
		return
	}
	in.readStd()
	switch in.ioReadBuf[0] {
	case 'A':
		in.keyChan <- arrowUp
	case 'B':
		in.keyChan <- arrowDown
	case 'C':
		in.keyChan <- arrowRight
	case 'D':
		in.keyChan <- arrowLeft
	default:
	}
}

func (in *input) read() event {
	select {
	case key := <-in.keyChan:
		return &keyPressed{key: key}
	case <-time.After(time.Millisecond * 4):
		return &eventNone{}
	}
}
