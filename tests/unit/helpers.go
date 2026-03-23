// Package helpers provides various helper functions for devops-scripts.
package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

const (
	// CertDir is the default directory for storing certificates.
	CertDir = "certs"
)

// GenerateKeyPair generates a new RSA key pair.
func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// CertPath returns the path to a certificate file.
func CertPath(commonName string) string {
	return filepath.Join(CertDir, fmt.Sprintf("%s.crt", commonName))
}

// KeyPath returns the path to a private key file.
func KeyPath(commonName string) string {
	return filepath.Join(CertDir, fmt.Sprintf("%s.key", commonName))
}

// CertPEM returns the PEM-encoded certificate.
func CertPEM(commonName string) (*pem.Block, error) {
	cert, err := newCert(commonName)
	if err != nil {
		return nil, err
	}
	return &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}, nil
}

// KeyPEM returns the PEM-encoded private key.
func KeyPEM(commonName string) (*pem.Block, error) {
	privateKey, _, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}, nil
}

// newCert generates a new certificate with the given common name.
func newCert(commonName string) ([]byte, error) {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: x509.Subject{
			CommonName: commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  false,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, template, template, publicKey, privateKey)
	if err != nil {
		return nil, err
	}
	return certDER, nil
}

// createCert creates a new certificate file.
func createCert(commonName string) error {
	certPEM, err := CertPEM(commonName)
	if err != nil {
		return err
	}
	keyPEM, err := KeyPEM(commonName)
	if err != nil {
		return err
	}
	certFile := CertPath(commonName)
	keyFile := KeyPath(commonName)
	if err := os.MkdirAll(CertDir, 0755); err != nil {
		return err
	}
	if err := os.WriteFile(certFile, certPEM.Bytes, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(keyFile, keyPEM.Bytes, 0600); err != nil {
		return err
	}
	return nil
}