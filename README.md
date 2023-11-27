# NILSAFER

Nilsafer is a utility to help you get value from the pointer address without worrying with panic. Nilsafer will return zero value from its type if the pointer address is nil.

## Installation

```bash
go get -u github.com/riskiamad/nilsafer
```

## Getting Started

### Simple Usage Example with String

```go
package main

import (
	"fmt"
	"github.com/riskiamad/nilsafer"
)

func main() {
	var s = "hello"
	var nilString *string = nil

	fmt.Println(nilsafer.ValueOrZero(&s))
	fmt.Println(nilsafer.ValueOrZero(nilString))
}

/*output:
hello

 */

```