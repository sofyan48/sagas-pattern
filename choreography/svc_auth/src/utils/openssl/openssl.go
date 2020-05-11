package openssl

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"strings"
)

type Openssl struct {
}

func OpensslHandler() *Openssl {
	return &Openssl{}
}

type OpensslInterface interface {
	GenerateKey(filename string) (string, error)
	MD5Hash(value string) string
}

// MD5Hash ...
func (ssl *Openssl) MD5Hash(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateKey ..
func (ssl *Openssl) GenerateKey(filename string) (string, error) {
	reader := rand.Reader
	bitSize := 1024

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return "", err
	}

	publicKey := key.PublicKey

	files := os.Getenv("SECRET_STORAGE") + "/" + filename

	// err = ssl.saveGobKey(files+".private.key", key)
	// if err != nil {
	// 	return "", err
	// }
	err = ssl.savePEMKey(files+".private.pem", key)
	if err != nil {
		return "", err
	}
	// err = ssl.saveGobKey(files+".public.key", publicKey)
	// if err != nil {
	// 	return "", err
	// }
	err = ssl.savePublicPEMKey(files+".public.pem", publicKey)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(files, "./", "/"), nil
}

func (ssl *Openssl) saveGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	return encoder.Encode(key)
}

func (ssl *Openssl) savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	return pem.Encode(outFile, privateKey)
}

func (ssl *Openssl) savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	return pem.Encode(pemfile, pemkey)
}
