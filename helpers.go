package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
)

func getCertificateModulus(filename string) (*big.Int, error) {
	pemData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM data")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain an RSA public key")
	}

	return publicKey.N, nil
}

func getPrivateKeyModulus(filename string) (*big.Int, error) {
	pemData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM data")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey.PublicKey.N, nil
}

func certificateFilesMatch(certificateFile string, privateKeyFile string) (bool, error) {
	key0, err := getCertificateModulus(certificateFile)
	if err != nil {
		return false, err
	}

	key1, err := getPrivateKeyModulus(privateKeyFile)
	if err != nil {
		return false, err
	}

	return key0.Cmp(key1) == 0, nil
}
