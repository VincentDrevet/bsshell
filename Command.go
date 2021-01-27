package sbshell

import "io"

type Command struct {
	Name        string
	Description string
	Action      func(w io.Writer, arg string, t *Terminal) error
	//Args        []Argument
	SubCommands []Command
}

func NewCommand(name, description string, action func(w io.Writer, arg string, t *Terminal) error, subs []Command) Command {
	return Command{
		Name:        name,
		Description: description,
		Action:      action,
		//Args:        args,
		SubCommands: subs,
	}
}

func (c *Command) Run(w io.Writer, arg string, t *Terminal) {
	c.Action(w, arg, t)
}

func (c *Command) Match(s string) bool {
	return c.Name == s
}
