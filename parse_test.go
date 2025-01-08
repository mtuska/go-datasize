package datasize_test

import (
	"testing"

	"github.com/mtuska/go-datasize"
)

func TestParseBits(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    datasize.BitData
		wantErr bool
	}{
		{
			name:    "Valid Bits without space",
			args:    args{input: "1024b"},
			want:    datasize.BitData(1024),
			wantErr: false,
		},
		{
			name:    "Valid Bits with space",
			args:    args{input: "1024 b"},
			want:    datasize.BitData(1024),
			wantErr: false,
		},
		{
			name:    "Valid Bits with comma",
			args:    args{input: "1,024b"},
			want:    datasize.BitData(1024),
			wantErr: false,
		},
		{
			name:    "Invalid Bits",
			args:    args{input: "1024X"},
			want:    datasize.BitData(0),
			wantErr: true,
		},
		{
			name:    "Empty String",
			args:    args{input: ""},
			want:    datasize.BitData(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := datasize.ParseBits(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBytes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    datasize.ByteData
		wantErr bool
	}{
		{
			name:    "Valid Bytes without space",
			args:    args{input: "1024B"},
			want:    datasize.ByteData(1024),
			wantErr: false,
		},
		{
			name:    "Valid Bytes with space",
			args:    args{input: "1024 B"},
			want:    datasize.ByteData(1024),
			wantErr: false,
		},
		{
			name:    "Valid Bytes with comma",
			args:    args{input: "1,024B"},
			want:    datasize.ByteData(1024),
			wantErr: false,
		},
		{
			name:    "Invalid Bytes",
			args:    args{input: "1024X"},
			want:    datasize.ByteData(0),
			wantErr: true,
		},
		{
			name:    "Empty String",
			args:    args{input: ""},
			want:    datasize.ByteData(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := datasize.ParseBytes(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
