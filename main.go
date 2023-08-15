package main

import (
	"fmt"
	"os"
	"time"

	in "split-clone/input"
	sp "split-clone/split"
	ut "split-clone/util"
)

var (
	fileName string
	input    in.Input
	prefix   string
	e        in.ArgPosition
)

func main() {
	start := time.Now()
	// 入力情報を処理
	args := os.Args
	e, err := in.ValidateInput(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	input = ut.GetParam(args, e)
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
