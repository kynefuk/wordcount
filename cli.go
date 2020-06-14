package main

import "io"

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	// メイン処理
	return 0
}
