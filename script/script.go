package script

import (
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/libsv/go-bt/v2/sighash"
	pushtx "github.com/murray-distributed-technologies/go-pushtx/script"
)

func NewLockingScript(hash []byte, address string) (*bscript.Script, error) {
	var err error
	s := &bscript.Script{}
	s.AppendOpcodes(bscript.OpSHA256)
	if err = s.AppendPushData(hash); err != nil {
		return nil, err
	}

	s.AppendOpcodes(bscript.OpEQUALVERIFY)
	if s, err = pushtx.AppendP2PKH(s, address); err != nil {
		return nil, err
	}
	return s, nil
}

func NewUnlockingScript(pubKey, sig, token []byte, sigHashFlag sighash.Flag) (*bscript.Script, error) {
	sigBuf := []byte{}
	sigBuf = append(sigBuf, sig...)
	sigBuf = append(sigBuf, uint8(sigHashFlag))

	scriptBuf := [][]byte{sigBuf, pubKey, token}

	s := &bscript.Script{}
	err := s.AppendPushDataArray(scriptBuf)
	if err != nil {
		return nil, err
	}
	return s, nil
}
