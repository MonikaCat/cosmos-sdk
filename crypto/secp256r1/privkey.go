package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"io"
	"math/big"
)

var pubKeyCurve = elliptic.P256()

const PrivKeyNistp256Size = 256

// PrivKeyNistp256 implements crypto.PrivKey.
type PrivKeyNistp256 struct {
	PrivKey [PrivKeyNistp256Size]byte
}

func genPrivKey(rand io.Reader) (PrivKeyNistp256, error) {
	privatekey, err := ecdsa.GenerateKey(pubKeyCurve, rand) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		return _, err
	}
	x := new(big.Int)
	x = privatekey.D
	y := x.Bytes()

	var z PrivKeyNistp256
	copy(z.PrivKey[:], y)
	return z, nil
}
