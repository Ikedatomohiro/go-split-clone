package option

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"split-clone/util"
)

func Bytes(file *os.File, bytes []byte, prefix string, fileName string) {
	reader := bufio.NewReader(file)
	for {
		bytes := make([]byte, 3000)
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
			break
		}

		err = ioutil.WriteFile(prefix+fileName, bytes, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fileName = util.GetFilename(fileName)
	}
}
