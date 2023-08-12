package input

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	optionExist = false
	optionValue int64
)

func CheckInput(args []string) (optionExist bool, err error) {
	// 引数がない場合はエラーを出力して終了
	if len(args[1:]) == 0 {
		return false, errors.New("split: missing file operand")
	}
	fileNameExist := false
	for i, arg := range args[1:] {
		// すでにoptionとfileNameがある場合はエラー
		if optionExist && fileNameExist {
			return optionExist, errors.New("Too many arguments")
		}
		// optionかどうか
		if strings.HasPrefix(arg, "-") {
			if optionExist || fileNameExist {
				return optionExist, errors.New("Invalid ption")
			}
			optionExist = true
			// optionは、-l, -n, -bのいずれか
			if arg != "-l" && arg != "-n" && arg != "-b" {
				return optionExist, errors.New("Invalid option 2")
			}
			// optionの次の引数があるか
			if i+2 >= len(args) {
				fmt.Println("optionsss: ", arg, i, i+2, len(args))
				return optionExist, errors.New("Need number after option")
			}
			// optionの次は整数またはオプションが-bの時はk,m,gをつけた整数
			pattern := `^\d+$`
			if arg == "-b" {
				pattern = `^\d+[kmgKMG]?$`
			}
			re := regexp.MustCompile(pattern)
			if !re.MatchString(args[i+2]) {
				return optionExist, errors.New("Invalid option number")
			}
		} else {
			if optionExist && i == 1 {
				continue
			}
			fileNameExist = true
		}
	}
	return optionExist, nil
}

func GetParam(args []string, optionExsits bool) (input Input) {
	if optionExsits {
		arg := args[1]
		pattern := `^\d+[kmgKMG]$`
		re := regexp.MustCompile(pattern)

		if arg == "-b" && re.MatchString(args[2]) {
			unit := strings.ToUpper(args[2][len(args[2])-1:])
			value, _ := strconv.Atoi(args[2][:len(args[2])-1])
			switch unit {
			case "K":
				optionValue = int64(value * 1000)
			case "M":
				optionValue = int64(value * 1000000)
			case "G":
				optionValue = int64(value * 1000000000)
			default:
				optionValue = int64(value)
			}
		} else {
			val, _ := strconv.Atoi(args[2])
			optionValue = int64(val)
		}
		input = Input{
			Option:      string([]byte{arg[1]}),
			OptionValue: optionValue,
			Uint:        args[2],
			FileName:    args[3],
		}
	} else {
		input = Input{Option: "l", OptionValue: 1000, FileName: args[1]}
	}
	return input
}
