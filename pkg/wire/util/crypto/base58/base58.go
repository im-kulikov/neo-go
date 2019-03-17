package base58

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/CityOfZion/neo-go/pkg/wire/util/crypto/hash"
)

const prefix rune = '1'

var decodeMap = map[rune]int64{
	'1': 0, '2': 1, '3': 2, '4': 3, '5': 4,
	'6': 5, '7': 6, '8': 7, '9': 8, 'A': 9,
	'B': 10, 'C': 11, 'D': 12, 'E': 13, 'F': 14,
	'G': 15, 'H': 16, 'J': 17, 'K': 18, 'L': 19,
	'M': 20, 'N': 21, 'P': 22, 'Q': 23, 'R': 24,
	'S': 25, 'T': 26, 'U': 27, 'V': 28, 'W': 29,
	'X': 30, 'Y': 31, 'Z': 32, 'a': 33, 'b': 34,
	'c': 35, 'd': 36, 'e': 37, 'f': 38, 'g': 39,
	'h': 40, 'i': 41, 'j': 42, 'k': 43, 'm': 44,
	'n': 45, 'o': 46, 'p': 47, 'q': 48, 'r': 49,
	's': 50, 't': 51, 'u': 52, 'v': 53, 'w': 54,
	'x': 55, 'y': 56, 'z': 57,
}

// Decode decodes the base58 encoded string.
func Decode(s string) ([]byte, error) {
	var (
		startIndex = 0
		zero       = 0
	)
	for i, c := range s {
		if c == prefix {
			zero++
		} else {
			startIndex = i
			break
		}
	}

	var (
		n   = big.NewInt(0)
		div = big.NewInt(58)
	)
	for _, c := range s[startIndex:] {
		charIndex, ok := decodeMap[c]
		if !ok {
			return nil, fmt.Errorf(
				"invalid character '%c' when decoding this base58 string: '%s'", c, s,
			)
		}
		n.Add(n.Mul(n, div), big.NewInt(charIndex))
	}

	out := n.Bytes()
	buf := make([]byte, (zero + len(out)))
	copy(buf[zero:], out[:])

	return buf, nil
}

// Encode encodes a byte slice to be a base58 encoded string.
func Encode(bytes []byte) string {
	var (
		lookupTable = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		x           = new(big.Int).SetBytes(bytes)
		r           = new(big.Int)
		m           = big.NewInt(58)
		zero        = big.NewInt(0)
		encoded     string
	)

	for x.Cmp(zero) > 0 {
		x.QuoRem(x, m, r)
		encoded = string(lookupTable[r.Int64()]) + encoded
	}

	return encoded
}

// CheckDecode decodes the given string.
func CheckDecode(s string) (b []byte, err error) {
	b, err = Decode(s)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(s); i++ {
		if s[i] != '1' {
			break
		}
		b = append([]byte{0x00}, b...)
	}

	if len(b) < 5 {
		return nil, fmt.Errorf("Invalid base-58 check string: missing checksum. -1")
	}

	hash, err := hash.DoubleSha256(b[:len(b)-4])

	if err != nil {
		return nil, fmt.Errorf("Could not double sha256 data")
	}

	if bytes.Compare(hash[0:4], b[len(b)-4:]) != 0 {
		return nil, fmt.Errorf("Invalid base-58 check string: invalid checksum. -2")
	}

	// Strip the 4 byte long hash.
	b = b[:len(b)-4]

	return b, nil
}

// CheckEncode encodes b into a base-58 check encoded string.
func CheckEncode(b []byte) (string, error) {
	hash, err := hash.DoubleSha256(b)
	if err != nil {
		return "", fmt.Errorf("Could not double sha256 data")
	}
	b = append(b, hash[0:4]...)

	return Encode(b), nil
}
