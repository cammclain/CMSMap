package scans

import (
	"cmsmap/internal/commands"
	"cmsmap/pkg/scans/reconscans"
	"fmt"
)

type SubdomainScanCommand struct{}

func (cmd *SubdomainScanCommand) Handle(args []string) string {
	if len(args) < 2 {
		return "Usage: scan_subdomains <domain>"
	}
	domain := args[1]

	// Asynchronous call to subdomain enumeration to avoid blocking the REPL
	go reconscans.Enumerate_Subdomains(domain)

	return fmt.Sprintf("Initiated subdomain scan for %s", domain)
}

func init() {
	commands.RegisterCommand("scan_subdomains", new(SubdomainScanCommand))
}
