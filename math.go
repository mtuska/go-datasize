package datasize

// Add
func (d DataSize) Add(other DataSize) DataSize {
	n := d
	n.bits += other.bits
	if n.bits >= byteSize {
		n.bytes += byteData(n.bits / 8)
		n.bits %= 8
	}
	n.bytes += other.bytes
	return n
}

// AddBits
func (d DataSize) AddBits(bits bitData) DataSize {
	n := d
	n.bits += bits
	if n.bits >= byteSize {
		n.bytes += byteData(n.bits / 8)
		n.bits %= byteSize
	}
	return n
}

// AddBytes
func (d DataSize) AddBytes(bytes byteData) DataSize {
	n := d
	n.bytes += bytes
	return n
}

// Sub
func (d DataSize) Sub(other DataSize) DataSize {
	n := d
	n.bytes -= other.bytes
	n.bits -= other.bits
	if n.bits < -byteSize {
		n.bytes -= byteData(n.bits / 8)
		n.bits %= byteSize
	}
	return n
}

// SubBits
func (d DataSize) SubBits(bits bitData) DataSize {
	n := d
	n.bits -= bits
	if n.bits < -byteSize {
		n.bytes -= byteData(n.bits / 8)
		n.bits %= byteSize
	}
	return n
}

// SubBytes
func (d DataSize) SubBytes(bytes byteData) DataSize {
	n := d
	n.bytes -= bytes
	return n
}

// Mul
func (d DataSize) Mul(factor int64) DataSize {
	n := d
	n.bytes *= byteData(factor)
	n.bits *= bitData(factor)
	if n.bits >= byteSize {
		n.bytes += byteData(n.bits / 8)
		n.bits %= byteSize
	}
	return n
}

// Div
func (d DataSize) Div(divisor int64) DataSize {
	n := d
	n.bytes /= byteData(divisor)
	n.bits /= bitData(divisor)
	if n.bits < 0 {
		n.bytes -= byteData(n.bits / 8)
		n.bits %= byteSize
	}
	return n
}

// Clamp
func (d DataSize) Clamp(min, max DataSize) DataSize {
	if d.bytes < min.bytes || (d.bytes == min.bytes && d.bits < min.bits) {
		return min
	}
	if d.bytes > max.bytes || (d.bytes == max.bytes && d.bits > max.bits) {
		return max
	}
	return d
}
