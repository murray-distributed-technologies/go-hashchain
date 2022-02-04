package main

import (
	"encoding/hex"
	"fmt"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/murray-distributed-technologies/go-hashchain/token"
	"github.com/murray-distributed-technologies/go-hashchain/transaction"
	"github.com/murray-distributed-technologies/go-hashchain/woc"
)

func main() {
	destPrivKey, _ := wif.DecodeWIF("")
	privKey, _ := wif.DecodeWIF("")
	changeAddress, _ := bscript.NewAddressFromPublicKey(privKey.PrivKey.PubKey(), true)
	address, _ := bscript.NewAddressFromPublicKey(destPrivKey.PrivKey.PubKey(), true)

	var t0 [32]byte
	// currently using 32 byte hex string for t0
	res, _ := hex.DecodeString("4bc1e4b940e46fe42bfe00f2be03d5a41c0f312478d2a31e32d1864c75ce0e2d")
	copy(t0[:], res[:])
	var n int
	n = 5
	tokenNum := 0
	hashchain := &token.Hashchain{
		N:  n,
		T0: t0,
	}

	hashchain, _ = token.GenerateTokenChain(hashchain)
	for i := 0; i < tokenNum; i++ {
		hashchain, _ = token.GetNextToken(hashchain)
	}

	var sats uint64
	var vOut uint32

	txId := ""
	vOut = 0
	amount := uint64(5000)

	o, _ := woc.GetTransactionOutput(txId, int(vOut))

	sats = uint64(o.Value * 100000000)
	scriptPubKey, err := bscript.NewFromHexString(o.ScriptPubKey.Hex)
	if err != nil {
		fmt.Println(err)
	}

	txIdBytes, _ := hex.DecodeString(txId)

	utxo := &bt.UTXO{
		TxID:          txIdBytes,
		Vout:          vOut,
		LockingScript: scriptPubKey,
		Satoshis:      sats,
	}

	rawTxString, err := transaction.CreateHashTransaction(utxo, privKey.PrivKey, hashchain.CurrentToken, hashchain.CurrentToken, address.AddressString, changeAddress.AddressString, amount)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rawTxString)

}
