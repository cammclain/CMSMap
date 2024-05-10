package scans

import (
	"cmsmap/internal/commands"
)

type ScanCommand struct{}

func (s *ScanCommand) Handle(args []string) string {
	if len(args) < 2 {
		return "Usage: scan <url>"
	}
	// Implement the scanning logic here
	return "Scanning " + args[1]
}

func init() {
	commands.RegisterCommand("scan", new(ScanCommand))
}
