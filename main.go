package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"split-clone/option"
)

func main() {

	// 入力情報を処理

	// オプションの定義
	l := flag.Int("l", 1000, "split at most N lines")
	n := flag.Int("n", 0, "split at most N files")
	b := flag.Int("b", 0, "split at most N bytes")

	flag.Parse()

	// 引数の取得
	args := flag.Args()

	// 引数がない場合はエラーを出力して終了
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "split: missing file operand")
		os.Exit(1)
	}
	// ファイルを読み込む
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// ファイルの情報を取得
	info, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// ファイルサイズを取得
	fileSize := info.Size()
	fmt.Println("filesize: ", fileSize)
	// 分割する単位を決定
	var unit int64
	var line int
	if *l > 0 {
		line = int(*l)
	} else if *n > 0 {
		unit = fileSize / int64(*n)
	} else if *b > 0 {
		unit = int64(*b)
	} else {
		fmt.Fprintln(os.Stderr, "split: invalid option")
		os.Exit(1)
	}
	fmt.Println(unit, line)
	// ファイルを分割する
	// 出力先のファイル名を決定
	var prefix string
	if len(args) > 1 {
		prefix = filepath.Base(args[1])
	}
	fileName := "xaa"
	if prefix != "" {
		fileName = "aa"
	}
	bytes := make([]byte, *b)
	// option.Lines(file, line, prefix, fileName)
	option.Bytes(file, bytes, prefix, fileName)
}
