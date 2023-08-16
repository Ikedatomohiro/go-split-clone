package input_test

import (
	"testing"

	in "split-clone/input"
)

func TestValidateInputt(t *testing.T) {
	var tests = []struct {
		name    string
		args    []string
		want    in.ArgPosition
		wantErr bool
	}{
		{
			name:    "success: -lオプション",
			args:    []string{"split", "-l", "100", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 0},
			wantErr: false,
		},
		{
			name:    "success: オプションなし",
			args:    []string{"split", "test.txt"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 1, Prefix: 0},
			wantErr: false,
		},
		{
			name:    "success: -nオプション",
			args:    []string{"split", "-n", "3", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 0},
			wantErr: false,
		},
		{
			name:    "success: -bオプション",
			args:    []string{"split", "-b", "3000", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 0},
			wantErr: false,
		},
		{
			name:    "success: -bオプション m",
			args:    []string{"split", "-b", "100m", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 0},
			wantErr: false,
		},
		{
			name:    "success: -lオプション + prefix",
			args:    []string{"split", "-l", "100", "test.txt", "sample_"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 4},
			wantErr: false,
		},
		{
			name:    "success: -lオプション -aオプション + prefix",
			args:    []string{"split", "-l", "100", "-a", "4", "test.txt", "sample_"},
			want:    in.ArgPosition{Option: 1, AOption: 3, FileName: 5, Prefix: 6},
			wantErr: false,
		},
		{
			name:    "success: -lオプション -aオプション + prefix",
			args:    []string{"split", "-a", "4", "-l", "100", "test.txt", "sample_"},
			want:    in.ArgPosition{Option: 3, AOption: 1, FileName: 5, Prefix: 6},
			wantErr: false,
		},
		{
			name:    "failure: -lオプション -aオプションが0 + prefix",
			args:    []string{"split", "-a", "0", "-l", "100", "test.txt", "sample_"},
			want:    in.ArgPosition{Option: 0, AOption: 1, FileName: 0, Prefix: 0},
			wantErr: true,
		},
		{
			name:    "success: オプションなし、プレフィックスあり",
			args:    []string{"split", "test.txt", "sample_"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 1, Prefix: 2},
			wantErr: false,
		},
		{
			name:    "failure: オプションなし、プレフィックスの後にオプション引数がある",
			args:    []string{"split", "test.txt", "sample_", "-b", "1000"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 1, Prefix: 2},
			wantErr: true,
		},
		{
			name:    "failure: 引数なし",
			args:    []string{"split"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 0, Prefix: 0},
			wantErr: true,
		},
		{
			name:    "failure: -lオプション + prefixの後にも引数がある",
			args:    []string{"split", "-l", "100", "test.txt", "sample_", "dummy"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 3, Prefix: 4},
			wantErr: true,
		},
		{
			name:    "failure: オプション・数値の後にファイル名がない",
			args:    []string{"split", "-l", "100"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 0, Prefix: 0},
			wantErr: true,
		},
		{
			name:    "success: チェックは通過。ファイルを開く時にファイルなしのエラーとなる(bをファイル名、1000をprefixとみなす)",
			args:    []string{"split", "b", "1000", "test.txt"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 1, Prefix: 3},
			wantErr: false,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-bn", "100m", "test.txt"},
			want:    in.ArgPosition{Option: 0, AOption: 0, FileName: 0, Prefix: 0},
			wantErr: true,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-l", "100m", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 0, Prefix: 0},
			wantErr: true,
		},
		{
			name:    "failure: オプション設定が多すぎる誤り",
			args:    []string{"split", "-l", "100m", "-n", "3", "test.txt"},
			want:    in.ArgPosition{Option: 1, AOption: 0, FileName: 0, Prefix: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := in.ValidateInput(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
