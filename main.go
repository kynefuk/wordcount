package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	if len(os.Args) > 1 {
		input, err := extractInput(os.Args)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(cli.Run(input))
	}
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(cli.Run(string(body)))
}

func extractInput(args []string) (string, error) {
	if len(args) != 2 {
		log.Printf("argument must be only single sentence. Your specified %d arguments", len(args)-1)
		return "", fmt.Errorf("argument must be only single sentence. Your specified %d arguments", len(args)-1)
	}
	input := os.Args[1]
	return input, nil
}
