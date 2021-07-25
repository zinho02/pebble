package main

import (
	"crypto/pqc/dilithium"
	"fmt"
	"log"

	"github.com/liboqs-go/oqs"
	"gopkg.in/square/go-jose.v2"
)

func main() {
	sigName := "Dilithium5"
	signer := oqs.Signature{}
	defer signer.Clean() // clean up even in case of panic

	if err := signer.Init(sigName, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nSignature details:")
	fmt.Println(signer.Details())

	msg := []byte("This is the message to sign")
	//pubKey, err := signer.GenerateKeyPair()
	privKey, err := dilithium.GenerateKeyDilithium5()
	pubKey := privKey.PublicKey
	if err != nil {
		log.Fatal(err)
	}

	encrypter, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS512, Key: privKey},
		nil,
	)

	signature, _ := encrypter.Sign(msg)

	msg2, _ := signature.Verify(&pubKey)

	fmt.Println(string(msg2))
	//fmt.Println(signature)
}
