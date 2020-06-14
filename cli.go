//標準入力から英語の文章を受け取り単語の数を数え、多い順TOP3を改行して出力するプログラムを作る
//単語の区切りはスペースまたは改行とする。なお連続するスペースや改行、文末のドットやカンマは無視する。
//同じ単語数の単語がある場合はアルファベット順に出力する。

package main

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Delimiter is newline character or space
var delimiter = regexp.MustCompile(`[\r\n ]+`)

const OUTPUT_COUNT = 3

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(arg string) int {
	inputSentence := arg
	splited := delimiter.Split(inputSentence, -1)

	set := make(map[string]struct{})
	for _, v := range splited {
		word := strings.Trim(v, ",.")

		if len(word) == 0 {
			continue
		}
		set[word] = struct{}{}
	}

	var ranking Ranking
	for key := range set {
		count := strings.Count(inputSentence, key)
		rank := Rank{name: key, count: count}
		ranking = append(ranking, rank)
	}

	sort.Slice(ranking, ranking.LessByName)
	sort.Sort(sort.Reverse(ranking))

	var outputCount = 1
	for _, rank := range ranking {
		if outputCount > OUTPUT_COUNT {
			break
		}
		fmt.Println(rank.name)
		outputCount++
	}
	return 0
}
