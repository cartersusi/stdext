package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

// Encrypt encrypts data using AES-GCM with a 256-bit key.
// If no key is provided, it will use the STDEXT_ENC_KEY environment variable.
// If that is not set, it will return an error.
//
// Parameters:
//   - data: the data to encrypt
//   - key: the key to use for encryption
//
// Returns:
//   - the encrypted data
//   - an error if one occurred
func Encrypt(data []byte, key ...interface{}) ([]byte, error) {
	var key_bytes []byte
	if len(key) > 0 {
		switch v := key[0].(type) {
		case string:
			key_bytes = []byte(v)
		case []byte:
			key_bytes = v
		default:
			return nil, errors.New("Invalid key type")
		}
	}

	if len(key_bytes) == 0 {
		key_string := os.Getenv("STDEXT_ENC_KEY")
		if key_string == "" {
			return nil, errors.New("STDEXT_ENC_KEY not set and no key provided")
		}
	}

	block, err := aes.NewCipher(key_bytes)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt decrypts data using AES-GCM with a 256-bit key.
// If no key is provided, it will use the STDEXT_ENC_KEY environment variable.
// If that is not set, it will return an error.
//
// Parameters:
//   - data: the data to decrypt
//   - key: the key to use for decryption
//
// Returns:
//   - the decrypted data
//   - an error if one occurred
func Decrypt(data []byte, key ...interface{}) ([]byte, error) {
	var key_bytes []byte
	if len(key) > 0 {
		switch v := key[0].(type) {
		case string:
			key_bytes = []byte(v)
		case []byte:
			key_bytes = v
		default:
			return nil, errors.New("Invalid key type")
		}
	}

	if len(key_bytes) == 0 {
		key_string := os.Getenv("STDEXT_ENC_KEY")
		if key_string == "" {
			return nil, errors.New("STDEXT_ENC_KEY not set and no key provided")
		}
	}

	block, err := aes.NewCipher(key_bytes)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
