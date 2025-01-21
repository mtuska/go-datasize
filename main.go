package datasize

import (
	"fmt"
)

const byteSize = 8

type DataType string

const (
	Bits         DataType = "bits"          // Power of 2(1024)
	BinaryBytes           = "binary_bytes"  // Power of 2(1024)
	DecimalBytes          = "decimal_bytes" // Power of 10(1000)
)

type bitData int64

// Binary (1024-based) constants for DataSize.
const (
	Bit bitData = 1
)

// Decimal (1000-based) constants for DataSize.
const (
	KiloBit bitData = 1000 * Bit
	MegaBit         = 1000 * KiloBit
	GigaBit         = 1000 * MegaBit
	TeraBit         = 1000 * GigaBit
	PetaBit         = 1000 * TeraBit
	ExaBit          = 1000 * PetaBit
)

type byteData int64

// Binary (1024-based) constants for ByteData.
const (
	Byte byteData = 1 << (10 * iota)
	KibiByte
	MebiByte
	GibiByte
	TebiByte
	PebiByte
	ExbiByte
)

// Decimal (1000-based) constants for ByteData.
const (
	KiloByte byteData = 1000 * Byte
	MegaByte          = 1000 * KiloByte
	GigaByte          = 1000 * MegaByte
	TeraByte          = 1000 * GigaByte
	PetaByte          = 1000 * TeraByte
	ExaByte           = 1000 * PetaByte
)

type DataSize struct {
	bits     bitData
	bytes    byteData
	datatype DataType
}

func (d DataSize) SetDataType(datatype DataType) DataSize {
	d.datatype = datatype
	return d
}

// ConvertToFloat converts DataSize to the specified unit as float64.
func (d DataSize) ConvertToFloat(unit int64) float64 {
	return float64(d.bytes) / float64(unit)
}

// String method for DataSize.
// Returns the value in an exact unit, or bits if the value isn't a multiple of a larger unit.
func (d DataSize) String() string {
	switch d.datatype {
	case BinaryBytes:
		switch {
		case d.bytes >= ExbiByte && d.bytes%ExbiByte == 0:
			return fmt.Sprintf("%dEiB", d.bytes/ExbiByte)
		case d.bytes >= PebiByte && d.bytes%PebiByte == 0:
			return fmt.Sprintf("%dPiB", d.bytes/PebiByte)
		case d.bytes >= TebiByte && d.bytes%TebiByte == 0:
			return fmt.Sprintf("%dTiB", d.bytes/TebiByte)
		case d.bytes >= GibiByte && d.bytes%GibiByte == 0:
			return fmt.Sprintf("%dGiB", d.bytes/GibiByte)
		case d.bytes >= MebiByte && d.bytes%MebiByte == 0:
			return fmt.Sprintf("%dMiB", d.bytes/MebiByte)
		case d.bytes >= KibiByte && d.bytes%KibiByte == 0:
			return fmt.Sprintf("%dKiB", d.bytes/KibiByte)
		default:
			return fmt.Sprintf("%dB", d.bytes)
		}
	case DecimalBytes:
		switch {
		case d.bytes >= ExaByte && d.bytes%ExaByte == 0:
			return fmt.Sprintf("%dEB", d.bytes/ExaByte)
		case d.bytes >= PetaByte && d.bytes%PetaByte == 0:
			return fmt.Sprintf("%dPB", d.bytes/PetaByte)
		case d.bytes >= TeraByte && d.bytes%TeraByte == 0:
			return fmt.Sprintf("%dTB", d.bytes/TeraByte)
		case d.bytes >= GigaByte && d.bytes%GigaByte == 0:
			return fmt.Sprintf("%dGB", d.bytes/GigaByte)
		case d.bytes >= MegaByte && d.bytes%MegaByte == 0:
			return fmt.Sprintf("%dMB", d.bytes/MegaByte)
		case d.bytes >= KiloByte && d.bytes%KiloByte == 0:
			return fmt.Sprintf("%dKB", d.bytes/KiloByte)
		default:
			return fmt.Sprintf("%dB", d.bytes)
		}
	case Bits:
		switch {
		case d.bits >= ExaBit && d.bits%ExaBit == 0:
			return fmt.Sprintf("%dEb", d.bits/ExaBit)
		case d.bits >= PetaBit && d.bits%PetaBit == 0:
			return fmt.Sprintf("%dPb", d.bits/PetaBit)
		case d.bits >= TeraBit && d.bits%TeraBit == 0:
			return fmt.Sprintf("%dTb", d.bits/TeraBit)
		case d.bits >= GigaBit && d.bits%GigaBit == 0:
			return fmt.Sprintf("%dGb", d.bits/GigaBit)
		case d.bits >= MegaBit && d.bits%MegaBit == 0:
			return fmt.Sprintf("%dMb", d.bits/MegaBit)
		case d.bits >= KiloBit && d.bits%KiloBit == 0:
			return fmt.Sprintf("%dKb", d.bits/KiloBit)
		default:
			return fmt.Sprintf("%db", d.bits)
		}
	}
	return fmt.Sprintf("%db %dB", d.bits, d.bytes)
}

// HumanString method for DataSize.
func (d DataSize) HumanString() string {
	switch d.datatype {
	case BinaryBytes:
		switch {
		case d.bytes >= ExbiByte:
			return fmt.Sprintf("%.2f EiB", float64(d.bytes)/float64(ExbiByte))
		case d.bytes >= PebiByte:
			return fmt.Sprintf("%.2f PiB", float64(d.bytes)/float64(PebiByte))
		case d.bytes >= TebiByte:
			return fmt.Sprintf("%.2f TiB", float64(d.bytes)/float64(TebiByte))
		case d.bytes >= GibiByte:
			return fmt.Sprintf("%.2f GiB", float64(d.bytes)/float64(GibiByte))
		case d.bytes >= MebiByte:
			return fmt.Sprintf("%.2f MiB", float64(d.bytes)/float64(MebiByte))
		case d.bytes >= KibiByte:
			return fmt.Sprintf("%.2f KiB", float64(d.bytes)/float64(KibiByte))
		default:
			return fmt.Sprintf("%d B", d.bytes)
		}
	case DecimalBytes:
		switch {
		case d.bytes >= ExaByte:
			return fmt.Sprintf("%.2f EB", float64(d.bytes)/float64(ExaByte))
		case d.bytes >= PetaByte:
			return fmt.Sprintf("%.2f PB", float64(d.bytes)/float64(PetaByte))
		case d.bytes >= TeraByte:
			return fmt.Sprintf("%.2f TB", float64(d.bytes)/float64(TeraByte))
		case d.bytes >= GigaByte:
			return fmt.Sprintf("%.2f GB", float64(d.bytes)/float64(GigaByte))
		case d.bytes >= MegaByte:
			return fmt.Sprintf("%.2f MB", float64(d.bytes)/float64(MegaByte))
		case d.bytes >= KiloByte:
			return fmt.Sprintf("%.2f KB", float64(d.bytes)/float64(KiloByte))
		default:
			return fmt.Sprintf("%d B", d.bytes)
		}
	case Bits:
		_bits := float64(d.bits)
		_bits += float64(d.bytes) * byteSize
		switch {
		case _bits >= float64(ExaBit):
			return fmt.Sprintf("%.2f Eb", _bits/float64(ExaBit))
		case _bits >= float64(PetaBit):
			return fmt.Sprintf("%.2f Pb", _bits/float64(PetaBit))
		case _bits >= float64(TeraBit):
			return fmt.Sprintf("%.2f Tb", _bits/float64(TeraBit))
		case _bits >= float64(GigaBit):
			return fmt.Sprintf("%.2f Gb", _bits/float64(GigaBit))
		case _bits >= float64(MegaBit):
			return fmt.Sprintf("%.2f Mb", _bits/float64(MegaBit))
		case _bits >= float64(KiloBit):
			return fmt.Sprintf("%.2f Kb", _bits/float64(KiloBit))
		default:
			return fmt.Sprintf("%.2f b", _bits)
		}
	}
	return fmt.Sprintf("%db %dB", d.bits, d.bytes)
}
