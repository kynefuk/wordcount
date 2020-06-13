//標準入力から英語の文章を受け取り単語の数を数え、多い順TOP3を改行して出力するプログラムを作る
//単語の区切りはスペースまたは改行とする。なお連続するスペースや改行、文末のドットやカンマは無視する。
//同じ単語数の単語がある場合はアルファベット順に出力する。
package main

import (
	"os"
)

func main() {
	os.Exit(Run(os.Args))
}

func Run(args []string) int {
	//処理
	return 0
}
