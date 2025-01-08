# go-datasize

`go-datasize` is a lightweight package for working with data sizes in Go. It provides utilities for conversion, formatting, and arithmetic with sizes in bits and bytes.

## Installation

```bash
go get github.com/mtuska/go-datasize
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mtuska/go-datasize"
)

func main() {
    var size datasize.ByteData = 1024 * 1024 // 1 MB
    fmt.Println(size.String()) // "1.00 MB"
}
```