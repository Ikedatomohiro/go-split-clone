package option_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	in "split-clone/input"
	op "split-clone/option"
	"testing"
)

type Result struct {
	fileContents map[string]string
}

func TestLines(t *testing.T) {
	var tests = []struct {
		name        string
		fileContent string
		args        in.Input
		wantResult  Result
		wantErr     bool
	}{
		{
			name:        "success",
			fileContent: "test samples\ntest",
			args: in.Input{
				Option:       "l",
				OptionValue:  1,
				SuffixLength: 2,
				FileName:     "test.txt",
				Prefix:       "x",
			},
			wantResult: Result{
				fileContents: map[string]string{
					"xaa": "test samples\n",
					"xab": "test\n",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックファイルを作成
			mock, err := os.CreateTemp("./", "test*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(mock.Name())
			defer mock.Close()
			_, err = io.WriteString(mock, tt.fileContent)
			if err != nil {
				fmt.Println("ファイルへの書き込みに失敗:", err)
				return
			}
			mock.Seek(0, io.SeekStart)
			// テスト実施
			err = op.Lines(tt.args, mock)
			if err != nil {
				t.Errorf("Lines returned an error: %v", err)
			}
			for file, value := range tt.wantResult.fileContents {
				content, err := ioutil.ReadFile(file)
				defer os.Remove(file)
				if err != nil {
					t.Fatalf("Failed to read temp file: %v", err)
				}
				// ファイルが存在してもファイル名と内容が一致していなければエラーとなる
				if !tt.wantErr && string(content) != value {
					t.Errorf("Lines wrote %q, expected %q", content, value)
				}
			}

		})
	}
}

func TestBytes(t *testing.T) {
	var tests = []struct {
		name        string
		fileContent string
		args        in.Input
		wantResult  Result
		wantErr     bool
	}{
		{
			name:        "success",
			fileContent: "test samples",
			args: in.Input{
				Option:       "b",
				OptionValue:  5,
				SuffixLength: 2,
				FileName:     "test.txt",
				Prefix:       "x",
			},
			wantResult: Result{
				fileContents: map[string]string{
					"xaa": "test ",
					"xab": "sampl",
					"xac": "es",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックファイルを作成
			mock, err := os.CreateTemp("./", "test*.txt")

			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(mock.Name())
			defer mock.Close()
			_, err = io.WriteString(mock, tt.fileContent)
			if err != nil {
				fmt.Println("ファイルへの書き込みに失敗:", err)
				return
			}
			mock.Seek(0, io.SeekStart)
			// テスト実施
			err = op.Bytes(tt.args, mock)
			if err != nil {
				t.Errorf("Bytes returned an error: %v", err)
			}
			for file, value := range tt.wantResult.fileContents {
				content, err := ioutil.ReadFile(file)
				defer os.Remove(file)
				if err != nil {
					t.Fatalf("Failed to read temp file: %v", err)
				}
				// ファイルが存在してもファイル名と内容が一致していなければエラーとなる
				if !tt.wantErr && string(content) != value {
					t.Errorf("Bytes wrote %q, expected %q", content, value)
				}
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	var tests = []struct {
		name        string
		fileContent string
		args        in.Input
		wantResult  Result
		wantErr     bool
	}{
		{
			name:        "success",
			fileContent: "test samples test",
			args: in.Input{
				Option:       "n",
				OptionValue:  2,
				SuffixLength: 2,
				FileName:     "test.txt",
				Prefix:       "x",
			},
			wantResult: Result{
				fileContents: map[string]string{
					"xaa": "test sam",
					"xab": "ples test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックファイルを作成
			mock, err := os.CreateTemp("./", "test*.txt")

			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(mock.Name())
			defer mock.Close()
			_, err = io.WriteString(mock, tt.fileContent)
			if err != nil {
				fmt.Println("ファイルへの書き込みに失敗:", err)
				return
			}
			mock.Seek(0, io.SeekStart)
			// テスト実施
			err = op.Numbers(tt.args, mock)
			if err != nil {
				t.Errorf("Numbers returned an error: %v", err)
			}
			for file, value := range tt.wantResult.fileContents {
				content, err := ioutil.ReadFile(file)
				defer os.Remove(file)
				if err != nil {
					t.Fatalf("Failed to read temp file: %v", err)
				}
				// ファイルが存在してもファイル名と内容が一致していなければエラーとなる
				if !tt.wantErr && string(content) != value {
					t.Errorf("Numbers wrote %q, expected %q", content, value)
				}
			}
		})
	}
}
