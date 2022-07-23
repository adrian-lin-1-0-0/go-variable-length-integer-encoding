# Variable-Length Integer Encoding 

Variable-Length Integer Encoding Implement

[rfc9000](https://datatracker.ietf.org/doc/html/rfc9000#section-16)

```
+======+========+=============+=======================+
| 2MSB | Length | Usable Bits | Range                 |
+======+========+=============+=======================+
| 00   | 1      | 6           | 0-63                  |
+------+--------+-------------+-----------------------+
| 01   | 2      | 14          | 0-16383               |
+------+--------+-------------+-----------------------+
| 10   | 4      | 30          | 0-1073741823          |
+------+--------+-------------+-----------------------+
| 11   | 8      | 62          | 0-4611686018427387903 |
+------+--------+-------------+-----------------------+
```

## Install

```shell
go get "github.com/adrian-lin-1-0-0/go-variable-length-integer-encoding"
```

## Usage

```go
package main

import (
	"fmt"

	encode "github.com/adrian-lin-1-0-0/go-variable-length-integer-encoding"
)

func main() {
	b, _ := encode.ToVarLenInt(0x3f)
	fmt.Println(encode.ToUint64(b))
}
```