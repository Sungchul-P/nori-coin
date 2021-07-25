package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/Sungchul-P/nori-coin/utils"
	"math/big"
)

const (
	signature     string = "9239255424EAE14D259D4EA25C0A2F87CB59F521765412236DA8BE2DA50609DE6B7775B1560878B165535528CD04127DF6573A3AB4313FD14B6418682D513B46"
	privateKey    string = "307702010104209229cf3b6b67587cd4fadc354a7bf1de6dbe94ea286f960f2ce43b8aa96f1530a00a06082a8648ce3d030107a14403420004ef1d02ecfb16e1e28e11f6fb1d5fd3366ac31004316e4431caa66bd7d2ae20078a7cd9b86f397cc386f405350bac1093401196747ed12fb89759aeba52dccef9"
	hashedMessage string = "a60a187658ba11c6022486f606015c968d4d293996737c0c6620223f83e61faf"
)

func Start() {

	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	retoredKey, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	fmt.Println(retoredKey)
	fmt.Println()
	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]
	fmt.Printf("%d\n\n%d\n\n%d\n\n", sigBytes, rBytes, sBytes)

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println()
	fmt.Println(bigR, bigS)
}
