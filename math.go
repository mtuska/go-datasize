package datasize

/////////////
// BitData //
/////////////

// Add adds two BitData values.
func (b BitData) Add(other BitData) BitData {
	return b + other
}

// AddBytes adds a ByteData value to the current BitData value.
func (b BitData) AddBytes(other ByteData) BitData {
	return b + other.ToBits()
}

// Sub subtracts another BitData value.
func (b BitData) Sub(other BitData) BitData {
	return b - other
}

// SubBytes subtracts a ByteData value from the current BitData value.
func (b BitData) SubBits(other ByteData) BitData {
	return b - other.ToBits()
}

// Mul multiplies BitData by a scalar.
func (b BitData) Mul(factor int64) BitData {
	return b * BitData(factor)
}

// MulFloat multiplies BitData by a floating-point factor.
func (b BitData) MulFloat(factor float64) BitData {
	return BitData(float64(b) * factor)
}

// Div divides BitData by a scalar.
func (b BitData) Div(divisor int64) BitData {
	return b / BitData(divisor)
}

// Clamp limits the BitData value within a specified range.
func (b BitData) Clamp(min, max BitData) BitData {
	if b < min {
		return min
	}
	if b > max {
		return max
	}
	return b
}

//////////////
// ByteData //
//////////////

// Add adds two ByteData values.
func (b ByteData) Add(other ByteData) ByteData {
	return b + other
}

// AddBits adds a BitData value to the current ByteData value.
func (b ByteData) AddBits(other BitData) ByteData {
	return b + other.ToBytes()
}

// Sub subtracts another ByteData value.
func (b ByteData) Sub(other ByteData) ByteData {
	return b - other
}

// SubBits subtracts a BitData value from the current ByteData value.
func (b ByteData) SubBits(other BitData) ByteData {
	return b - other.ToBytes()
}

// Mul multiplies ByteData by a scalar.
func (b ByteData) Mul(factor int64) ByteData {
	return b * ByteData(factor)
}

// MulFloat multiplies ByteData by a floating-point factor.
func (b ByteData) MulFloat(factor float64) ByteData {
	return ByteData(float64(b) * factor)
}

// Div divides ByteData by a scalar.
func (b ByteData) Div(divisor int64) ByteData {
	return b / ByteData(divisor)
}

// Clamp limits the ByteData value within a specified range.
func (b ByteData) Clamp(min, max ByteData) ByteData {
	if b < min {
		return min
	}
	if b > max {
		return max
	}
	return b
}
