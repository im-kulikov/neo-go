package transaction

import (
	"github.com/CityOfZion/neo-go/pkg/io"
	"github.com/CityOfZion/neo-go/pkg/util"
)

// Input represents a Transaction input (CoinReference).
type Input struct {
	// The hash of the previous transaction.
	PrevHash util.Uint256 `json:"txid"`

	// The index of the previous transaction.
	PrevIndex uint16 `json:"vout"`
}

// DecodeBinary implements Serializable interface.
func (in *Input) DecodeBinary(br *io.BinReader) {
	br.ReadBytes(in.PrevHash[:])
	in.PrevIndex = br.ReadU16LE()
}

// EncodeBinary implements Serializable interface.
func (in *Input) EncodeBinary(bw *io.BinWriter) {
	bw.WriteBytes(in.PrevHash[:])
	bw.WriteU16LE(in.PrevIndex)
}
