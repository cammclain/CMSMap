package help

import (
	"cmsmap/internal/commands"
	"fmt"
)

type HelpCommand struct{}

func (h *HelpCommand) Handle(args []string) string {
	if len(args) > 1 {
		if handler, ok := commands.GetHandler(args[1]); ok {
			return handler.Help()
		}
		return "No help available for " + args[1]
	}
	return h.Help()
}

func (h *HelpCommand) Help() string {
	helpText := "Available commands:\n"
	for cmdName, handler := range commands.GetAllHandlers() {
		helpText += fmt.Sprintf("%s - %s\n", cmdName, handler.Help())
	}
	return helpText
}

func init() {
	commands.RegisterCommand("help", new(HelpCommand))
}
