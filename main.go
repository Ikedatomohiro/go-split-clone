package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	// 入力情報を処理

	// ファイルを読み込む

	// ファイルを分割する

	// オプションの定義
	l := flag.Int("l", 1000, "split at most N lines")
	n := flag.Int("n", 0, "split at most N files")
	b := flag.Int("b", 0, "split at most N bytes")

	flag.Parse()

	// 引数の取得
	args := flag.Args()

	// 引数がない場合はエラーを出力して終了
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "split: missing file operand")
		os.Exit(1)
	}

	// ファイルを開く
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// ファイルの情報を取得
	info, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// ファイルサイズを取得
	fileSize := info.Size()
	fmt.Println("filesize: ", fileSize)
	// 分割する単位を決定
	var unit int64
	if *l > 0 {
		unit = int64(*l)
	} else if *n > 0 {
		unit = fileSize / int64(*n)
	} else if *b > 0 {
		unit = int64(*b)
	} else {
		fmt.Fprintln(os.Stderr, "split: invalid option")
		os.Exit(1)
	}
	fmt.Println(unit)

	// 出力先のファイル名を決定
	fileName := filepath.Base(args[0])
	if fileName == "." {
		fileName = "x"
	}
	name := "xaa"
	fmt.Println("current: ", name)
	name = getFilename(name)
	fmt.Println("next: ", name)

	// option.Lines()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	fileCount := 0
	outFile, err := os.Create(fmt.Sprintf("out_%d.txt", fileCount))
	for scanner.Scan() {
		lineCount++

		if lineCount > *l {
			lineCount = 1
			fileCount++
			outFile.Close()
			outFile, err = os.Create(fmt.Sprintf("out_%d.txt", fileCount))
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

	fmt.Println(lineCount)
	os.Exit(100000)
	// ファイルを分割
	reader := bufio.NewReader(file)
	var i int
	var j int64
	for {
		os.Exit(100000)
		// 出力先のファイルを作成
		outFileName := fmt.Sprintf("%02d", i)
		outFile, err := os.Create(outFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer outFile.Close()

		// ファイルに書き込み
		var k int64
		for k = 0; k < unit; k++ {
			b, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(b)
			// outFile.WriteByte(b)
			j++
		}

		// ファイルを閉じる
		// outFile.Close()

		// ファイルの終端に到達した場合は終了
		if j == fileSize {
			break
		}

		// ファイルを分割する単位を再計算
		if *l > 0 {
			unit = int64(*l)
		} else if *n > 0 {
			unit = (fileSize - j) / int64(*n-i)
		} else if *b > 0 {
			unit = int64(*b)
		}

		// 次のファイルに進む
		i++
	}

	// 	// txtar形式で出力
	// 	// fmt.Printf("main.go--\n%s", txtar("*.go"))
	// }

	// // txtar形式で出力するための関数
	//
	//	func txtar(pattern string) string {
	//		var buf []byte
	//		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	//			if err != nil {
	//				return err
	//			}
	//			if info.IsDir() || !filepath.Match(pattern, info.Name()) {
	//				return nil
	//			}
	//			buf = append(buf, fmt.Sprintf("%s--\n", path)...)
	//			file, err := os.Open(path)
	//			if err != nil {
	//				return err
	//			}
	//			defer file.Close()
	//			reader := bufio.NewReader(file)
	//			for {
	//				line, err := reader.ReadString('\n')
	//				if err != nil {
	//					if err == io.EOF {
	//						break
	//					}
	//					return err
	//				}
	//				buf = append(buf, []byte(line)...)
	//			}
	//			buf = append(buf, "\n"...)
	//			return nil
	//		})
	//		return string(buf)
}

func getFilename(name string) string {
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
