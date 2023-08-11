package main

import (
	"fmt"
	"os"
	"time"

	in "split-clone/input"
	op "split-clone/option"
	vl "split-clone/validation"
)

var (
	fileName string
	input    in.Input
	exist    in.Exist
	prefix   string
)

func main() {
	start := time.Now()
	// 入力情報を処理
	args := os.Args
	optionExist, err := vl.CheckInput(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input = vl.GetParam(args, optionExist)
	// ファイルを読み込む
	file, err := os.Open(input.FileName)
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
	// 分割する単位を決定
	var unit int64
	var line int64
	// ファイルを分割する
	// 出力先のファイル名を決定

	fileName := "xaa"
	if input.Prefix != "" {
		fileName = "aa"
	}
	// fmt.Println(prefix, fileName)
	// os.Exit(1)
	switch input.Option {
	case "l":
		line = input.OptionValue
		op.Lines(file, line, prefix, fileName)
	case "n":
		// ファイルサイズを取得
		fileSize := info.Size()
		unit = fileSize / int64(input.OptionValue)
		bytes := make([]byte, input.OptionValue)
		op.Bytes(file, bytes, prefix, fileName)
	case "b":
		unit = int64(input.OptionValue)
		bytes := make([]byte, input.OptionValue)
		op.Bytes(file, bytes, prefix, fileName)
	default:
		fmt.Fprintln(os.Stderr, "split: invalid option")
		os.Exit(1)
	}
	fmt.Println(unit, line)
	elapsed := time.Since(start)
	fmt.Printf("elapsed time: %v\n", elapsed)
	os.Exit(1)
}
