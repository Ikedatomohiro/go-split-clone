package input_test

import (
	"testing"

	in "split-clone/input"
)

func TestValidateInputt(t *testing.T) {
	var tests = []struct {
		name    string
		args    []string
		want    in.Exist
		wantErr bool
	}{
		{
			name:    "success: -lオプション",
			args:    []string{"split", "-l", "100", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: オプションなし",
			args:    []string{"split", "test.txt"},
			want:    in.Exist{Option: false, Prefix: false, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: -nオプション",
			args:    []string{"split", "-n", "3", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: -bオプション",
			args:    []string{"split", "-b", "3000", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: -bオプション m",
			args:    []string{"split", "-b", "100m", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: -lオプション + prefix",
			args:    []string{"split", "-l", "100", "test.txt", "sample_"},
			want:    in.Exist{Option: true, Prefix: true, FileName: true},
			wantErr: false,
		},
		{
			name:    "success: オプションなし、プレフィックスあり",
			args:    []string{"split", "test.txt", "sample_"},
			want:    in.Exist{Option: false, Prefix: true, FileName: true},
			wantErr: false,
		},
		{
			name:    "failure: オプションなし、プレフィックスの後にオプション引数がある",
			args:    []string{"split", "test.txt", "sample_", "-b", "1000"},
			want:    in.Exist{Option: false, Prefix: true, FileName: true},
			wantErr: true,
		},
		{
			name:    "failure: 引数なし",
			args:    []string{"split"},
			want:    in.Exist{Option: false, Prefix: false, FileName: false},
			wantErr: true,
		},
		{
			name:    "failure: -lオプション + prefixの後にも引数がある",
			args:    []string{"split", "-l", "100", "test.txt", "sample_", "dummy"},
			want:    in.Exist{Option: true, Prefix: true, FileName: true},
			wantErr: true,
		},
		{
			name:    "failure: オプション・数値の後にファイル名がない",
			args:    []string{"split", "-l", "100"},
			want:    in.Exist{Option: true, Prefix: false, FileName: false},
			wantErr: true,
		},
		{
			name:    "success: チェックは通過。ファイルを開く時にファイルなしのエラーとなる(bをファイル名、1000をprefixとみなす)",
			args:    []string{"split", "b", "1000", "test.txt"},
			want:    in.Exist{Option: false, Prefix: true, FileName: true},
			wantErr: false,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-bn", "100m", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: false},
			wantErr: true,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-l", "100m", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: false},
			wantErr: true,
		},
		{
			name:    "failure: オプション設定が多すぎる誤り",
			args:    []string{"split", "-l", "100m", "-n", "3", "test.txt"},
			want:    in.Exist{Option: true, Prefix: false, FileName: false},
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
