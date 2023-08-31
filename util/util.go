package util

import (
	"regexp"
	ip "split-clone/input"
	"strconv"
	"strings"
)

var (
	optionValue int64
	prefix      string
)

func GetParam(args []string, ap ip.ArgPosition) ip.Input {
	in := ip.Input{
		Option:       "l",
		OptionValue:  1000,
		SuffixLength: 2,
		FileName:     args[ap.FileName],
		Prefix:       "x",
	}
	if ap.Option > 0 {
		arg := args[ap.Option]
		pattern := `^\d+[kmgKMG]$`
		re := regexp.MustCompile(pattern)
		num := args[ap.Option+1]
		if arg == "-b" && re.MatchString(num) {
			unit := strings.ToUpper(num[len(num)-1:])
			value, _ := strconv.Atoi(num[:len(num)-1])
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
			val, _ := strconv.Atoi(args[ap.Option+1])
			optionValue = int64(val)
		}
		in.Option = strings.TrimPrefix(args[ap.Option], "-")
		in.OptionValue = optionValue
		in.FileName = args[ap.FileName]
		in.Prefix = "x"
		if ap.Prefix > 0 {
			in.Prefix = args[ap.Prefix]
		}
	}
	if ap.AOption > 0 {
		in.SuffixLength, _ = strconv.Atoi(args[ap.AOption+1])
	}
	if ap.Prefix > 0 {
		in.Prefix = args[ap.Prefix]
	}
	return in
}

func GetFilename(name string) string {
	if name == "" {
		return "aa"
	}
	bytes := []byte(name)
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] < 'z' {
			bytes[i]++
			break
		} else {
			bytes[i] = 'a'
			if i == 0 {
				bytes = append([]byte{'a'}, bytes...)
			}
		}
	}
	return string(bytes)
}

func GetDefaultFileName(num int) string {
	return strings.Repeat("a", num)
}
