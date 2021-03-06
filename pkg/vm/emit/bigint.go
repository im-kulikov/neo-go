package emit

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

// wordSizeBytes is a size of a big.Word (uint) in bytes.`
const wordSizeBytes = bits.UintSize / 8

// BytesToInt converts data in little-endian format to
// an integer.
func BytesToInt(data []byte) *big.Int {
	n := new(big.Int)
	size := len(data)
	if size == 0 {
		return big.NewInt(0)
	}

	isNeg := data[size-1]&0x80 != 0

	size = getEffectiveSize(data, isNeg)
	if size == 0 {
		if isNeg {
			return big.NewInt(-1)
		}

		return big.NewInt(0)
	}

	lw := size / wordSizeBytes
	ws := make([]big.Word, lw+1)
	for i := 0; i < lw; i++ {
		base := i * wordSizeBytes
		for j := base + 7; j >= base; j-- {
			ws[i] <<= 8
			ws[i] ^= big.Word(data[j])
		}
	}

	for i := size - 1; i >= lw*wordSizeBytes; i-- {
		ws[lw] <<= 8
		ws[lw] ^= big.Word(data[i])
	}

	if isNeg {
		for i := 0; i <= lw; i++ {
			ws[i] = ^ws[i]
		}

		shift := byte(wordSizeBytes-size%wordSizeBytes) * 8
		ws[lw] = ws[lw] & (^big.Word(0) >> shift)

		n.SetBits(ws)
		n.Neg(n)

		return n.Sub(n, big.NewInt(1))
	}

	return n.SetBits(ws)
}

// getEffectiveSize returns minimal number of bytes required
// to represent a number (two's complement for negatives).
func getEffectiveSize(buf []byte, isNeg bool) int {
	var b byte
	if isNeg {
		b = 0xFF
	}

	size := len(buf)
	for ; size > 0; size-- {
		if buf[size-1] != b {
			break
		}
	}

	return size
}

// IntToBytes converts integer to a slice in little-endian format.
func IntToBytes(n *big.Int) []byte {
	sign := n.Sign()
	if sign == 0 {
		return []byte{0}
	}

	var ws []big.Word
	if sign == 1 {
		ws = n.Bits()
	} else {
		n1 := new(big.Int).Add(n, big.NewInt(1))
		if n1.Sign() == 0 { // n == -1
			return []byte{0xFF}
		}

		ws = n1.Bits()
	}

	data := wordsToBytes(ws)

	size := len(data)
	for ; data[size-1] == 0; size-- {
	}

	data = data[:size]

	if data[size-1]&0x80 != 0 {
		data = append(data, 0)
	}

	if sign == -1 {
		for i := range data {
			data[i] = ^data[i]
		}
	}

	return data
}

func wordsToBytes(ws []big.Word) []byte {
	lb := len(ws) * wordSizeBytes
	bs := make([]byte, lb, lb+1)

	if wordSizeBytes == 8 {
		for i := range ws {
			binary.LittleEndian.PutUint64(bs[i*wordSizeBytes:], uint64(ws[i]))
		}
	} else {
		for i := range ws {
			binary.LittleEndian.PutUint32(bs[i*wordSizeBytes:], uint32(ws[i]))
		}
	}

	return bs
}
