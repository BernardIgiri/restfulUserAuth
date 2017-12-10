package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/bernardigiri/go-pkcs7"
)

type Decrypter struct {
	key []byte
}

func NewConfigDecrypter(key []byte) Decrypter {
	return Decrypter{key: key}
}

func (d Decrypter) Decrypt(encodedCipherText string) (plaintextBytes []byte, err error) {
	// DECODE
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCipherText)
	if err != nil {
		return
	}
	blockCipher, err := aes.NewCipher(d.key)
	if err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}
	// GET IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		err = errors.New("ciphertext is not a multiple of the block size")
		return
	}
	// DECRYPT
	mode := cipher.NewCBCDecrypter(blockCipher, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	// UNPAD
	plaintextBytes, err = pkcs7.Unpad(ciphertext, aes.BlockSize)
	return
}
