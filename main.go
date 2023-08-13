package main

import (
	"fmt"
	"os"
	"time"

	in "split-clone/input"
	sp "split-clone/split"
)

var (
	fileName string
	input    in.Input
	prefix   string
	e        in.Exist
)

func main() {
	start := time.Now()
	// 入力情報を処理
	args := os.Args
	e, err := in.CheckInput(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input = in.GetParam(args, e.Option)
	// ファイルを読み込む
	file, err := os.Open(input.FileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	// ファイルを分割する
	err = sp.Split(input, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed time: %v\n", elapsed)
}
