package vm_test

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	src := `
		package foo
		 import (
			"github.com/CityOfZion/neo-go/pkg/interop/crypto"
		 )
		func Main() []byte {
			src := []byte{0x97}
			hash := crypto.SHA256(src)
			return hash
		}
	`
	eval(t, src, []byte{0x2a, 0xa, 0xb7, 0x32, 0xb4, 0xe9, 0xd8, 0x5e, 0xf7, 0xdc, 0x25, 0x30, 0x3b, 0x64, 0xab, 0x52, 0x7c, 0x25, 0xa4, 0xd7, 0x78, 0x15, 0xeb, 0xb5, 0x79, 0xf3, 0x96, 0xec, 0x6c, 0xac, 0xca, 0xd3})
}

func TestSHA1(t *testing.T) {
	src := `
		package foo
		 import (
			"github.com/CityOfZion/neo-go/pkg/interop/crypto"
		 )
		func Main() []byte {
			src := []byte{0x97}
			hash := crypto.SHA1(src)
			return hash
		}
	`
	eval(t, src, []byte{0xfa, 0x13, 0x8a, 0xe3, 0x56, 0xd3, 0x5c, 0x8d, 0x77, 0x8, 0x3c, 0x40, 0x6a, 0x5b, 0xe7, 0x37, 0x45, 0x64, 0x3a, 0xae})
}

func TestHash160(t *testing.T) {
	src := `
		package foo
		 import (
			"github.com/CityOfZion/neo-go/pkg/interop/crypto"
		 )
		func Main() []byte {
			src := []byte{0x97}
			hash := crypto.Hash160(src)
			return hash
		}
	`
	eval(t, src, []byte{0x5f, 0xa4, 0x1c, 0x76, 0xf7, 0xe8, 0xca, 0x72, 0xb7, 0x18, 0xff, 0x59, 0x22, 0x91, 0xc2, 0x3a, 0x3a, 0xf5, 0x58, 0x6c})
}

func TestHash256(t *testing.T) {
	src := `
		package foo
		 import (
			"github.com/CityOfZion/neo-go/pkg/interop/crypto"
		 )
		func Main() []byte {
			src := []byte{0x97}
			hash := crypto.Hash256(src)
			return hash
		}
	`
	eval(t, src, []byte{0xc0, 0x85, 0x26, 0xad, 0x17, 0x36, 0x53, 0xee, 0xb8, 0xc7, 0xf4, 0xae, 0x82, 0x8b, 0x6e, 0xa1, 0x84, 0xac, 0x5a, 0x3, 0x8a, 0xf6, 0xc3, 0x68, 0x23, 0xfa, 0x5f, 0x5d, 0xd9, 0x1b, 0x91, 0xa2})
}
