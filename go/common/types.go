// Should only use structs and methods that are prefixed with the word
// Safe concurrently, anything else is not safe to use concurrently
package common

import (
	"math/big"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

type PairData struct {
	TokenAddress string `json:"tokenAddress"`
	Fee          int    `json:"fee"`
}

type PairsData struct {
	PData []PairData `json:"data"`
}

type SafeNonce struct {
	NonceMu sync.Mutex
	Nonce   *big.Int
}

func (nonce *SafeNonce) SafeIncrementNonce() {
	nonce.NonceMu.Lock()
	nonce.Nonce.Add(nonce.Nonce, big.NewInt(1))
	nonce.NonceMu.Unlock()
}

type Reserves struct {
	Token0 *big.Int
	Token1 *big.Int
}
