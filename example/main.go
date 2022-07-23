package main

import (
	"encoding/hex"
	"fmt"

	encode "github.com/adrian-lin-1-0-0/go-variable-length-integer-encoding"
)

func main() {
	b, _ := encode.ToVarLenInt(0x3f)
	fmt.Println(hex.EncodeToString(b))
	fmt.Println(encode.ToUint64(b))
}
