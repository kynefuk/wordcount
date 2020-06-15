package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCliRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	cases := []struct {
		name     string
		input    []string
		out      string
		err      string
		exitCode int
	}{
		{name: "output is correct", input: []string{"dog dog dog. cat. fish fish. go go go go."}, out: "go\ndog\nfish\n", err: "", exitCode: 0},
		{name: "invalid input", input: []string{"bb", "aa"}, out: "", err: "arguments must be single sentence. Your specified 2 arguments", exitCode: 1},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actualCode := cli.Run(c.input)
			if actualCode != c.exitCode {
				t.Errorf("expected %d, but got: %d", c.exitCode, actualCode)
			}
			if !strings.Contains(outStream.String(), c.out) {
				t.Errorf("expected %s, but got: %s", c.out, outStream.String())
			}
			if !strings.Contains(errStream.String(), c.err) {
				t.Errorf("expected %s, but got: %s", c.err, errStream.String())
			}
		})
	}
}
