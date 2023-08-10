package main

import (
	"fmt"
	"os"
	"split-clone/validation"
)

func main() {

	// 入力情報を処理
	fileName, l, n, b := validation.GetOptionParam()
	fmt.Println(fileName, *l, *n, *b)
	os.Exit(1)
	// ファイルを読み込む
	file, err := os.Open(fileName)
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
	os.Exit(1)
	// ファイルサイズを取得
	fileSize := info.Size()
	fmt.Println("filesize: ", fileSize)
	// 分割する単位を決定
	var unit int64
	var line int
	// if *l > 0 {
	// 	line = int(*l)
	// } else if *n > 0 {
	// 	unit = fileSize / int64(*n)
	// } else if *b > 0 {
	// 	unit = int64(*b)
	// } else {
	// 	fmt.Fprintln(os.Stderr, "split: invalid option")
	// 	os.Exit(1)
	// }
	fmt.Println(unit, line)
	// ファイルを分割する
	// 出力先のファイル名を決定
	// var prefix string
	// if len(args) > 1 {
	// 	prefix = filepath.Base(args[1])
	// }
	// fileName := "xaa"
	// if prefix != "" {
	// 	fileName = "aa"
	// }
	// bytes := make([]byte, *b)
	// option.Lines(file, line, prefix, fileName)
	// option.Bytes(file, bytes, prefix, fileName)
}
