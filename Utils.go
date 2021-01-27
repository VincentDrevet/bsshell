package sbshell

import "errors"

func SearchCommand(availablecommands []Command, inputuser string) (Command, error) {
	for _, cmd := range availablecommands {
		if cmd.Match(inputuser) {
			return cmd, nil
		}
	}
	return Command{}, errors.New("Unknown command")
}

func SearchArgs(args []Argument, arg string) (Argument, error) {
	for _, e := range args {
		if e.Name == arg {
			return e, nil
		}
	}
	return Argument{}, errors.New("Unknown Argument")
}
