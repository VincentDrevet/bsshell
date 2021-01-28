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
