package encrpt

import (
	"crypto/sha256"
)

var h = sha256.New()

func simpleSha256(data []byte) []byte {
	return h.Sum(data)
}

func xorBytes(a, b []byte) []byte {
	n := min(len(a), len(b))
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sliceByteXOR(data []byte, b byte) []byte {
	ret := make([]byte, len(data))
	for i, val := range data {
		ret[i] = val ^ b
	}
	return ret
}

func sliceTo(data []byte, s int) []byte {
	if len(data) == s {
		return data
	}

	ret := make([]byte, s)
	m := max(len(data), s)
	for i := 0; i < m; i++ {
		ret[i%s] = ret[i%s] ^ data[i%len(data)]
	}
	return ret
}
