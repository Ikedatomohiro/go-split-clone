package option

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	in "split-clone/input"
	"split-clone/util"
)

var (
	lineCount int64
	unit      int64
	line      int64
	prefix    string
	fileName  string
)

func Lines(in in.Input, file *os.File) {
	scanner := bufio.NewScanner(file)
	fileName = setDefaultFileName(in)
	prefix = in.Prefix
	outFile, err := os.Create(fmt.Sprintf(prefix + fileName))
	line = in.OptionValue
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

func Bytes(in in.Input, file *os.File) {
	reader := bufio.NewReader(file)
	fileName = setDefaultFileName(in)
	prefix = in.Prefix
	unit = int64(in.OptionValue)
	bytes := make([]byte, in.OptionValue)
	for {
		end := outputBytes(reader, bytes, prefix, fileName)
		if end {
			break
		}
		fileName = util.GetFilename(fileName)
	}
}

func Numbers(in in.Input, info fs.FileInfo, file *os.File) {
	fileSize := info.Size()
	unit = fileSize / int64(in.OptionValue)
	remainder := fileSize % int64(in.OptionValue)
	bytes := make([]byte, unit)
	reader := bufio.NewReader(file)
	fileName = setDefaultFileName(in)
	prefix = in.Prefix
	lineCount = 0
	for lineCount < in.OptionValue {
		if lineCount == in.OptionValue-1 {
			bytes = make([]byte, unit+remainder)
		}
		end := outputBytes(reader, bytes, prefix, fileName)
		if end {
			break
		}
		fileName = util.GetFilename(fileName)
		lineCount++
	}
}

func outputBytes(reader *bufio.Reader, bytes []byte, prefix string, fileName string) (end bool) {
	n, err := io.ReadFull(reader, bytes)
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			// EOFに達し、要求したバイト数が読み取られなかった場合
			err = ioutil.WriteFile(prefix+fileName, bytes[:n], 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else if err != io.EOF {
			// EOFでない他のエラー
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return true
	}
	err = ioutil.WriteFile(prefix+fileName, bytes, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return false
}

func setDefaultFileName(in in.Input) string {
	fileName := "xaa"
	if in.Prefix != "" {
		fileName = "aa"
	}
	if fileName == "" {
		return "aaa"
	}
	return fileName
}
