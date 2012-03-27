package bindata

import (
	"bytes"
	"compress/gzip"
	"io"
)

// MinimizePng returns the decompressed binary data.
// It panics if an error occurred.
func MinimizePng() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0xea,0x0c,
0xf0,0x73,0xe7,0xe5,0x92,0xe2,0x62,0x60,0x60,0xe0,0xf5,0xf4,
0x70,0x09,0x02,0xd2,0x0d,0x20,0xcc,0xc1,0x06,0x24,0x0f,0xdb,
0x25,0x9e,0x06,0x52,0x8c,0xc5,0x41,0xee,0x4e,0x0c,0xeb,0xce,
0xc9,0xbc,0x04,0x72,0xd8,0x92,0xbc,0xdd,0x5d,0x18,0xfe,0x83,
0xe0,0x82,0xbd,0xcb,0x27,0x03,0x45,0x38,0x0b,0x3c,0x22,0x8b,
0x19,0x18,0xb8,0x85,0x41,0x98,0x91,0x61,0xd6,0x1c,0x09,0xa0,
0x20,0x7b,0x89,0xa7,0xaf,0x2b,0xfb,0x1d,0x66,0x69,0x26,0x7e,
0xfe,0x5a,0x63,0xce,0x2e,0xa0,0x90,0x6c,0x66,0x48,0x44,0x89,
0x73,0x7e,0x6e,0x6e,0x6a,0x5e,0x09,0x03,0x08,0x38,0x17,0xa5,
0x26,0x96,0xa4,0xa6,0x28,0x94,0x67,0x96,0x64,0x28,0xb8,0x7b,
0xfa,0x06,0xa4,0xe8,0xa5,0xb2,0x03,0xc5,0xcb,0x3c,0x5d,0x1c,
0x43,0x2a,0x6e,0xbd,0xbd,0xdc,0xc8,0xc9,0x60,0xc0,0xc1,0x62,
0xf8,0xfa,0x9f,0xf5,0x6b,0x3e,0xd6,0xa5,0x17,0x52,0x9a,0xdc,
0xcf,0x71,0x30,0xe0,0x01,0x6e,0x73,0xdc,0xcb,0xa7,0xce,0x7e,
0xef,0x61,0xff,0xc1,0x9c,0x91,0x83,0x81,0x51,0x81,0x81,0xa5,
0x81,0x41,0x80,0x81,0xc9,0x81,0x01,0xc2,0xc1,0xa7,0xf5,0x40,
0x2b,0x9b,0xa1,0x00,0xa3,0xdc,0xfb,0x37,0x5f,0x66,0x81,0xb8,
0x9e,0xae,0x7e,0x2e,0xeb,0x9c,0x12,0x9a,0x00,0x01,0x00,0x00,
0xff,0xff,0x91,0x03,0x30,0x10,0x1f,0x01,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}