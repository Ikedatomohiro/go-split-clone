package option

import (
	"bufio"
	"fmt"
	"io"
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

func Lines(in in.Input, file *os.File) error {
	scanner := bufio.NewScanner(file)
	fileName = util.GetDefaultFileName(in.SuffixLength)
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
				return err
			}
			defer outFile.Close()
		}
		outFile.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Bytes(in in.Input, file *os.File) error {
	reader := bufio.NewReader(file)
	fileName = util.GetDefaultFileName(in.SuffixLength)
	prefix = in.Prefix
	unit = int64(in.OptionValue)
	bytes := make([]byte, in.OptionValue)
	for {
		end, err := outputBytes(reader, bytes, prefix, fileName)
		if err != nil {
			return err
		}
		if end {
			break
		}
		fileName = util.GetFilename(fileName)
	}
	return nil
}

func Numbers(in in.Input, file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := info.Size()
	unit = fileSize / int64(in.OptionValue)
	remainder := fileSize % int64(in.OptionValue)
	bytes := make([]byte, unit)
	reader := bufio.NewReader(file)
	fileName = util.GetDefaultFileName(in.SuffixLength)
	prefix = in.Prefix
	lineCount = 0
	for lineCount < in.OptionValue {
		if lineCount == in.OptionValue-1 {
			bytes = make([]byte, unit+remainder)
		}
		end, err := outputBytes(reader, bytes, prefix, fileName)
		if err != nil {
			return err
		}
		if end {
			break
		}
		fileName = util.GetFilename(fileName)
		lineCount++
	}
	return nil
}

func outputBytes(reader *bufio.Reader, bytes []byte, prefix string, fileName string) (end bool, err error) {
	n, err := io.ReadFull(reader, bytes)
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			// EOFに達し、要求したバイト数が読み取られなかった場合
			err = ioutil.WriteFile(prefix+fileName, bytes[:n], 0644)
			if err != nil {
				return true, err
			}
		} else if err != io.EOF {
			return false, err
		}
		return true, nil
	}
	err = ioutil.WriteFile(prefix+fileName, bytes, 0644)
	if err != nil {
		return false, err
	}
	return false, nil
}
