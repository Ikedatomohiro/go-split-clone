package util

import (
	"fmt"
	"os"
	"regexp"
	ip "split-clone/input"
	"strconv"
	"strings"
)

var (
	optionValue int64
	prefix      string
)

func GetParam(args []string, e ip.Exist) (in ip.Input) {
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
		in = ip.Input{
			Option:      string([]byte{arg[1]}),
			OptionValue: optionValue,
			FileName:    args[3],
		}
		if e.Prefix {
			in.Prefix = args[4]
		}
	} else {
		in = ip.Input{Option: "l", OptionValue: 1000, FileName: args[1]}
		if e.Prefix {
			in.Prefix = args[2]
		}
	}
	fmt.Println("o: ", e)
	fmt.Println("in: ", in)
	return in
}

func GetFilename(name string) string {
	if name == "" {
		return "aaa"
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

func ShowUsage() {
	fmt.Println("Usage: split [OPTION]... [FILE] [PREFIX]")
	os.Exit(1)
}
