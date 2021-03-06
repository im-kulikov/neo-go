package wrappers

import (
	"bytes"
	"sort"

	"github.com/CityOfZion/neo-go/pkg/core/state"
	"github.com/CityOfZion/neo-go/pkg/crypto/keys"
	"github.com/CityOfZion/neo-go/pkg/util"
)

// AccountState wrapper used for the representation of
// state.Account on the RPC Server.
type AccountState struct {
	Version    uint8             `json:"version"`
	ScriptHash util.Uint160      `json:"script_hash"`
	IsFrozen   bool              `json:"frozen"`
	Votes      []*keys.PublicKey `json:"votes"`
	Balances   []Balance         `json:"balances"`
}

// Balances type for sorting balances in rpc response.
type Balances []Balance

func (b Balances) Len() int           { return len(b) }
func (b Balances) Less(i, j int) bool { return bytes.Compare(b[i].Asset[:], b[j].Asset[:]) != -1 }
func (b Balances) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Balance response wrapper.
type Balance struct {
	Asset util.Uint256 `json:"asset"`
	Value util.Fixed8  `json:"value"`
}

// NewAccountState creates a new Account wrapper.
func NewAccountState(a *state.Account) AccountState {
	balances := make(Balances, 0, len(a.Balances))
	for k, v := range a.GetBalanceValues() {
		balances = append(balances, Balance{
			Asset: k,
			Value: v,
		})
	}

	sort.Sort(balances)

	// reverse scriptHash to be consistent with other client
	scriptHash, err := util.Uint160DecodeBytesBE(a.ScriptHash.BytesLE())
	if err != nil {
		scriptHash = a.ScriptHash
	}

	return AccountState{
		Version:    a.Version,
		ScriptHash: scriptHash,
		IsFrozen:   a.IsFrozen,
		Votes:      a.Votes,
		Balances:   balances,
	}
}
