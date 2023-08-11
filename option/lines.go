package option

// var lineCount int64 = 0

// func Lines(file *os.File, line int64, prefix string, fileName string) {
// 	scanner := bufio.NewScanner(file)
// 	outFile, err := os.Create(fmt.Sprintf(prefix + fileName))
// 	for scanner.Scan() {
// 		lineCount++
// 		if lineCount > line {
// 			lineCount = 1
// 			outFile.Close()
// 			fileName = util.GetFilename(fileName)
// 			outFile, err = os.Create(fmt.Sprintf(prefix + fileName))
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, err)
// 				os.Exit(1)
// 			}
// 			defer outFile.Close()
// 		}
// 		outFile.WriteString(scanner.Text() + "\n")
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }
