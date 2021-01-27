package main

import (
	"fmt"
	"io"
	"os"

	sbshell "github.com/VincentDrevet/bsshell"
)

func main() {
	// Create an exit command for exit terminal
	ExitCmd := sbshell.NewCommand("exit", "Exit terminal", func(w io.Writer, arg string, t *sbshell.Terminal) error {
		os.Exit(0)
		return nil
	}, nil)

	HiCmd := sbshell.NewCommand("hi", "salutation", func(w io.Writer, arg string, t *sbshell.Terminal) error {
		fmt.Fprintf(w, "Hi ;)\n")
		return nil
	}, nil)

	// Init a new terminal and attach command to it
	t := sbshell.NewTerminal([]sbshell.Command{ExitCmd, HiCmd})

	// Run the terminal
	t.Run()

}
