package commands

import (
	"sync"
)

type CommandHandler interface {
	Handle(args []string) string
}

var (
	commands = make(map[string]CommandHandler)
	lock     sync.Mutex
)

func RegisterCommand(name string, handler CommandHandler) {
	lock.Lock()
	defer lock.Unlock()
	commands[name] = handler
}

func GetHandler(commandName string) (CommandHandler, bool) {
	lock.Lock()
	defer lock.Unlock()
	handler, found := commands[commandName]
	return handler, found
}

func init() {
	// Initialization of all commands
	// Example: RegisterCommand("scan", new(ScanCommand))
}
