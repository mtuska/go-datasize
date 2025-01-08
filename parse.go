package datasize

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parse(input string) (value float64, unit string, err error) {
	re := regexp.MustCompile(`^([\d,]+(\.\d+)?)\s?([a-zA-Z]+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		err = errors.New("invalid format, expected '<value><unit>' or '<value> <unit>'")
		return
	}

	value, err = strconv.ParseFloat(strings.ReplaceAll(matches[1], ",", ""), 64)
	if err != nil {
		return
	}
	unit = matches[3]
	return
}

/////////////
// BitData //
/////////////

// ParseBits parses a human-readable data-size/data-rate string (e.g., "10Mb"; "10Mbps") into BitData.
func ParseBits(input string) (BitData, error) {
	units := map[string]BitData{
		"b":  Bit,
		"Kb": KiloBit,
		"Mb": MegaBit,
		"Gb": GigaBit,
		"Tb": TeraBit,
		"Pb": PetaBit,
		"Eb": ExaBit,

		"Kib": KibiBit,
		"Mib": MebiBit,
		"Gib": GibiBit,
		"Tib": TebiBit,
		"Pib": PebiBit,
		"Eib": ExbiBit,
	}

	value, sunit, err := parse(input)
	if err != nil {
		return 0, err
	}

	unit, ok := units[sunit]
	if !ok {
		return 0, errors.New("unknown unit")
	}

	return BitData(value * float64(unit)), nil
}

//////////////
// ByteData //
//////////////

// ParseBytes parses a human-readable data-size/data-rate string (e.g., "10MB"; "10MBps") into ByteData.
func ParseBytes(input string) (ByteData, error) {
	units := map[string]ByteData{
		"B":  Byte,
		"KB": KiloByte,
		"MB": MegaByte,
		"GB": GigaByte,
		"TB": TeraByte,
		"PB": PetaByte,
		"EB": ExaByte,

		"KiB": KibiByte,
		"MiB": MebiByte,
		"GiB": GibiByte,
		"TiB": TebiByte,
		"PiB": PebiByte,
		"EiB": ExbiByte,
	}

	value, sunit, err := parse(input)
	if err != nil {
		return 0, err
	}

	unit, ok := units[sunit]
	if !ok {
		return 0, errors.New("unknown unit")
	}

	return ByteData(value * float64(unit)), nil
}
