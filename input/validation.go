package input

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	optionValue int64
)

func CheckInput(args []string) (e Exist, err error) {
	e = Exist{}
	// 引数がない場合はエラーを出力して終了
	if len(args[1:]) == 0 {
		return e, errors.New("split: missing file operand")
	}
	for i, arg := range args[1:] {
		// すでにoptionとfileNameがある場合はprefixとみなす
		if e.Option && e.FileName {
			// prefixがすでにある場合はエラー
			if e.Prefix {
				return e, errors.New("Too many arguments")
			}
			e.Prefix = true
			continue
		}
		// optionかどうか
		if strings.HasPrefix(arg, "-") {
			if e.Option || e.FileName {
				return e, errors.New("Invalid option")
			}
			e.Option = true
			// optionは、-l, -n, -bのいずれか
			if arg != "-l" && arg != "-n" && arg != "-b" {
				return e, errors.New("Invalid option 2")
			}
			// optionの次の引数があるか
			if i+2 >= len(args) {
				fmt.Println("optionsss: ", arg, i, i+2, len(args))
				return e, errors.New("Need number after option")
			}
			// optionの次は整数またはオプションが-bの時はk,m,gをつけた整数
			pattern := `^\d+$`
			if arg == "-b" {
				pattern = `^\d+[kmgKMG]?$`
			}
			re := regexp.MustCompile(pattern)
			if !re.MatchString(args[i+2]) {
				return e, errors.New("Invalid option number")
			}
		} else {
			if e.Option && i == 1 {
				continue
			}
			e.FileName = true
		}
	}
	if !e.FileName {
		return e, errors.New("split: missing file operand")
	}
	return e, nil
}

func GetParam(args []string, e Exist) (input Input) {
	if e.Option {
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
			FileName:    args[3],
		}
	} else {
		input = Input{Option: "l", OptionValue: 1000, FileName: args[1]}
	}
	if e.Prefix {
		input.Prefix = args[4]
	}
	return input
}
