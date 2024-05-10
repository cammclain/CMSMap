package reconscans

import (
	"bytes"
	"context"
	"io"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

// Modified to return the result and error
func Enumerate_Subdomains(domain string) (string, error) {
	subfinderOptions := &runner.Options{
		Threads:            10,
		Timeout:            10,
		MaxEnumerationTime: 10,
	}

	subfinder, err := runner.NewRunner(subfinderOptions)
	if err != nil {
		return "", err
	}

	subfinderOutput := &bytes.Buffer{}

	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{subfinderOutput}); err != nil {
		return "", err
	}

	return subfinderOutput.String(), nil
}
