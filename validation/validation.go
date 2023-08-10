package validation

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetOptionParam() (fileName string, l *int, n *int, b *int) {
	fileName, err := checkInput()
	if err != nil {
		fmt.Println("Invalid input . error: ", err)
		os.Exit(1)
	}
	l, n, b = setDefaultValue()
	return fileName, l, n, b
}

func checkInput() (fileName string, err error) {
	args := os.Args
	// 引数がない場合はエラーを出力して終了
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "split: missing file operand")
		os.Exit(1)
	}
	optionExist := false
	fileNameExist := false
	for i, arg := range args[1:] {
		// すでにoptionとfileNameがある場合はエラー
		if optionExist && fileNameExist {
			return "", errors.New("Too many arguments")
		}
		// optionかどうか
		if strings.HasPrefix(arg, "-") {
			if optionExist || fileNameExist {
				return "", errors.New("Invalid ption")
			}
			optionExist = true
			// optionは、-l, -n, -bのいずれか
			if arg != "-l" && arg != "-n" && arg != "-b" {
				return "", errors.New("Invalid option 2")
			}
			// optionの次の引数があるか
			if len(args) < 0 {
				return "", errors.New("Need number after option")
			}
			// optionの次は整数
			if _, err := strconv.Atoi(args[i+2]); err != nil {
				return "", errors.New("Invalid option number")
			}
		} else {
			if optionExist && i == 1 {
				continue
			}
			fileNameExist = true
			fileName = arg
		}
	}
	return fileName, nil
}

func setDefaultValue() (l *int, n *int, b *int) {
	l = flag.Int("l", 0, "split at most N lines")
	n = flag.Int("n", 0, "split at most N files")
	b = flag.Int("b", 0, "split at most N bytes")
	flag.Parse()
	count := 0
	if *l != 0 {
		count++
	}
	if *n != 0 {
		count++
	}
	if *b != 0 {
		count++
	}
	if count == 0 {
		*l = 1000
	}
	return l, n, b
}
