package jwt

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// return rsa private key
func GetRSAPrivateKey(filename string) *rsa.PrivateKey {
	signBytes, _ := ioutil.ReadFile(filename)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Println(err)
		return nil
	}
	return signKey
}

// return rsa public key
func GetRSAPublicKey(filename string) *rsa.PublicKey {
	verifyBytes, _ := ioutil.ReadFile(filename)
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Println(err)
		return nil
	}
	return verifyKey
}
