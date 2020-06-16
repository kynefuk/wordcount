package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Exit codes represent an exit status code
const (
	ExitCodeOK = iota
	ExitCodeError
)

// OutPutCount is number of output words
const OutPutCount = 3

// Delimiter is newline character or space
var Delimiter = regexp.MustCompile(`[\r\n ]+`)

// Set is a pseudo set object
type Set map[string]struct{}

// CLI is a command line object
type CLI struct {
	outStream, errStream io.Writer
}

func createSet(splited []string) Set {
	set := Set{}
	for _, v := range splited {
		word := strings.Trim(v, ",.")

		// if word is like empty character
		if len(word) == 0 {
			continue
		}
		set[word] = struct{}{}
	}
	return set
}

func outputTop3Word(outStream io.Writer, ranking Ranking) {
	var count = 1
	for _, rank := range ranking {
		if count > OutPutCount {
			break
		}
		fmt.Fprintln(outStream, rank.name)
		count++
	}
}

func extractInputSentence(args []string) (string, error) {
	argsNum := len(args)
	if argsNum != 1 {
		return "", fmt.Errorf("arguments must be single sentence. Your specified %d arguments", argsNum)
	}
	input := args[0]
	return input, nil
}

// Run is a main process of cli
func (cli *CLI) Run(args []string) int {
	inputSentence, err := extractInputSentence(args)
	if err != nil {
		fmt.Fprintf(cli.errStream, "error occured: %s\n", err)
		return ExitCodeError
	}
	splited := Delimiter.Split(inputSentence, -1)

	set := createSet(splited)

	var ranking Ranking
	for key := range set {
		count := strings.Count(inputSentence, key)
		rank := Word{name: key, count: count}
		ranking = append(ranking, rank)
	}

	sort.Slice(ranking, ranking.LessByName)
	sort.Sort(sort.Reverse(ranking))

	outputTop3Word(cli.outStream, ranking)
	return ExitCodeOK
}
