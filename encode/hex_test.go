package encode

import "testing"

func TestHexString(t *testing.T) {

	bs := []byte{0x1, 0xf, 0xe, 0xaf}

	s := HexEncodeToString(bs)

	t.Logf("s:%s\n", s)

	b, _ := HexDecodeString(s)

	t.Logf("b:%v\n", b)
}
