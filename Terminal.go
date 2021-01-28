package sbshell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Terminal struct {
	Prompt   string
	Commands []Command
	Output   io.Writer
}

func NewTerminal() Terminal {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s", err.Error())
	}

	Exitcmd := NewCommand("exit", "exit shell", func(w io.Writer, arg string, t *Terminal) error {
		os.Exit(0)
		return nil
	}, nil)

	Clearcmd := NewCommand("clear", "clear terminal screen", func(w io.Writer, arg string, t *Terminal) error {
		fmt.Fprintf(w, "\x1b[2J\x1b[H")
		return nil
	}, nil)

	return Terminal{
		Prompt:   hostname + "> ",
		Output:   os.Stdout,
		Commands: []Command{Clearcmd, Exitcmd},
	}
}

func (t *Terminal) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprintf(t.Output, t.Prompt)
		scanner.Scan()
		msg := strings.Split(string(scanner.Bytes()), " ")

		cmd := msg[0]
		validcmd, err := SearchCommand(t.Commands, cmd)
		if err != nil {
			fmt.Fprintf(t.Output, "%s\n", err.Error())
		} else {
			// On vérifie si la commande prend des arguments si elle n'en prend pas on n'existe la commande
			if len(validcmd.SubCommands) == 0 || len(msg) == 1 {
				validcmd.Run(t.Output, "", t)
			} else {
				sub, err := SearchCommand(validcmd.SubCommands, msg[1])
				if err != nil {
					fmt.Fprintf(t.Output, "%s\n", err.Error())
				} else {
					sub.Run(t.Output, "", t)
				}
			}
		}
	}
}

func (t *Terminal) SetPrompt(prompt string) {
	t.Prompt = prompt
}

func (t *Terminal) SetOutput(w io.Writer) {
	t.Output = w
}

func (t *Terminal) AddCommand(cmd Command) {
	t.Commands = append(t.Commands, cmd)
}
