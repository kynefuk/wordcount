package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Exit codes represent an exit code
const (
	ExitCodeOK = iota
	ExitCodeError
)

// OutPutCount is number of output words
const OutPutCount = 3

// delimiter is newline character or space
var delimiter = regexp.MustCompile(`[\r\n ]+`)

// Set is a pseudo set
type Set map[string]struct{}

// CLI is a command line object
type CLI struct {
	outStream, errStream io.Writer
}

func createSet(splited []string) Set {
	set := Set{}
	for _, v := range splited {
		word := strings.Trim(v, ",.")

		if len(word) == 0 {
			continue
		}
		set[word] = struct{}{}
	}
	return set
}

func outputTop3Word(outStream io.Writer, ranking Ranking) {
	var outputCount = 1
	for _, rank := range ranking {
		if outputCount > OutPutCount {
			break
		}
		fmt.Fprintln(outStream, rank.name)
		outputCount++
	}
}

func extractInputSentence(args []string) (string, error) {
	argsNum := len(args)
	if len(args) != 1 {
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
	splited := delimiter.Split(inputSentence, -1)

	set := createSet(splited)

	var ranking Ranking
	for key := range set {
		count := strings.Count(inputSentence, key)
		rank := Rank{name: key, count: count}
		ranking = append(ranking, rank)
	}

	sort.Slice(ranking, ranking.LessByName)
	sort.Sort(sort.Reverse(ranking))

	outputTop3Word(cli.outStream, ranking)
	return ExitCodeOK
}
