package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	if terminal.IsTerminal(0) {
		input, err := extractInputSentence(os.Args)
		log.Printf("error: %T", err)
		if err != nil {
			os.Exit(ExitCodeError)
		}
		os.Exit(cli.Run(input))

	} else {
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(ExitCodeError)
		}
		os.Exit(cli.Run(string(body)))
	}
}

func extractInputSentence(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("argument must be only single sentence. Your specified %d arguments", len(args)-1)
	}
	input := os.Args[1]
	return input, nil
}
