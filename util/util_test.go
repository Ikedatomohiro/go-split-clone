package util_test

import (
	"testing"

	in "split-clone/input"
	ut "split-clone/util"
)

type GetParamArgs struct {
	args []string
	e    in.Exist
}

func TestGetFilename(t *testing.T) {
	var tests = []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			args:    "xaa",
			want:    "xab",
			wantErr: false,
		},
		{
			name:    "success: zの処理",
			args:    "xaz",
			want:    "xba",
			wantErr: false,
		},
		{
			name:    "success: zの処理",
			args:    "zzz",
			want:    "aaaa",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ut.GetFilename(tt.args)
			if got != tt.want {
				t.Errorf("GetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetParam(t *testing.T) {
	var tests = []struct {
		name    string
		args    GetParamArgs
		want    in.Input
		wantErr bool
	}{
		{
			name: "success: オプションなし、プレフィックスなし（-l 1000のオプションをデフォルトとする）",
			args: GetParamArgs{
				args: []string{"split", "test.txt"},
				e:    in.Exist{Option: false, Prefix: false, FileName: true},
			},
			want:    in.Input{Option: "l", OptionValue: 1000, FileName: "test.txt", Prefix: ""},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "3000", "test.txt"},
				e:    in.Exist{Option: true, Prefix: false, FileName: true},
			},
			want:    in.Input{Option: "b", OptionValue: 3000, FileName: "test.txt", Prefix: ""},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "300m", "test.txt"},
				e:    in.Exist{Option: true, Prefix: false, FileName: true},
			},
			want:    in.Input{Option: "b", OptionValue: 300000000, FileName: "test.txt", Prefix: ""},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-n", "3", "test.txt"},
				e:    in.Exist{Option: true, Prefix: false, FileName: true},
			},
			want:    in.Input{Option: "n", OptionValue: 3, FileName: "test.txt", Prefix: ""},
			wantErr: false,
		},
		{
			name: "success: オプションなし、プレフィックスあり",
			args: GetParamArgs{
				args: []string{"split", "test.txt", "sample_"},
				e:    in.Exist{Option: false, Prefix: true, FileName: true},
			},
			want:    in.Input{Option: "l", OptionValue: 1000, FileName: "test.txt", Prefix: "sample_"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスあり",
			args: GetParamArgs{
				args: []string{"split", "-n", "3", "test.txt", "sample_"},
				e:    in.Exist{Option: true, Prefix: true, FileName: true},
			},
			want:    in.Input{Option: "n", OptionValue: 3, FileName: "test.txt", Prefix: "sample_"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ut.GetParam(tt.args.args, tt.args.e)
			if got != tt.want {
				t.Errorf("GetParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}
