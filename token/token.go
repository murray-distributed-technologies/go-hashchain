package token

import (
	"bytes"

	"github.com/libsv/go-bk/crypto"
)

type Hashchain struct {
	N               int      // number of events requiring a token to access
	T0              [32]byte // initial token
	Tn              []byte   // Hash^n of t_0
	CurrentToken    []byte   // current token
	CurrentTokenNum int      // number of last token. Token 1 = t_n-1
	LastToken       []byte
}

func GenerateTokenChain(hashchain *Hashchain) (*Hashchain, error) {
	var err error
	if hashchain, err = GenerateTokenN(hashchain); err != nil {
		return nil, err
	}
	hashchain.CurrentToken = hashchain.Tn
	hashchain.CurrentTokenNum = 0
	return hashchain, nil
}

func GenerateTokenN(hashchain *Hashchain) (*Hashchain, error) {
	t0 := hashchain.T0[:]
	tN := t0
	for i := 0; i < hashchain.N; i++ {
		tN = crypto.Sha256(tN)
	}
	hashchain.Tn = tN
	return hashchain, nil
}

func GetNextToken(hashchain *Hashchain) (*Hashchain, error) {
	t0 := hashchain.T0[:]
	token := t0
	for i := 0; i < (hashchain.N - hashchain.CurrentTokenNum); i++ {
		token = crypto.Sha256(token)
	}
	hashchain.CurrentTokenNum++
	hashchain.CurrentToken = token
	return hashchain, nil
}

func CheckToken(hashchain *Hashchain) (bool, error) {
	token := crypto.Sha256(hashchain.LastToken)
	if bytes.Compare(hashchain.CurrentToken, token) == 0 {
		return true, nil
	}
	return false, nil
}
