package util

import (
	"fmt"
	"os"
)

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
