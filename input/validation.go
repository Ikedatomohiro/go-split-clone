package input

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateInput(args []string) (ap ArgPosition, err error) {
	ap = ArgPosition{}
	// 引数がない場合はエラーを出力して終了
	if len(args) == 1 {
		return ap, errors.New("split: missing file operand")
	}
	for i := 1; i < len(args); i++ {
		arg := args[i]
		// すでにoptionとfileNameがある場合はprefixとみなす
		if ap.Option > 0 && ap.FileName > 0 {
			// prefixがすでにある場合はエラー
			if ap.Prefix > 0 {
				return ap, errors.New("Too many arguments")
			}
			ap.Prefix = i
			continue
		}
		// optionかどうか
		if strings.HasPrefix(arg, "-") {
			if ap.Option > 0 || ap.FileName > 0 {
				return ap, errors.New("Invalid option")
			}
			ap.Option = i
			// optionは、-l, -n, -bのいずれか
			if arg != "-l" && arg != "-n" && arg != "-b" {
				return ap, errors.New("Invalid option 2")
			}
			// optionの次の引数があるか
			if i > len(args) {
				return ap, errors.New("Need number after option")
			}
			// optionの次は整数またはオプションが-bの時はk,m,gをつけた整数
			pattern := `^\d+$`
			if arg == "-b" {
				pattern = `^\d+[kmgKMG]?$`
			}
			re := regexp.MustCompile(pattern)
			if !re.MatchString(args[i+1]) {
				return ap, errors.New("Invalid option number")
			}
		} else {
			if ap.Option > 0 && i == ap.Option+1 {
				continue
			}
			if ap.FileName > 0 {
				ap.Prefix = i
				continue
			}
			ap.FileName = i
		}
	}
	if ap.FileName < 1 {
		return ap, errors.New("split: missing file operand")
	}
	return ap, nil
}
