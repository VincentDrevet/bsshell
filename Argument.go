package sbshell

import "io"

type Argument struct {
	Name   string
	Action func(w io.Writer) error
}

func (a *Argument) Run(w io.Writer) {
	a.Action(w)
}

func NewArgument(name string, action func(w io.Writer) error) Argument {
	return Argument{
		Name:   name,
		Action: action,
	}
}
