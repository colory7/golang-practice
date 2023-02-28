package crypto_demo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

// ecdsaCmd represents the doc command
func keyPairs(keyName string) {
	//elliptic.P256(),elliptic.P384(),elliptic.P521()

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	privateBs := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	privateFile, err := os.Create(keyName + ".private.pem")
	if err != nil {
		log.Fatal(err)
	}
	_, err = privateFile.Write(privateBs)
	if err != nil {
		log.Fatal(err)
	}
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	publicBs := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	publicKeyFile, err := os.Create(keyName + ".public.pem")
	if err != nil {
		log.Fatal(err)
	}
	_, err = publicKeyFile.Write(publicBs)
	if err != nil {
		log.Fatal(err)
	}
}
