package input_test

import (
	"testing"

	in "split-clone/input"
)

func TestCheckInput(t *testing.T) {
	var tests = []struct {
		name    string
		args    []string
		want    bool
		wantErr bool
	}{
		{
			name:    "success",
			args:    []string{"split", "-l", "100", "test.txt"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success: オプションなし",
			args:    []string{"split", "test.txt"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "success",
			args:    []string{"split", "-n", "100", "test.txt"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success",
			args:    []string{"split", "-b", "3000", "test.txt"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success",
			args:    []string{"split", "-b", "100m", "test.txt"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "success: ファイルを開く時にファイルなしのエラーとなる",
			args:    []string{"split", "b", "1000", "test.txt"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-bn", "100m", "test.txt"},
			want:    true,
			wantErr: true,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-l", "100m", "test.txt"},
			want:    true,
			wantErr: true,
		},
		{
			name:    "failure: オプション設定の誤り",
			args:    []string{"split", "-l", "100m", "-n", "3", "test.txt"},
			want:    true,
			wantErr: true,
		},
		{
			name:    "failure: 引数なし",
			args:    []string{"split"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := in.CheckInput(tt.args)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("CheckInput(%v) = %v", tt.args, err)
				}
			} else {
				if got != tt.want {
					t.Errorf("CheckInput(%v) = %v", tt.args, got)
				}
			}
		})
	}
}
