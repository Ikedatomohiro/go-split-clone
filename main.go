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
	exist    in.Exist
	prefix   string
)

func main() {
	start := time.Now()
	// 入力情報を処理
	args := os.Args
	optionExist, err := in.CheckInput(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input = in.GetParam(args, optionExist)
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
