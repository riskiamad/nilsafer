# NILSAFER

Nilsafer is a utility to help you get value from the pointer address without worrying with panic. Nilsafer will return zero value or default value from its type if the pointer address is nil.

## Installation

```bash
go get -u github.com/riskiamad/nilsafer
```

## Getting Started

### Simple Usage ValueOrZero

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

### Simple Usage ValueOrDefault

```go
package main

import (
	"fmt"
	"github.com/riskiamad/nilsafer"
)

func main() {
	var s = "hello"
	var nilString *string = nil
	var d = "world"

	fmt.Println(nilsafer.ValueOrDefault(&s, d))
	fmt.Println(nilsafer.ValueOrDefault(nilString, d))
}

/*output:
hello
world
 */

```