package encode

import (
	"testing"
)

func printErr(i, v uint64, t *testing.T) {
	t.Errorf("%v should be %v", i, v)
}

func testToVarLenInt(v uint64, t *testing.T) {
	b, _ := ToVarLenInt(v)
	if i, _ := ToUint64(b); i != v {
		printErr(i, v, t)
	}
}

func TestToVarLenInt(t *testing.T) {
	tv := func(v uint64) {
		testToVarLenInt(v, t)
	}

	tv(0x00)
	tv(0x3e)
	tv(0x3f)
	tv(0x3fff)
	tv(0x3fffffff)
}
