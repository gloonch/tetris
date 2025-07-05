package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ListenInput(cmds chan<- string) {
	var sttyCbreak, sttyEcho, sttyNoEcho, sttyNoCbreak *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		sttyCbreak = exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1")
		sttyNoEcho = exec.Command("stty", "-f", "/dev/tty", "-echo")
		sttyEcho = exec.Command("stty", "-f", "/dev/tty", "echo")
		sttyNoCbreak = exec.Command("stty", "-f", "/dev/tty", "-cbreak")
	case "linux":
		sttyCbreak = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
		sttyNoEcho = exec.Command("stty", "-F", "/dev/tty", "-echo")
		sttyEcho = exec.Command("stty", "-F", "/dev/tty", "echo")
		sttyNoCbreak = exec.Command("stty", "-F", "/dev/tty", "-cbreak")
	default:
		// no need for other os
	}

	if sttyCbreak != nil {
		sttyCbreak.Run()
	}
	if sttyNoEcho != nil {
		sttyNoEcho.Run()
	}
	defer func() {
		if sttyEcho != nil {
			sttyEcho.Run()
		}
		if sttyNoCbreak != nil {
			sttyNoCbreak.Run()
		}
	}()

	f, err := os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer f.Close()

	buf := make([]byte, 3)
	for {
		n, _ := f.Read(buf)
		if n == 1 {
			switch buf[0] {
			case 'q':
				cmds <- "quit"
			case 'r', 'R':
				cmds <- "rotate"
			case 'p', 'P':
				cmds <- "pause"
			}
		} else if n == 3 {
			if buf[0] == 27 && buf[1] == 91 {
				switch buf[2] {
				case 65:
					cmds <- "up"
				case 66:
					cmds <- "down"
				case 67:
					cmds <- "right"
				case 68:
					cmds <- "left"
				}
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}
