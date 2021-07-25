package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/Sungchul-P/nori-coin/utils"
)

const (
	signature     string = "9239255424EAE14D259D4EA25C0A2F87CB59F521765412236DA8BE2DA50609DE6B7775B1560878B165535528CD04127DF6573A3AB4313FD14B6418682D513B46"
	privateKey    string = "307702010104209229cf3b6b67587cd4fadc354a7bf1de6dbe94ea286f960f2ce43b8aa96f1530a00a06082a8648ce3d030107a14403420004ef1d02ecfb16e1e28e11f6fb1d5fd3366ac31004316e4431caa66bd7d2ae20078a7cd9b86f397cc386f405350bac1093401196747ed12fb89759aeba52dccef9"
	hashedMessage string = "a60a187658ba11c6022486f606015c968d4d293996737c0c6620223f83e61faf"
)

func Start() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)

	fmt.Printf("[Private Key]\n%x\n\n", keyAsBytes)

	utils.HandleErr(err)

	message := "11.2 Signing Messages"

	hashedMessage := utils.Hash(message)

	fmt.Println("hashedMessage : ", hashedMessage)

	hashAsBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	signature := append(r.Bytes(), s.Bytes()...)

	utils.HandleErr(err)

	fmt.Printf("\n[Signature]\n%X\n", signature)

	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)

	fmt.Println(ok)
}
