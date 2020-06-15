package main

import (
	"flag"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	if terminal.IsTerminal(0) {
		flag.Parse()
		args := flag.Args()
		os.Exit(cli.Run(args))
	} else {
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(ExitCodeError)
		}
		args := []string{string(body)}
		os.Exit(cli.Run(args))
	}
}
