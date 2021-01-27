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

func NewTerminal(cmds []Command) Terminal {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s", err.Error())
	}
	return Terminal{
		Prompt:   hostname + "> ",
		Commands: cmds,
		Output:   os.Stdout,
	}
}

/*func (t *Terminal) Run() {
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
			if len(validcmd.Args) == 0 || len(msg) == 1 {
				validcmd.Run(t.Output, "")
			} else {
				arg, err := SearchArgs(validcmd.Args, msg[1])
				if err != nil {
					fmt.Fprintf(t.Output, "%s\n", err.Error())
				} else {
					arg.Run(t.Output)
				}
			}
		}
	}
}
*/
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
