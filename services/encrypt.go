package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"

	"github.com/macduyhai/SmartHomeVer2/config"
	// "github.com/macduyhai/SmartHomeVer2/middlewares"
)

type EncryptService interface {
	Base64Enc(b1 []byte) string
	Base64Dec(s1 string) ([]byte, error)
	RsaDecrypt(ciphertext []byte, key []byte) ([]byte, error)
	RsaEncrypt(origData []byte, key []byte) ([]byte, error)

	// Add(request dtos.AddRequest) (*dtos.AddResponse, error)
	// List(request dtos.ListRequest) (*dtos.ListResponse, error)
	// Delete(request dtos.DeleteRequest) (*dtos.DeviceResponse, error)
	// Edit(request dtos.EditRequest) (*dtos.EditResponse, error)
	// Upload(request dtos.UploadRequest) (*dtos.UploadResponse, error)
	// Getstatus(request dtos.GetstatusRequest) (*dtos.GetstatusResponse, error)
}

type ecryptServiceImpl struct {
	config *config.Config
}

//=================== MA HOA ==========================

func Base64Enc(b1 []byte) string {
	s1 := base64.StdEncoding.EncodeToString(b1)
	s2 := ""
	var LEN int = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}

func Base64Dec(s1 string) ([]byte, error) {
	s1 = strings.Replace(s1, "\n", "", -1)
	s1 = strings.Replace(s1, "\r", "", -1)
	s1 = strings.Replace(s1, " ", "", -1)
	return base64.StdEncoding.DecodeString(s1)
}

func RsaDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
