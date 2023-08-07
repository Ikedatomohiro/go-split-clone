package option

import (
	"bufio"
	"fmt"
	"os"
	"split-clone/util"
)

func Lines(file *os.File, line int, prefix string, fileName string) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	outFile, err := os.Create(fmt.Sprintf(prefix + fileName))
	for scanner.Scan() {
		lineCount++

		if lineCount > line {
			lineCount = 1
			outFile.Close()
			fileName = util.GetFilename(fileName)
			outFile, err = os.Create(fmt.Sprintf(prefix + fileName))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer outFile.Close()
		}
		outFile.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
