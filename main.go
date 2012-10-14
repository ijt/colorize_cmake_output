package main;

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main() {
	// Reset the coloring if the user hits Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Print(Reset)
		os.Exit(1)
	}()

	// Copy stdin to stdout, turning red when CMake Error lines appear.
	r := bufio.NewReader(os.Stdin)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break;
		}
		if l[0] != ' ' {
			if strings.Contains(l, "CMake Error") {
				fmt.Print(FgRed)
			} else if strings.Contains(l, "CMake Warning") {
				fmt.Print(FgYellow)
			} else {
				fmt.Print(Reset)
			}

		}
		fmt.Print(l)
	}
}

const (
        Reset = "\x1b[0m"
        Bright = "\x1b[1m"
        Dim = "\x1b[2m"
        Underscore = "\x1b[4m"
        Blink = "\x1b[5m"
        Reverse = "\x1b[7m"
        Hidden = "\x1b[8m"

        FgBlack = "\x1b[30m"
        FgRed = "\x1b[31m"
        FgGreen = "\x1b[32m"
        FgYellow = "\x1b[33m"
        FgBlue = "\x1b[34m"
        FgMagenta = "\x1b[35m"
        FgCyan = "\x1b[36m"
        FgWhite = "\x1b[37m"

        BgBlack = "\x1b[40m"
        BgRed = "\x1b[41m"
        BgGreen = "\x1b[42m"
        BgYellow = "\x1b[43m"
        BgBlue = "\x1b[44m"
        BgMagenta = "\x1b[45m"
        BgCyan = "\x1b[46m"
        BgWhite = "\x1b[47m"
)

