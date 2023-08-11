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
	// ファイルを分割する
	switch input.Option {
	case "l":
		op.Lines(input, file, fileName)
	case "n":
		op.Numbers(input, info, file, fileName)
	case "b":
		op.Bytes(input, file, fileName)
	default:
		fmt.Fprintln(os.Stderr, "split: invalid option")
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed time: %v\n", elapsed)
}
