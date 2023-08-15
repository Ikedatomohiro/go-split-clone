package input

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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
