package datasize_test

import (
	"testing"

	"github.com/mtuska/go-datasize"
)

func TestBitData_ConvertTo(t *testing.T) {
	type args struct {
		unit datasize.BitData
	}
	tests := []struct {
		name string
		b    datasize.BitData
		args args
		want datasize.BitData
		err  bool
	}{
		{"Bits to KibiBits", 1024, args{datasize.KibiBit}, 1, false},
		{"Bits to MegaBits", 1048576, args{datasize.MegaBit}, 1, false},
		{"Bits to GigaBits", 1073741824, args{datasize.GigaBit}, 1, false},

		{"Bits to KiloBits", 1000, args{datasize.KiloBit}, 1, false},
		{"Bits to MegaBits", 1000000, args{datasize.MegaBit}, 1, false},
		{"Bits to GigaBits", 1000000000, args{datasize.GigaBit}, 1, false},

		{"Panic when Bits to KiloBits", 512, args{datasize.KiloBit}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && tt.err == true {
					t.Errorf("Expected panic but did not get one")
				}
			}()
			if got := tt.b.ConvertTo(tt.args.unit); got != tt.want {
				t.Errorf("BitData.ConvertTo() = %v, want %v", got, tt.want)
			}
			if tt.err == true {
				t.Errorf("Expected panic but did not get one")
			}
		})
	}
}

func TestBitData_String(t *testing.T) {
	tests := []struct {
		name string
		b    datasize.BitData
		want string
	}{
		{"Gigabit 1.5", datasize.GigaBit.MulFloat(1.5), "1.50 Gb"},
		{"Gigabit 2", datasize.GigaBit * 2, "2.00 Gb"},
		{"ExaBit", datasize.ExaBit, "1.00 Eb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("BitData.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteData_ConvertTo(t *testing.T) {
	type args struct {
		unit datasize.ByteData
	}
	tests := []struct {
		name string
		b    datasize.ByteData
		args args
		want datasize.ByteData
		err  bool
	}{
		{"Bytes to KibiBytes", 1024, args{datasize.KibiByte}, 1, false},
		{"Bytes to MebiBytes", 1048576, args{datasize.MebiByte}, 1, false},
		{"Bytes to GibiBytes", 1073741824, args{datasize.GibiByte}, 1, false},

		{"Bytes to KiloBytes", 1000, args{datasize.KiloByte}, 1, false},
		{"Bytes to MegaBytes", 1000000, args{datasize.MegaByte}, 1, false},
		{"Bytes to GigaBytes", 1000000000, args{datasize.GigaByte}, 1, false},

		{"Panic when Bytes to KiloBytes", 512, args{datasize.KiloByte}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && tt.err == true {
					t.Errorf("Expected panic but did not get one")
				}
			}()
			if got := tt.b.ConvertTo(tt.args.unit); got != tt.want {
				t.Errorf("ByteData.ConvertTo() = %v, want %v", got, tt.want)
			}
			if tt.err == true {
				t.Errorf("Expected panic but did not get one")
			}
		})
	}
}

func TestByteData_String(t *testing.T) {
	tests := []struct {
		name string
		b    datasize.ByteData
		want string
	}{
		{"Gigabit 1.5", datasize.GigaByte.MulFloat(1.5), "1.50 GB"},
		{"Gigabit 2", datasize.GigaByte * 2, "2.00 GB"},
		{"ExaBit", datasize.ExaByte, "1.00 EB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("ByteData.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
