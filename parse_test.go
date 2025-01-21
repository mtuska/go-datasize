package datasize

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    DataSize
		wantErr bool
	}{
		{
			name:    "Valid Bytes without space",
			args:    args{input: "1024B"},
			want:    DataSize{bytes: 1024, datatype: DecimalBytes},
			wantErr: false,
		},
		{
			name:    "Valid Bytes with space",
			args:    args{input: "1024 B"},
			want:    DataSize{bytes: 1024, datatype: DecimalBytes},
			wantErr: false,
		},
		{
			name:    "Valid Bytes with comma",
			args:    args{input: "1,024B"},
			want:    DataSize{bytes: 1024, datatype: DecimalBytes},
			wantErr: false,
		},
		{
			name:    "Scientific Notation",
			args:    args{input: "9.7248104E-4 MB"},
			want:    DataSize{bytes: 972, datatype: DecimalBytes},
			wantErr: false,
		},
		{
			name:    "Scientific Notation Bits",
			args:    args{input: "9.7248104E-4 Mb"},
			want:    DataSize{bits: 3, bytes: 126, datatype: Bits},
			wantErr: false,
		},
		{
			name:    "Invalid Unit",
			args:    args{input: "1024X"},
			want:    DataSize{},
			wantErr: true,
		},
		{
			name:    "Empty String",
			args:    args{input: ""},
			want:    DataSize{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDs, err := Parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDs, tt.want) {
				t.Errorf("Parse() = %v, want %v", gotDs, tt.want)
			}
		})
	}
}
