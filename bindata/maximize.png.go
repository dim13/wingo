package bindata

import (
	"bytes"
	"compress/gzip"
	"io"
)

// MaximizePng returns the decompressed binary data.
// It panics if an error occurred.
func MaximizePng() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x09,0x6e,0x88,0x00,0xff,0xea,0x0c,
0xf0,0x73,0xe7,0xe5,0x92,0xe2,0x62,0x60,0x60,0xe0,0xf5,0xf4,
0x70,0x09,0x02,0xd2,0x0d,0x20,0xcc,0xc1,0x06,0x24,0x0f,0xdb,
0x25,0x9e,0x06,0x52,0x8c,0xc5,0x41,0xee,0x4e,0x0c,0xeb,0xce,
0xc9,0xbc,0x04,0x72,0xd8,0x92,0xbc,0xdd,0x5d,0x18,0xfe,0x83,
0xe0,0x82,0xbd,0xcb,0x27,0x03,0x45,0x38,0x0b,0x3c,0x22,0x8b,
0x19,0x18,0xb8,0x85,0x41,0x98,0x91,0x61,0xd6,0x1c,0x09,0xa0,
0x20,0x7b,0x89,0xa7,0xaf,0x2b,0xfb,0x1d,0x66,0x69,0x66,0x65,
0x85,0x77,0x82,0xa6,0x5e,0x40,0x21,0xd9,0xcc,0x90,0x88,0x12,
0xe7,0xfc,0xdc,0xdc,0xd4,0xbc,0x12,0x06,0x10,0x70,0x2e,0x4a,
0x4d,0x2c,0x49,0x4d,0x51,0x28,0xcf,0x2c,0xc9,0x50,0x70,0xf7,
0xf4,0x0d,0x48,0xd1,0x4b,0x65,0x07,0x8a,0x5f,0xf0,0x74,0x71,
0x0c,0xa9,0xb8,0xf5,0xf6,0xee,0x41,0x51,0x06,0x41,0x1e,0x87,
0x83,0x17,0xea,0xcf,0x5f,0x6c,0xd0,0xf6,0x4f,0x59,0xd6,0xdc,
0xbb,0xff,0x4e,0x07,0x03,0x12,0x58,0x25,0xe5,0xbb,0xcf,0xf8,
0x6c,0x46,0xfe,0x2f,0xf6,0x2a,0xe1,0x06,0x06,0x01,0x06,0x26,
0x07,0x06,0x0e,0x06,0x46,0x05,0x06,0x16,0x62,0x38,0x01,0x87,
0x7a,0x4f,0xf3,0x5f,0x9a,0x71,0x3f,0x56,0xdd,0xba,0x41,0x19,
0x9b,0xc2,0x1f,0x53,0xbf,0x33,0xcd,0x3a,0x7f,0x82,0xef,0x01,
0x07,0x29,0xc6,0x82,0x39,0x48,0xa0,0x61,0xb1,0x98,0xb6,0x00,
0xf3,0x82,0x92,0x4f,0x2a,0xc6,0x20,0xae,0xa7,0xab,0x9f,0xcb,
0x3a,0xa7,0x84,0x26,0x40,0x00,0x00,0x00,0xff,0xff,0x8d,0x79,
0x01,0x96,0x79,0x01,0x00,0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}