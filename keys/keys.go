package keys

import (
	"math/big"

	"github.com/libsv/go-bk/bec"
)

func NewPrivateKey(token []byte, privKey *bec.PrivateKey) (*bec.PrivateKey, error) {
	newPrivKey := &bec.PrivateKey{}
	tokenBigInt := &big.Int{}
	tokenBigInt = tokenBigInt.SetBytes(token)
	newPrivKey.D = privKey.D.Add(privKey.D, tokenBigInt)
	newPubKey, _ := NewPublicKey(token, privKey.PubKey())
	newPrivKey.PublicKey = *newPubKey.ToECDSA()
	return newPrivKey, nil
}

func NewPublicKey(token []byte, pubKey *bec.PublicKey) (*bec.PublicKey, error) {
	curve := bec.S256()

	newX, newY := curve.ScalarBaseMult(token)
	newX, newY = curve.Add(newX, newY, pubKey.X, pubKey.Y)
	newPubKey := &bec.PublicKey{}
	newPubKey.Curve = curve

	newPubKey.X = newX
	newPubKey.Y = newY

	return newPubKey, nil

}

func CheckTokenAgainstKeys(token []byte, pubKeyOld, pubKeyNew *bec.PublicKey) (bool, error) {
	curve := bec.S256()

	x, y := curve.ScalarBaseMult(token)

	var checkX *big.Int
	var checkY *big.Int

	checkX.Sub(pubKeyNew.X, pubKeyOld.X)
	checkY.Sub(pubKeyNew.Y, pubKeyOld.Y)

	rX := x.Cmp(checkX)
	rY := y.Cmp(checkY)

	if rX == 0 {
		if rY == 0 {
			return true, nil
		}
	}
	return false, nil

}
